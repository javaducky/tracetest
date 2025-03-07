# Running Tracetest with Elastic APM

:::note
[Check out the source code on GitHub here.](https://github.com/kubeshop/tracetest/tree/main/examples/tracetest-elasticapm-with-elastic-agent) 
:::

## Sample Node.js app with Elastic, OpenTelemetry and Tracetest

This is a simple quick start on how to configure a Node.js app to use Elastic APM Agent with traces, and Tracetest for enhancing your E2E and integration tests with trace-based testing. The infrastructure will use Elastic APM and Elasticsearch as the trace data store and OpenTelemetry Collector to receive traces from the Node.js app and send them to Elastic APM.

## Prerequisites

You will need [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/) installed on your machine to run this quick start app!

## Project structure

* `docker-compose.yml` - docker compose file that starts the whole environment
    * Elastic stack single node cluster with Elasticsearch, Kibana and, the APM Server.
    * OTel collector to support Tracetest.
    * Tracetest instance.
* `collector-config.yml` - OTel collector configuration file
* `app.js` - sample NodeJS application listening on port 8080 and instrumented with Elastic Nodejs APM agent.

The project is built with Docker Compose.

### 1. Node.js app
The `docker-compose.yaml` contains a service called `app` for the Node.js app.

### 2. Tracetest
The `docker-compose.yaml` file, `collector.config.yaml`, and `tracetest.config.yaml` contain configs for setting up Tracetest, Elastic APM, Elasticsearch, Kibana, and the OpenTelemetry Collector.

### Docker Compose Network
All `services` in the `docker-compose.yaml` are on the same network and will be reachable by hostname from within other services. E.g. `apm-server:8200` in the `elastic-apm-agent.js` will map to the `apm-server` service, where the port `8200` is the port where the Elastic APM accepts traces. And, `https://es01:9200` in the `tracetest.config.yaml` will map to the `es01` service and port `9200` where Tracetest will fetch trace data from Elasticsearch.

## Node.js app

The Node.js app is a simple Express app, contained in the `app.js` file.

The Elastic APM tracing is contained in the `elastic-apm-agent.js` file. Traces will be sent to the OpenTelemetry Collector.

Here's the content of the `elastic-apm-agent.js` file:

```js
const apm = require('elastic-apm-node').start({
  serviceName: 'sample-app',
  serverUrl: 'http://apm-server:8200',
})
```

Traces will be sent to either the Elastic APM endpoint.

The hostname and port is:

- HTTP: `http://apm-server:8200`

Enabling the tracer is done by preloading the trace file.

```bash
node -r ./elastic-apm-agent.js app.js
```

In the `package.json` you will see two npm script for running the respective tracers alongside the `app.js`.

```json
"scripts": {
  "with-elastic-apm-tracer":"node -r ./elastic-apm-agent.js app.js"
},
```

To start the server, run this command.

```bash
npm run with-elastic-apm-tracer
```

As you can see the `Dockerfile` uses the command above.

```Dockerfile
FROM node:slim
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 8080
CMD [ "npm", "run", "with-elastic-apm-tracer" ]
```

And, the `docker-compose.yaml` contains just one service for the Node.js app.

```yaml
app:
  image: quick-start-nodejs-elastic-apm
  hostname: app
  build: .
  ports:
    - "8080:8080"
  depends_on:
    apm-server:
      condition: service_started
```

## Tracetest

The `docker-compose.yaml` contains 7 services that configure Tracetest and Elastic.

- **Postgres** - Postgres is a prerequisite for Tracetest to work. It stores trace data when running the trace-based tests.
- [**OpenTelemetry Collector**](https://opentelemetry.io/docs/collector/) - A vendor-agnostic implementation of how to receive, process and export telemetry data.
- [**Elasticsearch**](https://www.elastic.co/elasticsearch/) - Data store and search engine. (Also contains a `setup` service to configure Elasticsearch) 
- [**Elastic APM**](https://www.elastic.co/observability/application-performance-monitoring) - Elastic application performance monitoring and traces.
- [**Kibana**](https://www.elastic.co/kibana/) - Kibana is a free and open user interface that lets you visualize your Elasticsearch data and navigate the Elastic Stack.
- [**Tracetest**](https://tracetest.io/) - Trace-based testing that generates end-to-end tests automatically from traces.

They will start in this order:

1. Postgres & Setup
2. Elasticsearch
3. Kibana
4. Elastic APM Server
5. OpenTelemetry Collector
6. Tracetest

[View the entire `docker-compose.yaml` file here](https://github.com/kubeshop/tracetest/blob/main/examples/tracetest-elasticapm-with-elastic-agent/docker-compose.yaml).

Tracetest depends on Postgres and the OpenTelemetry Collector. The OpenTelemetry Collector depends on the Elastic APM Server that then depends on Elasticsearch and Kibana.

Both Tracetest and the OpenTelemetry Collector require config files to be loaded via a volume.

Elasticsearch, Kibana, and Elastic APM use a `.env` file to load their config.

## Steps to start the environment

To start both the the environment run this command:

```bash
docker compose up -d
```

## Connecting Tracetest to Elastic APM

But how does Tracetest fetch traces?

Tracetest connects to Elastic APM to fetch trace data.

In the Web UI, open settings, and select Elastic APM.

![](https://res.cloudinary.com/djwdcmwdz/image/upload/v1674566041/Blogposts/Docs/screely-1674566018046_ci0st9.png)

Or, if you prefer using the CLI, you can use this file config.

```yaml
type: DataStore
spec:
  name: Elastic Data Store
  type: elasticapm
  isDefault: true
    elasticapm:
      addresses:
        - https://es01:9200
      username: elastic
      password: changeme
      index: traces-apm-default
      insecureSkipVerify: true
```

Proceed to run this command in the terminal, and specify the file above.

```bash
tracetest datastore apply -f my/data-store/file/location.yaml
```

### Step-by-step guide

Open `http://localhost:11633/` to configure the connection to Elasticsearch:

1. In Settings, configure Elastic APM as the Data Store.
2. Set `traces-apm-default` as the Index name.
3. Add the Address and set it to `https://es01:9200`.
4. Set the Username to `elastic` and password to `changeme`.
5. You will need to download the CA certificate from the docker image and upload it to the config under "Upload CA file".
    * The command to download the `ca.crt` file is:
    `docker cp tracetest-elasticapm-with-elastic-agent-es01-1:/usr/share/elasticsearch/config/certs/ca/ca.crt .`
    * Alternatively, you can skip CA certificate validation by setting the `Enable TLS but don't verify the certificate` option.
6. Test the connection and Save it, if all is successful.

Create a new test:
1. Use the "HTTP Request" option. Hit Next.
2. Name your test and add a description. Hit Next.
3. Configure the GET URL to be `http://app:8080` since the tests will be running in docker compose network. Hit Create.
4. Running the test should succeed.

## Open Kibana
Open `https://localhost:5601` and login using `elastic:changeme` credentials. The credentials can be changed in the `.env` file. Navigate to APM (upper lefthand corner menu) -> Services and you should see the `tracetest` service with the rest of the details.

## Steps to stop the environment
```bash
docker compose down -v

# Remove the built app docker image
docker rmi quick-start-nodejs:latest
```

## Learn more

Feel free to check out our [examples in GitHub](https://github.com/kubeshop/tracetest/tree/main/examples), and join our [Discord Community](https://discord.gg/8MtcMrQNbX) for more info!

# Running Tracetest with Tempo

:::note
[Check out the source code on GitHub here.](https://github.com/kubeshop/tracetest/tree/main/examples/quick-start-tempo-nodejs) 
:::

## Sample Node.js app with Tempo, OpenTelemetry and Tracetest

This is a simple quick start on how to configure a Node.js app to use OpenTelemetry instrumentation with traces and Tracetest for enhancing your e2e and integration tests with trace-based testing. The infrastructure will use Tempo as the trace data store, and OpenTelemetry Collector to receive traces from the Node.js app and send them to Tempo.

## Prerequisites

You will need [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/) installed on your machine to run this quick start app!

## Project Structure

The project is built with Docker Compose. It contains two distinct `docker-compose.yaml` files.

### 1. Node.js app
The `docker-compose.yaml` file and `Dockerfile` in the root directory are for the Node.js app.

### 2. Tracetest
The `docker-compose.yaml` file, `collector.config.yaml`, and `tracetest.config.yaml` in the `tracetest` directory are for the setting up Tracetest, Tempo, and the OpenTelemetry Collector.

The `tracetest` directory is self-contained and will run all the prerequisites for enabling OpenTelemetry traces and trace-based testing with Tracetest.

### Docker Compose Network
All `services` in the `docker-compose.yaml` are on the same network and will be reachable by hostname from within other services. E.g. `tempo:4317` in the `collector.config.yaml` will map to the `tempo` service, where the port `4317` is the port where Tempo accepts traces. And, `tempo:9095` in the `tracetest.config.yaml` will map to the `tempo` service and port `9095` where Tracetest will fetch trace data from Tempo.

## Node.js app

The Node.js app is a simple Express app, contained in the `app.js` file.

The OpenTelemetry tracing is contained in the `tracing.otel.grpc.js` or `tracing.otel.http.js` files, respectively.
Traces will be sent to the OpenTelemetry Collector.

Here's the content of the `tracing.otel.grpc.js` file:

```js
const opentelemetry = require('@opentelemetry/sdk-node')
const { getNodeAutoInstrumentations } = require('@opentelemetry/auto-instrumentations-node')
const { OTLPTraceExporter } = require('@opentelemetry/exporter-trace-otlp-grpc');

const sdk = new opentelemetry.NodeSDK({
  traceExporter: new OTLPTraceExporter({ url: 'http://otel-collector:4317' }),
  instrumentations: [getNodeAutoInstrumentations()],
})
sdk.start()
```

Depending on which of these you choose, traces will be sent to either the `grpc` or `http` endpoint.

The hostnames and ports for these are:

- GRPC: `http://otel-collector:4317`
- HTTP: `http://otel-collector:4318/v1/traces`

Enabling the tracer is done by preloading the trace file.

```bash
node -r ./tracing.otel.grpc.js app.js
```

In the `package.json` you will see two npm scripts for running the respective tracers alongside the `app.js`.

```json
"scripts": {
  "with-grpc-tracer":"node -r ./tracing.otel.grpc.js app.js",
  "with-http-tracer":"node -r ./tracing.otel.http.js app.js"
},
```

To start the server, run this command:

```bash
npm run with-grpc-tracer
# or
npm run with-http-tracer
```

As you can see the `Dockerfile` uses the command above.

```Dockerfile
FROM node:slim
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 8080
CMD [ "npm", "run", "with-grpc-tracer" ]
```

And, the `docker-compose.yaml` contains just one service for the Node.js app.

```yaml
version: '3'
services:
  app:
    image: quick-start-nodejs
    build: .
    ports:
      - "8080:8080"
```

To start it, run this command:

```bash
docker compose build # optional if you haven't already built the image
docker compose up
```

This will start the Node.js app. But, you're not sending the traces anywhere.

Let's fix this by configuring Tracetest and OpenTelemetry Collector.

## Tracetest

The `docker-compose.yaml` in the `tracetest` directory is configured with four services.

- **Postgres** - Postgres is a prerequisite for Tracetest to work. It stores trace data when running the trace-based tests.
- [**OpenTelemetry Collector**](https://opentelemetry.io/docs/collector/) - A vendor-agnostic implementation of how to receive, process and export telemetry data.
- [**Tempo**](https://grafana.com/oss/tempo/) - Grafana Tempo is an open source, easy-to-use, and high-scale distributed tracing backend.
- [**Tracetest**](https://tracetest.io/) - Trace-based testing that generates end-to-end tests automatically from traces.

```yaml
version: '3'
services:

  tracetest:
    image: kubeshop/tracetest
    platform: linux/amd64
    volumes:
      - ./tracetest/tracetest.config.yaml:/app/config.yaml
    ports:
      - 11633:11633
    depends_on:
      postgres:
        condition: service_healthy
      tempo:
        condition: service_started
      otel-collector:
        condition: service_started
    healthcheck:
      test: ["CMD", "wget", "--spider", "localhost:11633"]
      interval: 1s
      timeout: 3s
      retries: 60

  postgres:
    image: postgres:14
    environment:
      POSTGRES_PASSWORD: postgres
      POSTGRES_USER: postgres
    healthcheck:
      test: pg_isready -U "$$POSTGRES_USER" -d "$$POSTGRES_DB"
      interval: 1s
      timeout: 5s
      retries: 60

  otel-collector:
    image: otel/opentelemetry-collector-contrib:0.59.0
    command:
      - "--config"
      - "/otel-local-config.yaml"
    volumes:
      - ./tracetest/collector.config.yaml:/otel-local-config.yaml
    depends_on:
      - tempo

  tempo:
    image: grafana/tempo:1.5.0
    command: ["-config.file=/etc/tempo.yaml"]
    volumes:
    - ./tracetest/tempo.config.yaml:/etc/tempo.yaml

```

Tracetest depends on Postgres, Tempo and the OpenTelemetry Collector. All three services require config files to be loaded via a volume. The volumes are mapped from the root directory into the `tracetest` directory and the respective config files.

**Why?** To start both the Node.js app and Tracetest, we will run this command:

```bash
docker-compose -f docker-compose.yaml -f tracetest/docker-compose.yaml up # add --build if the images are not built already
```

The `tempo.config.yaml` file contains the initial config for running Tempo.

The key takeaway is the server block.

```yaml
#...
server:
  http_listen_port: 3100
  grpc_listen_port: 9095
#...
```

We'll use this below.

Check out the full Tempo config for reference.

```yaml
auth_enabled: false

server:
  http_listen_port: 3100
  grpc_listen_port: 9095

distributor:
  receivers:                           # this configuration will listen on all ports and protocols that tempo is capable of.
    jaeger:                            # the receives all come from the OpenTelemetry collector.  more configuration information can
      protocols:                       # be found there: https://github.com/open-telemetry/opentelemetry-collector/tree/master/receiver
        thrift_http:                   #
        grpc:                          # for a production deployment you should only enable the receivers you need!
        thrift_binary:
        thrift_compact:
    zipkin:
    otlp:
      protocols:
        http:
        grpc:
    opencensus:

ingester:
  trace_idle_period: 10s               # the length of time after a trace has not received spans to consider it complete and flush it
  max_block_bytes: 1_000_000           # cut the head block when it hits this size or ...
  #traces_per_block: 1_000_000
  max_block_duration: 5m               #   this much time passes

compactor:
  compaction:
    compaction_window: 1h              # blocks in this time window will be compacted together
    max_compaction_objects: 1000000    # maximum size of compacted blocks
    block_retention: 1h
    compacted_block_retention: 10m

storage:
  trace:
    backend: local                     # backend configuration to use
    wal:
      path: /tmp/tempo/wal            # where to store the the wal locally
      #bloom_filter_false_positive: .05 # bloom filter false positive rate.  lower values create larger filters but fewer false positives
      #index_downsample: 10             # number of traces per index record
    local:
      path: /tmp/tempo/blocks
    pool:
      max_workers: 100                 # the worker pool mainly drives querying, but is also used for polling the blocklist
      queue_depth: 10000

```

The `tracetest.config.yaml` file contains the basic setup of connecting Tracetest to the Postgres instance, and defining the trace data store and exporter. The data store is set to Tempo, meaning the traces will be stored in Tempo and Tracetest will fetch them from Tempo when running tests. The exporter is set to the OpenTelemetry Collector.

But how does Tracetest fetch traces?

Tracetest uses `tempo.endpoint:tempo:9095` to connect to Tempo and fetch trace data.

```yaml
# tracetest.config.yaml

postgresConnString: "host=postgres user=postgres password=postgres port=5432 sslmode=disable"

poolingConfig:
  maxWaitTimeForTrace: 10m
  retryDelay: 5s

googleAnalytics:
  enabled: true

demo:
  enabled: []

experimentalFeatures: []

telemetry:
  dataStores:
    tempo:
      type: tempo
      tempo:
        endpoint: tempo:9095
        tls:
          insecure: true

server:
  telemetry:
    dataStore: tempo
    exporter: collector
    applicationExporter: collector

```

How do traces reach Tempo?

The `collector.config.yaml` explains that. It receives traces via either `grpc` or `http`. Then, exports them to Tempo's OTLP gRPC endpoint `tempo:4317`.

```yaml
# collector.config.yaml
receivers:
  otlp:
    protocols:
      grpc:
      http:

processors:
  batch:
    timeout: 100ms

exporters:
  logging:
    loglevel: debug
  otlp/2:
    endpoint: tempo:4317
    tls:
      insecure: true

service:
  pipelines:
    traces:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp/2]

```

## Run both the Node.js app and Tracetest

To start both the Node.js app and Tracetest, we will run this command:

```bash
docker-compose -f docker-compose.yaml -f tracetest/docker-compose.yaml up # add --build if the images are not built already
```

This will start your Tracetest instance on `http://localhost:11633/`. Go ahead and open it up.

Start creating tests! Make sure to use the `http://app:8080/` url in your test creation, because your Node.js app and Tracetest are in the same network.

## Learn more

Feel free to check out our [examples in GitHub](https://github.com/kubeshop/tracetest/tree/main/examples), and join our [Discord Community](https://discord.gg/8MtcMrQNbX) for more info!

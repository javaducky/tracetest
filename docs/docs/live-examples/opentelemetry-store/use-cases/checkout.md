

# OpenTelemetry Store - Add item into the shopping cart

In this use case, we want to validate the following story:

```
As a consumer, after choosing products
I want to pay for all products in my shopping cart
So I can ship them to my address and use them.
```

TODO: simple mermaid about this service

You can trigger this use case by calling the endpoint `POST /api/cart`, with the following request body:
```json
{
  "item": {
    "productId": "OLJCESPC7Z",
    "quantity": 1
  },
  "userId": "2491f868-88f1-4345-8836-d5d8511a9f83"
}
```

It should return a payload similar to this:
```json
{
  "userId": "2491f868-88f1-4345-8836-d5d8511a9f83",
  "items": [
    {
      "productId": "OLJCESPC7Z",
      "quantity": 1
    }
  ]
}
```

https://github.com/open-telemetry/opentelemetry-demo/blob/main/src/frontend/cypress/e2e/Checkout.cy.ts

## Building a Test for This Scenario

Using Tracetest, we can [create a test](../../../web-ui/creating-tests.md) that will execute an API call on `POST /api/cart` and validate two properties:
- TODO

### Traces

Running these tests for the first time will create an Observability trace like the image above, where you can see spans for the API call, validation (an API internal operation), and database calls:
- TODO: add trace image

### Assertions

With this trace, now we can build [assertions](../../../concepts/assertions.md) on Tracetest and validate the API response and the database latency:

- TODO: explain each assertion

Now you can validate this entire use case.

### Test Definition

If you want to replicate this entire test on Tracetest see by yourself, you can replicate these steps on our Web UI or using our CLI, saving the following test definition as the file `test-definition.yml` and later running:

```sh
tracetest test -d test-definition.yml --wait-for-results
```

```yaml
type: Test

```

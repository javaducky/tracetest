# OpenTelemetry Store - Check shopping cart contents

In this use case, we want to validate the following story:

```
As a consumer
I want to see my shopping cart
And see all products that I've added to it
So I can go to the checkout process
```

You can trigger this use case by calling the endpoint `GET /api/cart?sessionId={some-uuid}&currecyCode=`. It should return a payload similar to this:
```json
{
  "userId": "8c0465e2-32bb-4ecb-a9c8-5a2861629ff1",
  "items": [
    {
      "productId": "66VCHSJNUP",
      "quantity" : 1,
      "product": {
        "id": "66VCHSJNUP",
        "name": "Starsense Explorer Refractor Telescope",
        "description": "The first telescope that uses your smartphone to analyze the night sky and calculate its position in real time. StarSense Explorer is ideal for beginners thanks to the app’s user-friendly interface and detailed tutorials. It’s like having your own personal tour guide of the night sky",
        "picture": "/images/products/StarsenseExplorer.jpg",
        "priceUsd": {
          "currencyCode": "USD",
          "units": 349,
          "nanos": 950000000
        },
        "categories": [
          "telescopes"
        ]
      }
    }
  ]
}
```

## Building a Test for This Scenario

Using Tracetest, we can [create a test](../../../web-ui/creating-tests.md) that will execute an API call on `GET /api/cart?sessionId={some-uuid}&currecyCode=` and validate two properties:
- The product ID `66VCHSJNUP`, previously added, exists in the cart.
- The size of the shopping cart should be 1.

### Traces

Running these tests for the first time will create an Observability trace like the image above, where you can see spans for the API call, and database calls:
![](../images/check-shopping-cart-contents-trace.png)

### Assertions

With this trace, now we can build [assertions](../../../concepts/assertions.md) on Tracetest and validate the API response and the database latency:

- **The product ID `66VCHSJNUP`, previously added, exists in the cart.**
![](../images/check-shopping-cart-contents-product-catalog.png)

- **The size of the shopping cart should be 1.**
![](../images/check-shopping-cart-contents-item-lenght.png)

Now you can validate this entire use case.

### Test Definition

If you want to replicate this entire test on Tracetest see by yourself, you can replicate these steps on our Web UI or using our CLI, saving the following test definition as the file `test-definition.yml` and later running:

```sh
tracetest test -d test-definition.yml --wait-for-results
```

```yaml
type: Test
spec:
  id: 1JnDAc04g
  name: OpenTelemetry Store - Check shopping cart contents
  trigger:
    type: http
    httpRequest:
      url: http://otel-demo-v2.tracetest.io/api/cart?sessionId=8c0465e2-32bb-4ecb-a9c8-5a2861629ff1&currencyCode=
      method: GET
      headers:
      - key: Content-Type
        value: application/json
  specs:
  - selector: span[tracetest.span.type="rpc" name="hipstershop.ProductCatalogService/GetProduct"
      rpc.system="grpc" rpc.method="GetProduct" rpc.service="hipstershop.ProductCatalogService"]
    assertions:
    - attr:app.product.id = "66VCHSJNUP"
  - selector: span[tracetest.span.type="general" name="Tracetest trigger"]
    assertions:
    - attr:tracetest.response.body | json_path '$.items.lenght' > 0
```

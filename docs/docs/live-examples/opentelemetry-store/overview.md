# OpenTelemetry Astronomy Shop Demo

This system implements an Astronomy shop in a set of microservices in different languages with OpenTelemetry enabled, intended to be used as an example to OpenTelemetry instrumentation and observability.

- **Source Code**: https://github.com/open-telemetry/opentelemetry-demo
- **Running it locally**: [instructions](https://github.com/open-telemetry/opentelemetry-demo/blob/main/docs/docker_deployment.md#run-docker-compose)
- **Running on Kubernetes**: [instructions](https://github.com/open-telemetry/opentelemetry-demo/blob/main/docs/kubernetes_deployment.md)

## Running with Tracetest

TODO

## Use cases

- [Add item into shopping cart](./use-cases/add-item-into-shopping-cart.md): simulate a user choosing an item and adding it to the shopping cart
- [Check shopping cart content](./use-cases/check-shopping-cart-contents.md): simulate a user choosing different products and checking the shopping cart later 
- [Checkout](./use-cases/checkout.md): simulates a user choosing a product and later doing a checkout of that product, with billing and shipping info
- [Get recommended products](./use-cases/get-recommended-products.md): simulates a user querying for recommended products

## System architecture

This demonstration environment consists in a series of microservices, handling each aspect of the store, like Product Catalog, Payment, Currency, etc.

A detailed description of these services can be seen [here](https://github.com/open-telemetry/opentelemetry-demo/tree/main/docs#service-documentation)
and the architecture diagrams can be seen [here](https://github.com/open-telemetry/opentelemetry-demo/blob/main/docs/current_architecture.md).

# oas-combiner

**oas-combiner** merges many OpenAPI (Swagger) files into one with additional features.

oas-combiner expecting one yaml file with follow structure:
```yaml
openapi: 3.0.0
info:
  title: My Gateway
  description: My Gateway
  version: 0.1.0
servers:
  - url: 'https://gw.prod'
    description: Production

combine:
  include:
    - $ref: auth.yaml # oas-combiner extract paths and components from file
    - $ref: users.yaml

  defaultResponses: # optional
    401: { $ref: "#/components/responses/error-unauthorized-response" }
    403: { $ref: "#/components/responses/error-forbidden-response" }
    #...
```

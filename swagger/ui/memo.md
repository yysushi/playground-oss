# Playground for Swagger UI

## Docker

<https://github.com/swagger-api/swagger-ui/blob/master/docs/usage/installation.md#docker>

```shell-session
docker pull swaggerapi/swagger-ui
docker run -p 80:8080 swaggerapi/swagger-ui
```

```shell-session
docker run -p 80:8080 -v $(pwd):/tmp -e SWAGGER_JSON=/tmp/openapi.yaml swaggerapi/swagger-ui
```

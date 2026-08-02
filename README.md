# Swagger / Swaggo

This project exposes Swagger UI at `/swagger/index.html` and generates the OpenAPI spec from `swaggo` annotations.

To refresh the generated docs after changing handlers or routes:

```bash
swag init -g cmd/api/main.go -o cmd/api/docs --parseDependency --parseInternal
```

The generated files live in `cmd/api/docs/` and are imported by `cmd/api/main.go`.

# Findings of OpenAPI Generator

- Unable to generate e.g. a usecase file for each operation as OpenAPI generator groups by tag.
  Seems a workaround exists: https://github.com/OpenAPITools/openapi-generator/issues/5854
- Unclear if it supports reusable snippets in templates

# Alternative: Parse OpenAPI file and perform templating in custom Go library

1. Parse spec with a library like https://pkg.go.dev/github.com/getkin/kin-openapi/openapi3
2. Define a tree structure containing the output (challenge: support repeated files on multiple levels like tags, operations, models, ...)

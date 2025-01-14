# 😽 Cat Fact API
[![Go Reference](https://pkg.go.dev/badge/github.com/go-api-libs/cat-facts.svg)](https://pkg.go.dev/github.com/go-api-libs/cat-facts/pkg/catfacts)
[![Official Documentation](https://img.shields.io/badge/docs-API-blue)](https://alexwohlbruck.github.io/cat-facts/docs/)
[![OpenAPI](https://img.shields.io/badge/OpenAPI-3.1-blue)](/api/openapi.json)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-api-libs/cat-facts)](https://goreportcard.com/report/github.com/go-api-libs/cat-facts)
![Code Coverage](https://img.shields.io/badge/coverage-43%25-orange)
![API Health](https://img.shields.io/badge/API_health-89%25-green)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

An API that shows different Cat Facts! ([Source](https://freepublicapis.com/cat-fact-api))

## Installation

To install the library, use the following command:

```shell
go get github.com/go-api-libs/cat-facts/pkg/catfacts
```

## Usage

### Example 1: Get random cat facts

```go
package main

import (
	"context"

	"github.com/go-api-libs/cat-facts/pkg/catfacts"
)

func main() {
	c, err := catfacts.NewClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	facts, err := c.GetRandom(ctx, &catfacts.GetRandomParams{
		Amount:     1,
		AnimalType: "cat",
	})
	if err != nil {
		panic(err)
	}

	// Use facts array
}

```

### Example 2: 

```go
package main

import (
	"context"

	"github.com/go-api-libs/cat-facts/pkg/catfacts"
)

func main() {
	c, err := catfacts.NewClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	fact, err := c.GetFactByID(ctx, "591f9890d369931519ce3564")
	if err != nil {
		panic(err)
	}

	// Use fact object
}

```

## Additional Information

- [**Go Reference**](https://pkg.go.dev/github.com/go-api-libs/cat-facts/pkg/catfacts): The Go reference documentation for the client package.
- [**Official Documentation**](https://alexwohlbruck.github.io/cat-facts/docs/): The official API documentation.
- [**OpenAPI Specification**](./api/openapi.json): The OpenAPI 3.1.0 specification.
- [**Go Report Card**](https://goreportcard.com/report/github.com/go-api-libs/cat-facts): Check the code quality report.

## Contributing

If you have any contributions to make, please submit a pull request or open an issue on the [GitHub repository](https://github.com/go-api-libs/cat-facts).

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

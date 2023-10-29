# dashotv/tmdb

Golang TMDB (v3) API client

[![Build Status](https://drone.dasho.net/api/badges/dashotv/tmdb/status.svg?ref=refs/heads/main)](https://drone.dasho.net/dashotv/tmdb)
[![GoDoc](https://godoc.org/github.com/dashotv/tmdb?status.svg)](https://godoc.org/github.com/dashotv/tmdb)
[![Go Report Card](https://goreportcard.com/badge/github.com/dashotv/tmdb)](https://goreportcard.com/report/github.com/dashotv/tmdb)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

## Generated Code

This package is generated from the
[TMDB OpenAPI](https://developer.themoviedb.org/reference/intro/getting-started)
specification using the [Speakeasy](https://speakeasyapi.dev) code generator.

The generated code is in the `openapi` directory, and I've written scripts to wrap
it with a more convenient interface.

See the `Generation` section below for more information.

> [!NOTE]
> I have needed to make a few tweaks to the openapi spec to get things to work correctly in go.

## Additional Documentation

The Speakeasy generator also generates documentation for the API. This is available
in the openapi directory's [README.md](openapi/README.md) and the [docs](openapi/docs) directory.

## Status

The code is currently in alpha, and is not ready for production use. It is being used
by the DashoTV project, but is not yet stable.

## Usage

Install the package with:

```bash
go get github.com/dashotv/tmdb
```

Import the package with:

```go
import "github.com/dashotv/tmdb"
```

Create a new client with the code below. This will create the client and configure Bearer authentication.

```go
client := tvdb.New("TOKEN")
```

Your TMDB [Account page API section](https://www.themoviedb.org/settings/api) will have the token. See their docs for more info.

## Development

Crate a local `.env` file with the following content:

```
# .env
TMDB_API_TOKEN=<TOKEN>
```

Run the following to get the make targets:

```
> make help
Usage:
  make <target>

Targets:
  General:
    generate            Generate code from openapi.yml spec
    clean               Remove build related file
  Test:
    test                Run the tests of the project
    coverage            Run the tests of the project and export the coverage
  Lint:
    lint                Run all available linters
    lint-go             Use golintci-lint on your project
  Help:
    help                Show this help.

```

### Generation

To update the generated code, ensure you have the dependencies installed:

- Speakeasy CLI - See [Speakeasy](https://speakeasyapi.dev) for more information.
- Ruby - See [Ruby](https://www.ruby-lang.org/en/documentation/installation/) for more information.

Then run:

```bash
make generate
```

This will run the `scripts/generate.sh` script, which will:

- Run the speakeasy cli generator
- Rearrange the generated code
- Run a ruby script to build the wrapper client code

### Notes

There are several operations on the api that are disabled (this is handled in the
ruby script). These are operations that require additional priveleges (actions on
behalf of the user) or that don't work correctly. These are configured as an array
in the ruby script.

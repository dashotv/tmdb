# dashotv/tmdb

TMDB API in Go

[![Build Status](https://drone.dasho.net/api/badges/dashotv/tmdb/status.svg?ref=refs/heads/main)](https://drone.dasho.net/dashotv/tmdb)
[![GoDoc](https://godoc.org/github.com/dashotv/tmdb?status.svg)](https://godoc.org/github.com/dashotv/tmdb)
[![Go Report Card](https://goreportcard.com/badge/github.com/dashotv/tmdb)](https://goreportcard.com/report/github.com/dashotv/tmdb)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

## Usage

Install the package with:

```bash
go get github.com/dashotv/tmdb
```

Import the package with:

```go
import "github.com/dashotv/tvdb"
```

Create a new client with the code below. This will create the client and configure Bearer authentication.

```go
client, err := tvdb.New("https://api.themoviedb.org", "TOKEN")
```

Your TMDB [Account page API section](https://www.themoviedb.org/settings/api) will have the token. See their docs for more info.

## Development

Crate a local `.env` file with the following content:

```
# .env
TMDB_API_URL=https://api.themoviedb.org
TMDB_API_TOKEN=<TOKEN>
```

Run the following to get the make targets:

> make help

```

Usage:
  make <target>

Targets:
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

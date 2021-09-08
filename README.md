# football-data-sdk #
[![football-data-sdk (latest SemVer)](https://img.shields.io/github/v/release/matheustex/football-data-sdk?sort=semver)](https://github.com/matheustex/football-data-sdk/releases)
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/matheustex/football-data-sdk)
[![Test Status](https://github.com/google/go-github/workflows/tests/badge.svg)](https://github.com/matheustex/football-data-sdk/actions?query=workflow%3Atests)

football-data-sdk is a Go client library for accessing the [Football Data API](https://www.football-data.org/documentation/api).

Successful queries return native Go structs.

### Services

* Areas
* Competitions
* Matches
* Players
* Teams

## Installation ##

football-data-sdk is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/matheus-tex/football-data-sdk
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/matheus-tex/football-data-sdk"
```

and run `go get` without parameters.

## Usage ##

```go
import "github.com/matheus-tex/football-data-sdk"
```

```go
client := football.NewClient()

// list all competitions
competitions, err := client.Competitions.List(context.Background(), nil)
```

Some API methods have optional parameters that can be passed. For example:

```go
client := github.NewClient(nil)

// list public matches for a player
filters := &football.PlayerFiltersOptions{Limit: "5"}
matches, err := client.Players.Matches(context.Background(), "1", filters)
```

## License ##

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.

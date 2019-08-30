# GraphQL Playground Handler Function with Playground Settings in Golang

[![CircleCI]][CircleCILink]
[![Maintainability]][MILInk] [![Test Coverage]][TCLink] [![Go Report Card]][GRCLink]

[CircleCI]: https://circleci.com/gh/hiroaki-yamamoto/gqlplay.svg?style=svg
[CircleCILink]: https://circleci.com/gh/hiroaki-yamamoto/gqlplay
[Maintainability]: https://api.codeclimate.com/v1/badges/e199f4a18d1690f650cb/maintainability
[MILink]: https://codeclimate.com/github/hiroaki-yamamoto/gqlplay/maintainability
[Test Coverage]: https://api.codeclimate.com/v1/badges/e199f4a18d1690f650cb/test_coverage
[TCLink]: https://codeclimate.com/github/hiroaki-yamamoto/gqlplay/test_coverage
[Go Report Card]: https://goreportcard.com/badge/github.com/hiroaki-yamamoto/gqlplay
[GRCLink]: https://goreportcard.com/report/github.com/hiroaki-yamamoto/gqlplay


## What this?

This is a go binding of [Prisma's GraphQL PlayGround] with setting configuration
support.

[Prisma's GraphQL PlayGround]: https://github.com/prisma/graphql-playground

## Why I made this?

I like using GraphQL, especially, [GQLGen] is awesome.
However, the playground provided by [GQLGen] has some small-issues. i.e. I cam't
set configuration to handle CSRF. So, I decided to re-invent the wheel that can
set the configuration.

[GQLGen]: https://github.com/99designs/gqlgen

## How to use?

```Go
package main

import "log"

import (
  "github.com/hiroaki-yamamoto/gqlplay"
  "github.com/go-chi/chi"
)

func main() {
  player := gqlplay.Ground(gqlplay.Option{
    Settings: goplay.Settings{
      PollingEnabled: false,
      Credentials: SameOriginCredentials,
      HideTraceResponse: true,
    },
    Headers: map[string]string{
      "X-CSRF-TOKEN": "blablablabla...",
    }
  })
  r := chi.NewRouter()
  //...Other code...
  r.Get("/playground", player)
  svr := http.Server{
		Addr:    fmt.Sprintf("%s:%d", setting.Host, setting.Port),
		Handler: chi.ServerBaseContext(ctx, r),
  }
  log.Fatal(svr.ListenAndServe())
}
```

It must be easy.

## Contribution

Sending Issues / PR is welcome. PR is more appreciated. However, if you changed
the behavior, the corresponding changes, additions, and/or deletions of
the test code is mandatory, if you changed only the doc, the test
code is not needed.

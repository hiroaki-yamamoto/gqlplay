# GraphQL Playground Handler Function with Playground Settings in Golang

[![Maintainability]][MILInk] [![Test Coverage]][TCLink] [![Go Report Card]][GRCLink]

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

<!-- ## How to use?

```Go
package main

import "github.com/hiroaki-yamamoto/gqlplay"

func main() {
  gqlplay.Ground(gqlplay.Config{
    "schema.polling.enable": false,
    "request.credentials": "same-origin",
  })
}
``` -->

[GQLGen]: https://github.com/99designs/gqlgen

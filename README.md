# Data race in fuego 

Simple reproduction of a data race inside [fuego](https://github.com/go-fuego/fuego), when multiple
`fuego.Server` are created in parallel.

Run `go test -race .` to see the data race.

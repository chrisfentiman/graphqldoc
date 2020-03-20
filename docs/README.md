# Documentation generator for GraphQL distributed via NPM

Markdown generator for documenting GraphQL schema

## Use

Generate dir `doc/` with markdown files

```bash
$ graphqldoc http://localhost:8080/graphql
```

## Compile

Is need `go-bindata` [https://github.com/go-bindata/go-bindata](https://github.com/go-bindata/go-bindata)

```
go get github.com/go-bindata/go-bindata/...
```

```bash
$ go-bindata -o assets.go template/
$ sed -i '' 's/package\ main/package\ graphqldoc/g' assets.go # Change package of assets.go file
$ cd cmd/
$ go install -v
```

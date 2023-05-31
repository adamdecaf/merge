# merge

[![GoDoc](https://godoc.org/github.com/adamdecaf/merge?status.svg)](https://godoc.org/github.com/adamdecaf/merge)
[![Build Status](https://github.com/adamdecaf/merge/workflows/Go/badge.svg)](https://github.com/adamdecaf/merge/actions)
[![Coverage Status](https://codecov.io/gh/adamdecaf/merge/branch/master/graph/badge.svg)](https://codecov.io/gh/adamdecaf/merge)
[![Go Report Card](https://goreportcard.com/badge/github.com/adamdecaf/merge)](https://goreportcard.com/report/github.com/adamdecaf/merge)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/adamdecaf/merge/master/LICENSE)

merge is a Go package for quickly combining slices of objects together into. This package supports updating objects with the same key and returns sorted results.

## Install

```
go get github.com/adamdecaf/merge
```

## Example

Given a few example games where players score goals we can combine those to find their total scores.

```go
type Player struct {
	Name  string
	Goals int
}

game1 := []Player{
	{Name: "John Doe", Goals: 2},
	{Name: "Jane Doe", Goals: 1},
}
game2 := []Player{
	{Name: "Jane Doe", Goals: 2},
}
game3 := []Player{
	{Name: "Other Guy", Goals: 1},
}
game4 := []Player{
	{Name: "Jane Doe", Goals: 1},
	{Name: "Other Girl", Goals: 5},
}

out := Slices(
	func(p Player) string {
		return p.Name
	},
	func(p1 *Player, p2 *Player) {
		p1.Goals += p2.Goals
	},
	game1, game2, game3, game4,
)


expected := []Player{
	{Name: "John Doe", Goals: 2},
	{Name: "Jane Doe", Goals: 4},
	{Name: "Other Guy", Goals: 1},
	{Name: "Other Girl", Goals: 5},
}
```

## Supported and tested platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.

package merge

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type Player struct {
	Name  string
	Goals int
}

func TestSlices_Basic(t *testing.T) {
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

	expected := []Player{
		{Name: "John Doe", Goals: 2},
		{Name: "Jane Doe", Goals: 4},
		{Name: "Other Guy", Goals: 1},
		{Name: "Other Girl", Goals: 5},
	}

	out := Slices(
		func(p Player) string {
			return p.Name
		},
		func(p1 *Player, p2 Player) {
			p1.Goals += p2.Goals
		},
		game1, game2, game3, game4,
	)
	require.ElementsMatch(t, expected, out)
}

func TestSlices_Large(t *testing.T) {
	iterations := 10000
	mod := 27

	var games [][]Player
	for i := 0; i < iterations; i++ {
		var players []Player
		for j := 1; j < iterations; j *= 3 {
			players = append(players, Player{
				Name:  fmt.Sprintf("Forward %02.2d", i%mod),
				Goals: j / mod,
			})
		}
		games = append(games, players)
	}

	out := Slices(
		func(p Player) string {
			return p.Name
		},
		func(p1 *Player, p2 Player) {
			p1.Goals += p2.Goals
		},
		games...,
	)
	require.Len(t, out, mod)

	for i := range out {
		require.Equal(t, fmt.Sprintf("Forward %02.2d", i), out[i].Name)

		if i < 10 {
			require.Equal(t, 135044, out[i].Goals)
		} else {
			require.Equal(t, 134680, out[i].Goals)
		}
	}
}

func TestMergeStrings(t *testing.T) {
	a := []string{"b", "C", "A"}
	b := []string{"d", "D", "a"}

	got := Slices(
		func(s string) string {
			return strings.ToLower(s)
		},
		nil, // do nothing, just unique
		a, b,
	)

	expected := []string{"A", "b", "C", "d"}
	require.ElementsMatch(t, expected, got)
}

func TestEmpty(t *testing.T) {
	a := []string{"b", "C", "A"}
	b := []string{"d", "D", "a"}

	var fn func(string) string

	got := Slices(fn, nil, a, b)
	require.Empty(t, got)
}

func BenchmarkSlices(b *testing.B) {
	// Prepare data outside the timed loop (adapted from TestSlices_Large)
	iterations := 10000
	mod := 27

	var games [][]Player
	for i := 0; i < iterations; i++ {
		var players []Player
		for j := 1; j < iterations; j *= 3 {
			players = append(players, Player{
				Name:  fmt.Sprintf("Forward %02.2d", i%mod),
				Goals: j / mod,
			})
		}
		games = append(games, players)
	}

	keyFn := func(p Player) string {
		return p.Name
	}
	combiner := func(existing *Player, incoming Player) {
		existing.Goals += incoming.Goals
	}

	// Reset timer and run the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Slices(keyFn, combiner, games...)
	}
}

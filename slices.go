package merge

import (
	"github.com/igrmk/treemap/v2"
	"golang.org/x/exp/constraints"
)

func Slices[K constraints.Ordered, V any](key func(V) K, combiner func(*V, *V), slices ...[]V) []V {
	tm := treemap.New[K, V]()

	for i := range slices {
		for j := range slices[i] {
			k := key(slices[i][j])
			v, exists := tm.Get(k)
			if exists {
				combiner(&v, &slices[i][j])
				tm.Set(k, v)
			} else {
				tm.Set(k, slices[i][j])
			}
		}
	}

	var out []V
	for it := tm.Iterator(); it.Valid(); it.Next() {
		out = append(out, it.Value())
	}
	return out
}

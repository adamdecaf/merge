// Package merge provides functionality to merge multiple slices into a single sorted slice,
// combining values with the same key using a user-defined combiner function.
package merge

import (
	"github.com/igrmk/treemap/v2"
	"golang.org/x/exp/constraints"
)

// Slices merges multiple slices into a single sorted slice based on a key function.
// It uses a treemap to maintain sorted order and combines values with the same key
// using the provided combiner function. If no combiner is provided (nil), the first
// value encountered for a given key is kept, and subsequent values are ignored.
//
// The function is generic, allowing keys of any ordered type (K) and values of any type (V).
// The key function extracts a key from each value, and the combiner function (if provided)
// modifies the stored value based on a new value with the same key.
//
// Parameters:
//   - key: A function that extracts a comparable key of type K from a value of type V.
//   - combiner: A function that combines two values of type V, modifying the first value
//     based on the second. If nil, the first value for a key is retained.
//   - slices: Variable number of input slices of type V to merge.
//
// Returns:
//
//	A sorted slice containing the merged values, ordered by their keys.
func Slices[K constraints.Ordered, V any](key func(V) K, combiner func(*V, V), slices ...[]V) []V {
	tm := treemap.New[K, *V]()

	for _, sl := range slices {
		for i := range sl {
			v := sl[i]
			k := key(v)

			if ptr, found := tm.Get(k); found {
				if combiner != nil {
					combiner(ptr, v)
				}
			} else {
				ptr := new(V)
				*ptr = v
				tm.Set(k, ptr)
			}
		}
	}

	out := make([]V, 0, tm.Len())

	for it := tm.Iterator(); it.Valid(); it.Next() {
		out = append(out, *it.Value())
	}

	return out
}

package lib

import "iter"

// Pairs returns an iterator over successive pairs of values from seq.
func Pairs[V any](seq iter.Seq[V]) iter.Seq2[V, V] {
	return func(yield func(V, V) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v1, ok1 := next()
			if !ok1 {
				return
			}
			v2, ok2 := next()
			// If ok2 is false, v2 should be the
			// zero value; yield one last pair.
			if !yield(v1, v2) {
				return
			}
			if !ok2 {
				return
			}
		}
	}
}

func Filter[V any](seq iter.Seq[V], check func(V) bool) iter.Seq[V] {
    return func(yield func(V) bool) {
        for v := range seq {
            if !check(v) {
                continue
            }

            if !yield(v) {
                break
            }
        }
    }
}

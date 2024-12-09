package lib

// TODO: make generic
func Intersect(a, b map[int]struct{}) bool {
	if len(a) < len(b) {
		for key := range a {
			_, ok := b[key]
			if ok {
				return true
			}
		}
	} else {
		for key := range b {
			_, ok := a[key]
			if ok {
				return true
			}
		}
	}

	return false
}

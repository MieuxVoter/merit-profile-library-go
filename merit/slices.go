package merit

// contains is a shim for [slices.Contains] since we support Go 1.16.
// If we bumped to Go 1.18 we could use [slices.Contains] instead of this.
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

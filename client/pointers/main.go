package pointers

// String returns the address of a string
func String(v string) *string {
	return &v
}

// Int returns the address of an integer
func Int(v int) *int {
	return &v
}

// DereferenceStrings converts a string pointer slice to a string slice
func DereferenceStrings(s []*string) []string {
	n := make([]string, len(s))
	for i := range s {
		n[i] = *s[i]
	}
	return n
}

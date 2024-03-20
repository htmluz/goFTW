package iteration

func Repeat(a string, b int) string {
	var repeated string
	for i := 0; i < b; i++ {
		repeated += a
	}
	return repeated
}

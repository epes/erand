package erand

// PrintMinMax returns a string of length `width` that shows
// where the result ended between the min and max
func PrintMinMax(min int, max int, val int, width int) string {
	diff := float32(max - min)
	ratio := float32(val-min) / diff
	index := int(ratio * float32(width))

	chars := make([]rune, width)

	for i := range chars {
		if i == index {
			chars[i] = 'O'
		} else {
			chars[i] = '-'
		}
	}

	return string(chars)
}

package reverse_string

func ReverseString(input string) (output string) {
	rune_slice := []rune(input)
	var rune_final []rune
	for i := len(rune_slice) - 1; i >= 0; i-- {
		rune_final = append(rune_final, rune_slice[i])
	}
	output = string(rune_final)
	return output
}

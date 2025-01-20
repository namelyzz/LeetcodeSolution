package code

func preprocess(s string) string {
	sb := []rune{'^'}
	for _, ch := range s {
		sb = append(sb, '#', ch)
	}
	sb = append(sb, '#', '$')
	return string(sb)
}

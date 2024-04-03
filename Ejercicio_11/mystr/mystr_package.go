package mystr

func Concatenar(xs []string) string {
	sep := " "
	str := ""
	for _, x := range xs {
		str += x
		str += sep
	}
	return str
}

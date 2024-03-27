// Package perro ofrece una serie de funciones para adaptar nuestros datos a formato canino.
package perro

type humanYears int
type dogYears int

//Func Edad permite convertir la edad de un humano a edad perruna
func Edad(hy humanYears) dogYears {
	edadPerruna := hy * 7
	return dogYears(edadPerruna)
}

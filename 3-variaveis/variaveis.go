package main

import (
	"fmt"
)

var ruby, php, golang bool //a função var declara uma lista de variaveis
// func main() {
// 	var inteiro int
// 	fmt.Println(inteiro, ruby, php, golang)

// }

// -------------------------------------------------------------------------
var i, j int = 1, 2 //inicializando variaveis
// func main() {
// 	var c, python, java = true, false, "no!"
// 	fmt.Println(i, j, c, python, java)
// }

// declaração curta de variaveis
// :=
// só é possivel usar dentro de uma função
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

package main

import "fmt"

func add(x int, y int) int { //função recebe duas variaveis e retorna um resultado com um valor do mesmo tipo
	return x + y
}

// uma funçõa pode ter zero ou mais argumentos
// em go declaramos o tipo e depois a variavel

//quando dois parametros forem do mesmo tipo podemos omitir o tipo de todos menos o ultimo para encurtar o código
func soma(x, y int) int {
	return x + y
}

// função pode retornar qualquer numero de resultados
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //quando retorno esta sem argumentos, retorna os valores das variaveis atuais, isso é conhecido como retorno limpo
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(soma(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
}

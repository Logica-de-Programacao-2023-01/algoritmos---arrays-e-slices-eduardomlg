package main

import "fmt"

// Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.
func main() {
	var numeros = [3]int{1, 5, 3}
	soma := 0
	for i := 0; i < len(numeros); i++ {
		soma += numeros[i]
	}
	fmt.Print(soma)
}

package main

import "fmt"

// Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do
// Slice e os valores dos elementos. Em seguida, imprima o Slice e a soma dos valores armazenados.
func main() {
	var tamanho int
	fmt.Println("Informe o tamanho da lista ")
	fmt.Scan(&tamanho)

	slice := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Println("Informe os numeros da lista? ")
		fmt.Scan(&slice[i])
	}
	fmt.Println(slice)
	soma := 0
	for i := 0; i < len(slice); i++ {
		soma += slice[i]
	}
	fmt.Print(soma)
}

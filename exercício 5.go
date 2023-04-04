package main

import "fmt"

func main() {
	var matriz [3][2]int
	for linha := 0; linha < len(matriz); linha++ {
		for coluna := 0; coluna < len(matriz[linha]); coluna++ {
			var x int
			fmt.Println("Digite o número da linha ", linha, " e da coluna ", coluna)
			fmt.Scan(&x)
			matriz[linha][coluna] = x
		}
	}
}

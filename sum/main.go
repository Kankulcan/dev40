package main

import (
	"fmt"
)

func main() {
	nome := "Jose"
	i := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(sum(i...))
	fmt.Println(mostraNome(nome))
}

func mostraNome(n string) string {
	return fmt.Sprintf("Seja muito bem vindo %s", n)

}

func sum(i ...int) int {
	soma := 0
	for _, e := range i {
		soma += e
	}
	return soma
}

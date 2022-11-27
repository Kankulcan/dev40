package main

import (
	"fmt"
)

func main() {
}

func division(i float64, j float64) (float64, error) {
	if j == 0 {
		return 0, fmt.Errorf("não é possĩvel realizar uma divisão por zero!")
	}

	if i < 0 || j < 0 {
		return 0, fmt.Errorf("não aceitamos a divisão com números negativos!")
	}

	return i / j, nil

}

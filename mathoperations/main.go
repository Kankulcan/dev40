package main

import "errors"

func main() {

}

func somar(a float64, b float64) float64 {
	return a + b
}

func dividir(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("não é possível dividir por zero!")
	}

	return a / b, nil
}

func multiplicar(a float64, b float64) float64 {
	return a * b
}

func subtrair(a float64, b float64) float64 {
	return a - b
}

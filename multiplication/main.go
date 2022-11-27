package main

func main() {

}

func multiplication(i int, j int) int {
	return i * j
}

func fatorial(i int) int {
	fator := 1
	for j := 1; j <= i; j++ {
		fator *= j
	}

	return fator
}

func printNome(n string) string {
	return ("Seja bem-vindo " + n)
}

package main

import "testing"

func Test_sum(t *testing.T) {
	number := []struct{
		number []int
		expected int
	}{
		{number:[]int{1,2,3,4}, expected:10},
	}


	for _,e := range number{
		somaTest := sum(e.number...)
		
		if somaTest != e.expected {
			t.Errorf("Esperávamos receber %d mas recebemos %d", e.expected, somaTest)
		}

	}
}

func Test_mostraNome(t *testing.T){
	nomes := []struct {
		nome string
		expected string
	}{
		{nome: "Lionel Filho", expected: "Seja muito bem vindo Lionel Filho"},
		{nome: "Acácia Amarela", expected: "Seja muito bem vindo Acácia Amarela"},
		{nome: "@323as **", expected: "Seja muito bem vindo @323as **"},
		{nome: "Anúcio de Silva Sá", expected: "Seja muito bem vindo Anúcio de Silva Sá"},

	}

	for _,e := range nomes{
		resultado := mostraNome(e.nome)
		if resultado != ("Seja muito bem vindo "+ e.nome){
			t.Errorf("A saudação esperada era - Seja muito bem vindo %s", e.nome)
		}
	}
}
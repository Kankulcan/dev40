package main

import (
	"testing"
)

func Test_multipication(t *testing.T) {
	gotValues := []struct {
		value  int
		value2 int
		expect int
	}{
		{value: 4, value2: 3, expect: 12},
		{value: 2, value2: 5, expect: 10},
		{value: 10, value2: 7, expect: 70},
		{value: -3, value2: 2, expect: -6},
	}

	for _, val := range gotValues {
		result := multiplication(val.value, val.value2)
		if result != val.expect {
			t.Errorf("O valor esperado era %d, mas o encontrado foi %d", val.expect, result)
		}
	}

}

func Test_fatorial(t *testing.T) {
	gotValues := []struct {
		value  int
		expect int
	}{
		{value: 4, expect: 24},
		{value: 3, expect: 6},
		{value: 5, expect: 120},
		{value: 2, expect: 2},
		{value: 6, expect: 720},
	}

	for _, val := range gotValues {
		result := fatorial(val.value)

		if result != val.expect {
			t.Errorf("O valor esperado era %d, mas o valor encontrado foi %d", val.expect, result)
		}
	}
}

func Test_printNome(t *testing.T) {
	gotValues := []struct {
		name string
		expect string
	}{
		{name: "", expect: "Seja bem-vindo "},
		{name: "Juliana", expect: "Seja bem-vindo Juliana"},
		{name: "Catarina", expect: "Seja bem-vindo Catarina"},
		{name: "Arnold Fernandes Sauro", expect: "Seja bem-vindo Arnold Fernandes Sauro"},
	}

	for _,val := range gotValues{
		result := printNome(val.name)
		if result != val.expect {
			t.Errorf("O valor esperado era - %s", val.expect)
			t.Errorf("O valor recebido foi - %s", result)
		}
	}
}

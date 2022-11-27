package main

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func Test_division(t *testing.T) {
	t.Run("Testing division by zero", func(t *testing.T) {

		gotTests := []struct {
			number   float64
			divisor  float64
			expected float64
			erro     error
		}{
			{number: 46, divisor: 0, expected: 0, erro: errors.New("não é possĩvel realizar uma divisão por zero!")},
		}
	
		for _, value := range gotTests {
			result, err := division(value.number, value.divisor)
	
			if !reflect.DeepEqual(err.Error(),value.erro.Error()) {
				t.Errorf("esperávamos receber - %s", value.erro.Error())
				t.Errorf("recebemos - %s", err.Error())
			}
	
			if result != value.expected {
				t.Errorf("esperávamos %0.2f, mas recebemos %0.2f", value.expected, result)
			}
		}
	})

	t.Run("testing division by a negative number", func(t *testing.T) {
		erroMsg:= "não aceitamos a divisão com números negativos!"

		gotTests := []struct {
			number   float64
			divisor  float64
			expected float64
			erro     error
		}{
			{number: -4, divisor: 2, expected: 0, erro: errors.New(erroMsg)},
			{number: 23, divisor: -3, expected: 0, erro: errors.New(erroMsg)},
			{number: 234, divisor: -5, expected: 0, erro: errors.New(erroMsg)},
			{number: -6345, divisor: -10, expected: 0, erro: errors.New(erroMsg)},
		}
	
		for _, value := range gotTests {
			result, err := division(value.number, value.divisor)
	
			if strings.Compare(err.Error(),value.erro.Error()) != 0 {
				t.Errorf("esperávamos receber - %s", value.erro.Error())
				t.Errorf("recebemos - %s", err.Error())
			}
	
			if result != value.expected {
				t.Errorf("esperávamos %0.2f, mas recebemos %0.2f", value.expected, result)
			}
		}

	})
	
}


package main

import (
	"errors"
	"testing"
)

func TestSomar(t *testing.T) {
	got := somar(1, 2)
	want := float64(3)
	assertCorrectMessage(t, got, want)

}

func TestDividir(t *testing.T) {

	t.Run("Dividindo tudo por zero", func(t *testing.T) {
		got, err := dividir(2, 2)
		want := float64(1)
		wanterr := errors.New("não é possível dividir por zero!")
		assertCorrectMessageDivision(t, got, want, wanterr, err)

	})

	got, err := dividir(2, 2)
	want := float64(1)
	wanterr := errors.New("")
	assertCorrectMessageDivision(t, got, want, wanterr, err)
}

func TestMultiplicar(t *testing.T) {
	got := multiplicar(2, 2)
	want := float64(4)
	assertCorrectMessage(t, got, want)
}

func TestSubtrair(t *testing.T) {
	got := subtrair(3, 2)
	want := float64(1)
	assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("Expected: %0.2f,  Received: %0.2f", want, got)
	}
}

func assertCorrectMessageDivision(t testing.TB, got float64, want float64, wanterr error, err error) {
	t.Helper()
	if err != nil {
		strError := err.Error()
		if strError != wanterr.Error() {
			t.Errorf("erro esperado: %s - erro recebido: %s ", wanterr.Error(), strError)
		}
	}
	if got != want {
		t.Errorf("Expected: %0.2f,  Received: %0.2f", want, got)
	}
}

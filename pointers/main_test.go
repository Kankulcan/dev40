package main

import (
	"testing"
)

func Test_wallet(t *testing.T) {
	//É muito importante entedermos o que está se passando e por que precisamos utilizar 
	//ponteiros para por receber os valores da Wallet.
	//1. Criamos a Wallet sem definição de valores. Então ela está zerada.
	//Cada função ou método chamado, faz criar uma nova cópia dos parâmetros, portanto
	//quando chamamos Deposit ou Balance, estamos criando cópias de novos itens.
	//Para resolvermos esta situação, devemos aplicar o ponteiro nos métodos e funções.
	//para que possam instanciar o mesmo objeto já na memória.

	//Se quisermos que uma função mude o estado do objeto, precisamos passá-la com ponteiros.
	
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}



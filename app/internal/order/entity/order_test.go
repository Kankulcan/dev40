package entity

import (
	"fmt"
	"testing"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assertOrder(t, order.IsValid(), "the order must have an ID")
}

func TestGivenAnZeroPrice_WhenCreateANewOrder_ThenShouldReceivedAnError(t *testing.T) {
	order := Order{ID: "Camisa"}
	assertOrder(t, order.IsValid(), "the price must be more then zero")
}

func TestGivenAnZeroTax_WhenCreateANewOrder_ThenShouldReceivedAnError(t *testing.T) {
	order := Order{ID: "Camisa", Price: 100}
	assertOrder(t, order.IsValid(), "the tax should be more then zero")
}

func TestGivenACorrectParams_WhenCreateAnOrder_NoErrosExpected(t *testing.T) {
	order, err := NewOrder("Camisa",100.0, 1.00)
	if err != nil { 
		t.Error("O programa está errado.")
	}

	assert(t, order.ID, "Camisa")
	assert(t, fmt.Sprint(order.Price), "100") 
	assert(t, fmt.Sprint(order.Tax), "1")
}

func TestGivenPriceAndTax_WhenCreateAnOrder_CalculateFinalPrice(t *testing.T) {
	order, err := NewOrder("Camisa",100.0, 2.00)
	if err != nil { 
		t.Error("O programa está errado.")
	}
	
	err = order.CalculatePrice()
	if err != nil { 
		t.Error("O programa está errado.")
	}

	assert(t, order.ID, "Camisa")
	assert(t, fmt.Sprint(order.Price), "100") 
	assert(t, fmt.Sprint(order.Tax), "2")
	assert(t,fmt.Sprint(order.FinalPrice), "102") 

}

func assert(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s ", got, want)
	}
}

func assertOrder(t *testing.T, valid error, want string) {
	t.Helper()

	if valid != nil {

		got := valid.Error()
		if got != want {
			t.Errorf("got %s, want %s ", got, want)
		}
	}

}

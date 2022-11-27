package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price, tax float64) (*Order, error) {
	order :=  &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.IsValid()
	if err!= nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) CalculatePrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err!= nil {
		return err
	}
	return nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New("the order must have an ID")
	}

	if o.Price <= 0 {
		return errors.New("the price must be more then zero")
	}

	if o.Tax < 0 {
		return errors.New("the tax should be more then zero")
	}

	return nil
}

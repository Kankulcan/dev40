package main

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("Rectangle ", func(t *testing.T) {
		//Esse objeto irá guardar o valor por que está sendo instanciado já com o valor
		//Se utilizássemos um método para criar o objeto, só teríamos acesso ao objeto 
		//criado se utilizássemos um ponteiro. Isso por que cada método copia um novo
		//parâmetro.
		rectangle := Rectangle{10.0, 10.0} 
		got := perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got: %.2f  want: %.2f", got, want)
		}
	})

}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 12, height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 12, height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}

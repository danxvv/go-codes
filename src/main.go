package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadro struct {
	base float64
}

type rectangulo struct {
	width  float64
	height float64
}

func (s cuadro) area() float64 {
	return s.base * s.base
}

func (r rectangulo) area() float64 {
	return r.width * r.height
}

func calculate(f figuras2D) {
	fmt.Println("Area", f.area())
}

func main() {
	myCuadro := cuadro{base: 2.0}
	myRect := rectangulo{height: 3.4, width: 2.1}

	calculate(myCuadro)
	calculate(myRect)

	// Lista de interface
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)

}

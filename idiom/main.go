package main

import "fmt"

type TransmissionType int

const (
	TransmissionUndefined TransmissionType = iota
	TransmissionManual
	TransmissionAutomatic
)

type Car struct {
	brand        string
	color        string
	serialNumber string
	transmission TransmissionType
}

// NewCar implements fixed number of parameters, which is a pain to use, and also
// to extend later
func NewCar(serialNumber, brand, color string,
	transmissionType TransmissionType) *Car {
	return &Car{
		brand:        brand,
		color:        color,
		serialNumber: serialNumber,
		transmission: transmissionType,
	}
}

func main() {
	car := NewCar("123", "BMW", "Red", TransmissionAutomatic)
	fmt.Println(car)
}

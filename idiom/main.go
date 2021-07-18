package main

import "fmt"

type TransmissionType int

func (tt TransmissionType) String() string {
	switch tt {
	case TransmissionAutomatic:
		return "Automatic"
	case TransmissionManual:
		return "Manual"
	default:
		return ""
	}
}

const (
	TransmissionUndefined TransmissionType = iota
	TransmissionManual
	TransmissionAutomatic
)

type CarOptions struct {
	brand        string
	color        string
	transmission TransmissionType
}

type CarOption func(*CarOptions)

type Car struct {
	options      CarOptions
	serialNumber string
}

func (c *Car) String() string {
	return fmt.Sprintf("brand: %s color: %s serial: %s transmission: %s",
		c.options.brand, c.options.color, c.serialNumber, c.options.transmission)
}

// NewCar implements variable number of options as parameters
func NewCar(serialNumber string, options ...CarOption) *Car {
	args := CarOptions{}
	for _, option := range options {
		option(&args)
	}
	return &Car{
		options:      args,
		serialNumber: serialNumber,
	}
}

func WithBrand(brand string) CarOption {
	return func(o *CarOptions) {
		o.brand = brand
	}
}

func WithColor(color string) CarOption {
	return func(o *CarOptions) {
		o.color = color
	}
}

func WithTransmissionType(tt TransmissionType) CarOption {
	return func(o *CarOptions) {
		o.transmission = tt
	}
}

func main() {
	car := NewCar("123")
	fmt.Println(car)

	car = NewCar("456",
		WithBrand("BMW"),
		WithTransmissionType(TransmissionAutomatic),
	)
	fmt.Println(car)

	car = NewCar("789",
		WithBrand("BMW"),
		WithColor("Red"),
		WithTransmissionType(TransmissionAutomatic),
	)
	fmt.Println(car)
}

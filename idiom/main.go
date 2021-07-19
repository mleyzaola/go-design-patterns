package main

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	TransmissionType int
	SpeedUnitType    int
)

const (
	KmUnit SpeedUnitType = iota + 1
	MiUnit
)

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

func (ut SpeedUnitType) String() string {
	switch ut {
	case KmUnit:
		return "Km/hr"
	case MiUnit:
		return "Mi/hr"
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
	maxSpeed     float64
	maxSpeedType SpeedUnitType
	transmission TransmissionType
}

type CarOption func(*CarOptions)

type Car struct {
	options      CarOptions
	serialNumber string
}

func (c *Car) String() string {
	var values []string
	if val := c.options.brand; val != "" {
		values = append(values, val)
	}
	if val := c.options.color; val != "" {
		values = append(values, val)
	}
	if val := c.options.transmission; val != TransmissionUndefined {
		values = append(values, val.String())
	}
	if val := c.options.maxSpeed; val != 0 {
		values = append(values,
			strconv.FormatFloat(val, 'f', 0, 64)+
				" "+
				c.options.maxSpeedType.String(),
		)
	}
	return strings.Join(values, " ")
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

func WithMaxSpeed(maxSpeed float64, unitType SpeedUnitType) CarOption {
	return func(o *CarOptions) {
		o.maxSpeed = maxSpeed
		o.maxSpeedType = unitType
	}
}

// TODO: let the audience implement maxSpeedUnit of type Unit and support Km and Mi
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
		WithMaxSpeed(330, KmUnit),
	)
	fmt.Println(car)
}

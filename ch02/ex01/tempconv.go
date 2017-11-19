// Package tempconv performs Celsius, Fahrenheit and Kelvin conversions.
package tempconv

import "fmt"

type (
	// Celsius is a temperature in ºC
	Celsius float64
	// Fahrenheit is a temperature in ºF
	Fahrenheit float64
	// Kelvin is a temperature in ºK
	Kelvin float64
)

// Reference temperatures
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

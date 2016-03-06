// Package conv defines various units and conversions.
package conv

import "fmt"

type (
	Celsius    float64
	Fahrenheit float64
	Meters     float64
	Feet       float64
	Kilograms  float64
	Pounds     float64
	Hectares   float64
	Acres      float64
)

func (c Celsius) String() string    { return fmt.Sprintf("%f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%f°F", f) }
func CToF(c Celsius) Fahrenheit     { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius     { return Celsius((f - 32) * 5 / 9) }

func (m Meters) String() string { return fmt.Sprintf("%f m", m) }
func (f Feet) String() string   { return fmt.Sprintf("%f ft", f) }
func MToF(m Meters) Feet        { return Feet(m / 0.3048) }
func FToM(f Feet) Meters        { return Meters(f * 0.3048) }

func (k Kilograms) String() string { return fmt.Sprintf("%f kg", k) }
func (p Pounds) String() string    { return fmt.Sprintf("%f lbs", p) }
func KToP(k Kilograms) Pounds      { return Pounds(k * 1000 / 453.59237) }
func PToK(p Pounds) Kilograms      { return Kilograms(p * 453.59237 / 1000) }

func (h Hectares) String() string { return fmt.Sprintf("%f hectares", h) }
func (a Acres) String() string    { return fmt.Sprintf("%f acres", a) }
func HToA(h Hectares) Acres       { return Acres(h * 2.4710538148) }
func AToH(a Acres) Hectares       { return Hectares(a / 2.4710538148) }

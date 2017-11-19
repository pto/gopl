// Package conv defines various units and conversions.
package conv

import "fmt"

type (
	// Celsius is a temperature in ºC.
	Celsius float64
	// Fahrenheit is a temperature in ºF.
	Fahrenheit float64
	// Meters is a length in meters.
	Meters float64
	// Feet is a length in feet.
	Feet float64
	// Kilograms is a weight in kilograms.
	Kilograms float64
	// Pounds is a weight in pounds.
	Pounds float64
	// Hectares is an area in hectares.
	Hectares float64
	// Acres is an area in acres.
	Acres float64
	// Gallons is a volumne in gallons.
	Gallons float64
	// Liters is a volume in liters.
	Liters float64
)

func (c Celsius) String() string    { return fmt.Sprintf("%.6g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.6g°F", f) }

// CToF converts a temperature in Celsius to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a temperature in Fahrenheit to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (m Meters) String() string { return fmt.Sprintf("%.6g m", m) }
func (f Feet) String() string   { return fmt.Sprintf("%.6g ft", f) }

// MToF converts a distance in meters to feet.
func MToF(m Meters) Feet { return Feet(m / 0.3048) }

// FToM converts a distance in feet to meters.
func FToM(f Feet) Meters { return Meters(f * 0.3048) }

func (k Kilograms) String() string { return fmt.Sprintf("%.6g kg", k) }
func (p Pounds) String() string    { return fmt.Sprintf("%.6g lbs", p) }

// KToP converts a weight from kilograms to pounds.
func KToP(k Kilograms) Pounds { return Pounds(k * 1000 / 453.59237) }

// PToK converts a weight from pounds to kilgrams.
func PToK(p Pounds) Kilograms { return Kilograms(p * 453.59237 / 1000) }

func (h Hectares) String() string { return fmt.Sprintf("%.6g hectares", h) }
func (a Acres) String() string    { return fmt.Sprintf("%.6g acres", a) }

// HToA converts an area from hectares to acres.
func HToA(h Hectares) Acres { return Acres(h * 2.4710538148) }

// AToH converts an area from acres to hectares.
func AToH(a Acres) Hectares { return Hectares(a / 2.4710538148) }

func (g Gallons) String() string { return fmt.Sprintf("%.6g gallons", g) }
func (l Liters) String() string  { return fmt.Sprintf("%.6g liters", l) }

// GToL converts a volume from gallons to liters.
func GToL(g Gallons) Liters { return Liters(g * 3.78541) }

// LToG converts a volume from liters to gallons.
func LToG(l Liters) Gallons { return Gallons(l / 3.78541) }

package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Feet float64
type Meters float64
type Pound float64
type KiloGram float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Feet) String() string       { return fmt.Sprintf("%gft", f) }
func (m Meters) String() string     { return fmt.Sprintf("%gm", m) }
func (p Pound) String() string      { return fmt.Sprintf("%glb", p) }
func (k KiloGram) String() string   { return fmt.Sprintf("%gkg", k) }

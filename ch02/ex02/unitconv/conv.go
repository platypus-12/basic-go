package unitconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func FToM(f Feet) Meters        { return Meters(f / 3.2808) }
func MToF(m Meters) Feet        { return Feet(m * 3.2808) }
func PToK(p Pound) KiloGram     { return KiloGram(p / 2.2046) }
func KToP(k KiloGram) Pound     { return Pound(2.2046 * k) }

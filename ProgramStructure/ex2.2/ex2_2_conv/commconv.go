// Package commconv 一个通用的单位换算，包括了米和英尺以及千克和磅
package commconv

import "fmt"

type Meter float64
type Feet float64
type Kilogram float64
type Pound float64

func (ft Feet) String() string {
	return fmt.Sprintf("%g ft", ft)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kg", k)
}

func (l Pound) String() string {
	return fmt.Sprintf("%g lb", l)
}

func FtoM(ft Feet) Meter {
	return Meter(ft * 0.3048)
}

func MtoF(m Meter) Feet {
	return Feet(m / 0.3048)
}

func KtoP(k Kilogram) Pound {
	return Pound(k / 0.45359237)
}

func PtoK(p Pound) Kilogram {
	return Kilogram(p * 0.45359237)
}

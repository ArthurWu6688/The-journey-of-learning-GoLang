// Package tempconv 这是一个包注释，这个包的功能是带有单位的温度输出
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelwin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

const (
	AbsoluteZeroK Kelwin = 0
	FreezingK            = 273.15
	BoilingK             = 373.15
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g ˚C\n", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g ˚F\n", f)
}

func (k Kelwin) String() string {
	return fmt.Sprintf("%g K\n", k)
}

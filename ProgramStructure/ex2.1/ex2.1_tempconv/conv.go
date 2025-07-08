// Package tempconv 这是一个包注释，这个包的功能是温度转换
package tempconv

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CtoK(c Celsius) Kelwin {
	return Kelwin(c - 273.15)
}

func KtoC(k Kelwin) Celsius {
	return Celsius(k + 273.15)
}

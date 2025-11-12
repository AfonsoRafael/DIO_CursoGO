package main

import "fmt"

const ebolicao_kelvin float64 = 373.00
func main() {
var tempk float64 = ebolicao_kelvin
tempC := tempk - 273.00
fmt.Printf("A temperatura de ebulição da água em Kelvin é: %2.f K e em Celsius é %2.f", tempk, tempC)
}
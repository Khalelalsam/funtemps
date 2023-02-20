package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius

//////////Fahrenheit
func FahrenheitToCelsius(fahrenheit float64) float64 {

  celsius := (fahrenheit - 32) * 5/9


	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return celsius
}

func FahrenheitToKevlin(fahrenheit float64) float64 {

  Kelvin := (fahrenheit - 32) * 5/9 + 273.15
 
	return Kelvin
}
////cELSIUS
func CelsiusToFahrenheit(Celsius float64) float64 {

  fahrenheit := Celsius* 9/5 + 32
 
	return fahrenheit
}

func CelsiusToKelvin(Celsius float64) float64 {
  Kelvin := Celsius + 273.15
  
	return Kelvin
}
//kELVIN
func KelvinToCelsius(Kelvin float64) float64 {
  Celsius := Kelvin - 273.15
 
	return Celsius
}

func KelvinToFahrenheit(Kelvin float64) float64 {
  fahrenheit := (Kelvin - 273.15)* 9/5 + 32

	return  fahrenheit

}




//func FarhenheitToCelsius(fahrenheit float64) float64 {}
// De andre konverteringsfunksjonene implementere her
// ...

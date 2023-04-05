package main

import (
	"flag"
	"fmt"
	"funtemps/conv"
	"math"
)

// Definerer flag-variablene i hoved-"scope"
var cels float64
var kel float64
var temp string
var fahr float64
var out string
var funfacts string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader Celsius")
	flag.Float64Var(&kel, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&temp, "t", "C", "temperatur i sun")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")

}

func main() {

	flag.Parse()

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen--------------------------------------------
	//fmt.Println(fahr, out, funfacts)

	//fmt.Println("len(flag.Args())", len(flag.Args()))
	//fmt.Println("flag.NFlag()", flag.NFlag())

	//fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	if out == "C" && isFlagPassed("F") {

		celsius := conv.FahrenheitToCelsius(fahr)

		if celsius == math.Trunc(celsius) {

			celsiusInt := int(celsius)
			fmt.Printf("%.2f °F er %d °C.", fahr, celsiusInt)
		} else {
			fmt.Printf("%.2f °F er %.2f °C.", fahr, celsius)
		}

	} else if out == "C" && isFlagPassed("K") {

		Celsius := conv.KelvinToCelsius(kel)

		if Celsius == math.Trunc(Celsius) {

			Int := int(Celsius)

			fmt.Printf("%.2f  °K er %.d °C.", kel, Int)
		} else {
			fmt.Printf("%9.2f °K er %9.2f °C.", kel, Celsius)
		}

	} else if out == "F" && isFlagPassed("C") {

		f := conv.CelsiusToFahrenheit(cels)

		if f == math.Trunc(f) {

			fInt := int(f)
			fmt.Printf("%.2f °C er %.d °F.", cels, fInt)
		} else {
			fmt.Printf("%.2f °C er %.2f °F.", cels, f)
		}

	} else if out == "F" && isFlagPassed("K") {

		Fahrenheit := conv.KelvinToFahrenheit(kel)

		if Fahrenheit == math.Trunc(Fahrenheit) {

			Int := int(Fahrenheit)
			fmt.Printf("%9.2f °K er %.d °F.", kel, Int)
		} else {
			fmt.Printf("%9.2f °K er %9.2f °F.", kel, Fahrenheit)
		}

		//fmt.Printf("%.2f degrees kelvin is equal to %.2f degrees Farhrenheit.",kel,conv.KelvinToFahrenheit(kel))

	} else if out == "K" && isFlagPassed("C") {

		Kelvin := conv.CelsiusToKelvin(cels)

		if Kelvin == math.Trunc(Kelvin) {

			Int := int(Kelvin)

			fmt.Printf("%9.2f °C er %.d °K.", cels, Int)
		} else {
			fmt.Printf("%9.2f °C er %9.2f °K.", cels, Kelvin)
		}

	} else if out == "K" && isFlagPassed("F") {

		Kelvin := conv.FahrenheitToKevlin(fahr)

		if Kelvin == math.Trunc(Kelvin) {

			Int := int(Kelvin)

			fmt.Printf("%9.2f °F er %.d °K.", fahr, Int)
		} else {
			fmt.Printf("%9.2f °F er %9.2f °K.", fahr, Kelvin)
		}
	}

	if temp == "C" && isFlagPassed("funfacts") && funfacts != "terra" && funfacts != "luna" {

		fmt.Printf("Temperatur på ytre lag av Solen 5506.85°C.\n")
		fmt.Printf("Temperatur i Solens kjerne er 15 000 000°C.")

	} else if funfacts == "luna" && isFlagPassed("funfacts") {

		fmt.Printf("Temperatur på Månens overflate om natten er lik -183\n")
		fmt.Printf("Temperatur på Månens overflate om dagen er lik 106 ")

	} else if funfacts == "terra" && isFlagPassed("funfacts") {

		fmt.Printf("Laveste temperatur målt på Jordens overflate -89.4\n")
		fmt.Printf("Temperatur i Jordens indre kjerne 9392 ")
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

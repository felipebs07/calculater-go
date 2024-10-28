package main

import (
	"fmt"
	"strconv"
)

func main() {
	var quantity_notes int
	fmt.Print("Hello teacher, how many notes u will add? ")

	fmt.Scan(&quantity_notes)

	var sumTotal float64
	sumTotal = 0.0
	for i := 0; i < quantity_notes; i++ {
		var note float64
		fmt.Print("What value of note? ")
		fmt.Scan(&note)
		sumTotal = sumTotal + note
	}

	averageNote := sumTotal / float64(quantity_notes)
	println("The average is " + strconv.FormatFloat(averageNote, 'f', 2, 32))
}

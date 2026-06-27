package main

import (
	"fmt"
	"strings"
)

func main() {

	// Exercise 1 — Variables & String Format: My name is Asher, I am 20 years old, and it is true that I am a student.
	name := "Asher"
	age := 20

	fmt.Printf("My name is %s, I am %d years old, and it is true that I am a student. \n", name, age)

	// Exercise 2 — If/Else: Write a function called classify that takes an int num and returns if the num is negative, positive, or zero
	fmt.Println(classify(0))
	fmt.Println(classify(-10))
	fmt.Println(classify(10))

	// Exercise 3 — Switch: Write a function called dayType that takes a day name as a string and return if its weekday, weekend, or unknown

	if result, err := dayType("SUNDAY"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Type of day:", result)
	}

	// Exercise 4 — For Loop: Write a loop that prints the numbers 1 to 10, also tell if the number is even or odd
	// Exercise 5 — Range: Create a slice of 5 numbers. Loop through it using range and print the sum at the end.
}

func classify(num int) string {
	switch {
	case num < 0:
		return "negative"
	case num > 0:
		return "positive"
	default:
		return "zero"
	}
}

func dayType(day string) (string, error) {
	switch strings.TrimSpace(strings.ToLower(day)) {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		return "weekday", nil
	case "saturday", "sunday":
		return "weekend", nil
	default:
		return "unknown", fmt.Errorf("day unknown")
	}
}

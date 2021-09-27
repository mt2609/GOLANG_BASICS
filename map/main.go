package main

import "fmt"

func main() {

	color := map[string]string{
		"red":    "apple",
		"yellow": "banana",
		"white":  "peach",
	}
	pM(color)
}
func pM(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code", color, "is", hex)
	}
}

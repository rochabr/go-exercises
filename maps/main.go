package main

import "fmt"

func main() {

	//var colors map[string]string
	//colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	colors["white"] = "#ffffff"
	colors["whitey"] = "#ffffffa"

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%+v - %+v", color, hex)
		fmt.Println("")
	}
}

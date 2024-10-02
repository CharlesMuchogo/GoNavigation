package main

import (
	"github.com/CharlesMuchogo/GoNavigation/cmd"
	"github.com/CharlesMuchogo/GoNavigation/samples"
	"log"
)

func main() {

	navigation := &cmd.Navigator
	navigation.Navigate(samples.HomePage)

	for {
		cmd.ClearScreen()
		screen, err := navigation.Top()
		if err != nil {
			log.Fatal(err)
		}

		if navigation.IsEmpty() {
			break
		}
		screen()
	}

}

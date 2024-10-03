package navigation

import (
	"github.com/CharlesMuchogo/GoNavigation/cmd"
	"log"
)

var Navigator = cmd.NavStack{}

func InitialRoute(route func()) {
	navigation := &Navigator
	navigation.Navigate(route)

	for {
		cmd.ClearScreen()
		screen, err := navigation.Top()
		if navigation.IsEmpty() {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		screen()
	}
}

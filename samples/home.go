package samples

import (
	"Navigation/cmd"
	"github.com/charmbracelet/huh"
	"log"
)

var (
	selectedOption string
)

func HomePage() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Main Menu").Options(
				huh.NewOption("Screen 1", "1"),
				huh.NewOption("Screen 2", "2"),
				huh.NewOption("Screen 3", "3"),
				huh.NewOption("Exit", "4"),
			).Value(&selectedOption),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if selectedOption != "4" {
		cmd.Navigator.Navigate(func() {
			DescriptionPage(selectedOption)
		})
	} else {
		cmd.Navigator.Pop()
	}
}

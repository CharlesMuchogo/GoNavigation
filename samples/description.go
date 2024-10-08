package samples

import (
	"fmt"
	"github.com/CharlesMuchogo/GoNavigation/navigation"
	"github.com/charmbracelet/huh"
	"log"
)

func DescriptionPage(page string) {
	title := fmt.Sprintf("Welcome to page %s", page)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title(title).Options(
				huh.NewOption("Go back", "5"),
				huh.NewOption("Description 6", "6"),
				huh.NewOption("Description 7", "7"),
				huh.NewOption("Description 8", "8"),
			).Value(&selectedOption),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if selectedOption == "5" {
		navigation.Navigator.Pop()
	} else {
		navigation.Navigator.Navigate(func() {
			MoreDescriptionPage(selectedOption)
		})
	}
}

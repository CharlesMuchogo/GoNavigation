package samples

import (
	"fmt"
	"github.com/CharlesMuchogo/GoNavigation/navigation"
	"github.com/charmbracelet/huh"
	"log"
)

func MoreDescriptionPage(page string) {
	title := fmt.Sprintf("Welcome to more description for page %s", page)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title(title).Options(
				huh.NewOption("Go back", "5"),
			).Value(&selectedOption),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if selectedOption == "5" {
		navigation.Navigator.Pop()
	}
}

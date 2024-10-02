package samples

import (
	"fmt"
	"github.com/CharlesMuchogo/GoNavigation/cmd"
	"github.com/charmbracelet/huh"
	"log"
)

func DescriptionPage(page string) {
	title := fmt.Sprintf("Welcome to page %s", page)

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
		cmd.Navigator.Pop()
	}
}

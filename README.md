# GoNavigation Library

## Overview

`GoNavigation` is a simple navigation library built using Go for terminal-based applications. It allows for easy page navigation with a stack-based approach, enabling you to push, pop, and navigate between different screens.

The library includes:
- **Navigation Stack**: A stack-based system to handle the screen/page navigation.
- **Screen Clearing**: Cross-platform support to clear the terminal screen (Linux, macOS, and Windows).
- **Sample Pages**: Includes sample pages (`HomePage`, `DescriptionPage`) using the `huh` library for form-based UI.

## How to Use

- Set Up Initial Route: In the main.go, use navigation.InitialRoute() to set the initial screen.



```go
package main

import (
	"github.com/CharlesMuchogo/GoNavigation/navigation"
	"github.com/CharlesMuchogo/GoNavigation/samples"
)

func main() {
	navigation.InitialRoute(samples.HomePage)
}
```



- Navigate Between Pages: Use Navigate() to push a new page onto the stack, and Pop() to return to the previous page.

  

```go
package sample

import (
	"github.com/CharlesMuchogo/GoNavigation/navigation"
	"github.com/charmbracelet/huh"
	"log"
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
	// Navigates to the selected screen
	if selectedOption != "4" {
		navigation.Navigator.Navigate(func() {
			DescriptionPage(selectedOption)
		})
	} else {
      // Navigates back
		navigation.Navigator.Pop()
	}
}
```


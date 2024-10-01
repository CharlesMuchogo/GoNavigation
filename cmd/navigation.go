package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

var Navigator = NavStack{}

type NavStack struct {
	items []func()
	mu    sync.Mutex
}

func (s *NavStack) Navigate(page func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, page)
}

func (s *NavStack) Pop() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *NavStack) Top() (func(), error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.IsEmpty() {
		return nil, fmt.Errorf("no page found in the stack")
	}
	return s.items[len(s.items)-1], nil
}

func (s *NavStack) IsEmpty() bool {
	return len(s.items) == 0
}

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") // Darwin (macOS)
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ClearScreen() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {
		value() //we execute it
	}
}

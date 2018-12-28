package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Windows OS")
	case "linux":
		fmt.Println("Linux Distribution")
	default:
		fmt.Println("OS: ", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")

	}
}

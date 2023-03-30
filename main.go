package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 44
	height = 44
)

var (
	colors = []string{
		"\x1b[31m", // red
		"\x1b[32m", // green
		"\x1b[33m", // yellow
		"\x1b[34m", // blue
		"\x1b[35m", // magenta
		"\x1b[36m", // cyan
	}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		clearScreen()

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				char := getChar()
				color := colors[rand.Intn(len(colors))]

				fmt.Printf("%s%s", color, char)
			}

			fmt.Printf("\x1b[0m\n") // reset color and move to next line
		}

		time.Sleep(62 * time.Millisecond) // refresh rate 16hz
	}
}

func clearScreen() {
	fmt.Print("\x1b[2J\x1b[H") // clear screen and move cursor to top-left
}

func getChar() string {
	chars := []string{
		" ", "　", "•", "∙", "●", "◦", "◉", "◎", "⚫", "⚪", "❖", "✦", "✶", "✹", "✿", "❀", "❂", "❈", "❉", "❄", "❅", "❇", "❦", "❧", "❣", "❤", "❥", "❦", "❧", "❞", "❝", "❍", "❑", "❒", "⌘", "⌥", "⌃", "⇧", "⇪", "⌫", "⌦", "⏎", "↩", "⇥", "⇤", "⎋",
	}

	return chars[rand.Intn(len(chars))]
}

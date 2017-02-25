package main

import (
	"fmt"

	"bufio"
	"os"

	"os/exec"

	"strings"

	"github.com/miquella/ask"
)

type state struct {
	word    string
	guessed []byte
	missed  int
}

func main() {
	initStates()
	clearScreen()
	word, _ := ask.HiddenAsk("Bitte gib dein Wort ein: ")

	s := state{word, []byte{}, 0}

	for {
		clearScreen()
		w := renderWord(s)
		h, dead := renderHangman(s)
		g := renderGuesses(s)
		fmt.Println(w)
		fmt.Println(h)
		fmt.Println(g)

		if dead {
			clearScreen()
			fmt.Println("Leider verloren! Das Wort war:", s.word)
			fmt.Println("Danke fürs Spielen! Bis zum nächsten Mal.")
			break
		}

		buf := bufio.NewReader(os.Stdin)
		b, _ := buf.ReadByte()
		s.guessed = append(s.guessed, b)
		if !strings.Contains(s.word, string(b)) {
			s.missed++
		}

		if complete(s) {
			clearScreen()
			fmt.Println("Gratulation! Das Wort ist:", s.word)
			fmt.Println("Danke fürs Spielen! Bis zum nächsten Mal.")
			break

		}
	}

}
func renderGuesses(s state) string {
	return string(s.guessed)
}

func complete(s state) bool {
	w := renderWord(s)

	return !strings.Contains(w, "_")
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func renderHangman(i state) (string, bool) {
	if i.missed == len(states)-1 {
		return states[i.missed], true
	}
	return states[i.missed], false
}

func renderWord(s state) string {
	output := make([]byte, len(s.word))
	for i, l := range s.word {
		if contains(s.guessed, byte(l)) {
			output[i] = byte(l)
		} else {
			output[i] = '_'
		}
	}

	return string(output)
}
func contains(bytes []byte, i byte) bool {
	for _, b := range bytes {
		if b == i {
			return true
		}
	}

	return false
}

var states = make([]string, 9)

func initStates() {
	states[0] = `
`
	states[1] = `

    |
    |
    |
    |
    |
    |
`
	states[2] = `
    _________
    |
    |
    |
    |
    |
    |
`
	states[3] = `
    _________
    |         |
    |         0
    |
    |
    |
    |
`
	states[4] = `
    _________
    |         |
    |         0
    |         |
    |
    |
    |
`
	states[5] = `
    _________
    |         |
    |         0
    |        /|
    |
    |
    |
`
	states[6] = `
    _________
    |         |
    |         0
    |        /|\
    |
    |
    |
`
	states[7] = `
    _________
    |         |
    |         0
    |        /|\
    |        /
    |
    |
`
	states[8] = `
    _________
    |         |
    |         0
    |        /|\
    |        / \
    |
    |
`

}

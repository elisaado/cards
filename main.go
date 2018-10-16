package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	. "github.com/logrusorgru/aurora"
)

var s Set

func main() {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}

	parseFile()

	fmt.Println(Cyan(s.Name).Bold())
	fmt.Println(Cyan(s.Description), "\n")
	fmt.Println(Sprintf(Blue("From %s to %s"), Bold(s.TermLanguage), Bold(s.DefinitionLanguage)))

	StartSession()
}

func StartSession() {
	cardsLeft := Shuffle(s.Cards)

	for _, card := range cardsLeft {
		// Show progress
		fmt.Println(Sprintf(Cyan("Card %s out of %s"), Bold(strconv.Itoa(len(s.Cards))), Bold(strconv.Itoa(len(cardsLeft)))))

		fmt.Println(">", Magenta(card.Term).Bold())
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("< ")
		text, _ := reader.ReadString('\n')
		if len(text) == 0 { // ^D makes it have no \n and thus length 0
			fmt.Println()
		}

		text = strings.TrimSpace(text)

		if strings.ToLower(text) == strings.ToLower(card.Definition) {
			s.Cards = remove(s.Cards, card)
			fmt.Println(Green("✓ Correct!").Bold())
		} else {
			fmt.Println(Red("✕ Incorrect").Bold())
			fmt.Println(Red("The correct answer was"), Green(card.Definition).Bold())
		}

		fmt.Println()
	}

	if len(cardsLeft) != 0 {
		StartSession()
	} else {
		fmt.Println("wow good u finiscc")
	}
}

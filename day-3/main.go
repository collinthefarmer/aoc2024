package main

import (
	_ "embed"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed resc/input.txt
var input string

type Mul struct {
	A int
	B int
}

func (m *Mul) Exec() int {
	return m.A * m.B
}

func extractMulCommands(text string) []Mul {
	var commands []Mul

	commandPattern := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)")
	matches := commandPattern.FindAllString(text, -1)

	for _, match := range matches {
		as := match[4:strings.Index(match, ",")]
		a, err := strconv.Atoi(as)
		if err != nil {
			panic(err)
		}

		bs := match[strings.Index(match, ",")+1 : len(match)-1]
		b, err := strconv.Atoi(bs)
		if err != nil {
			panic(err)
		}

		commands = append(commands, Mul{A: a, B: b})
	}
	return commands
}

func extractEnabledMulCommands(text string) []Mul {
	var commands []Mul

	textLen := len(text)
	enabled := true
	var start int
	for i := range text {
		if enabled && i+7 < textLen && text[i:i+7] == "don't()" {
			enabledCommands := extractMulCommands(text[start:i])
			commands = append(commands, enabledCommands...)
			enabled = false
		} else if !enabled && i+4 < textLen && text[i:i+4] == "do()" {
			enabled = true
			start = i + 4
		}

	}

	return commands
}

func main() {
	print("do()"[0:4])
	commands := extractMulCommands(input)

	var sum int
	for _, cmd := range commands {
		sum += cmd.Exec()
	}

	log.Printf("Pt. 1: sum of mul commands: %v", sum)

	var enabledSum int
	enabledCommands := extractEnabledMulCommands(input)
	for _, cmd := range enabledCommands {
		enabledSum += cmd.Exec()
	}

	log.Printf("Pt. 2: sum of enabled mul commands %v", enabledSum)
}

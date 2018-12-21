package main

import (
	"fmt"

	"github.com/chzyer/readline"
	zxcvbn "github.com/nbutton23/zxcvbn-go"
)

const (
	Cyan          = 36
	Green         = 32
	Magenta       = 35
	Red           = 31
	Yellow        = 33
	BackgroundRed = 41
)

// Reset sequence
var ColorResetEscape = "\033[0m"

// ColorResetEscape translates a ANSI color number to a color escape.
func ColorEscape(color int) string {
	return fmt.Sprintf("\033[0;%dm", color)
}

// Colorize the msg using ANSI color escapes
func Colorize(msg string, color int) string {
	return ColorEscape(color) + msg + ColorResetEscape
}

func createStrengthPrompt(password []rune) string {
	symbol, color := "", Red
	strength := zxcvbn.PasswordStrength(string(password), nil)

	switch {
	case strength.Score <= 1:
		symbol = "✗"
		color = Red
	case strength.Score <= 2:
		symbol = "⚡"
		color = Magenta
	case strength.Score <= 3:
		symbol = "⚠"
		color = Yellow
	case strength.Score <= 4:
		symbol = "✔"
		color = Green
	}

	prompt := Colorize(symbol, color)
	if strength.Entropy > 0 {
		entropy := fmt.Sprintf(" %3.0f", strength.Entropy)
		prompt += Colorize(entropy, Cyan)
	} else {
		prompt += Colorize(" ENT", Cyan)
	}

	prompt += Colorize(" New Password: ", color)
	return prompt
}

func main() {
	rl, err := readline.New("")
	if err != nil {
		return
	}
	defer rl.Close()

	setPasswordCfg := rl.GenPasswordConfig()
	setPasswordCfg.SetListener(func(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
		rl.SetPrompt(createStrengthPrompt(line))
		rl.Refresh()
		return nil, 0, false
	})

	pswd, err := rl.ReadPasswordWithConfig(setPasswordCfg)
	if err != nil {
		return
	}

	fmt.Println("Your password was:", string(pswd))
}

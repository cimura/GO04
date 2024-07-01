package piscine

import (
	"os"
	"ft"
)

func PrintProgramname() {
	var argstr string
	arg := os.Args[0]
	argstr = arg[2:]

	for _, r := range argstr {
		ft.PrintRune(r)
	}
}
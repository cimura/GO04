package piscine

import (
	"os"
	"ft"
)

func Len(Args []string) int {
	count := 0
	for range Args {
		count++
	}
	return count
}

func Printparams() {
	var argstr string
	
	for i := 1; i < Len(os.Args); i++ {
		argstr += os.Args[i] + "\n"
	}
	for _, r := range argstr {
		ft.PrintRune(r)
	}
}
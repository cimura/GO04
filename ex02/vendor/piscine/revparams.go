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

func Revparams() {
	var argstr string
	
	for i := Len(os.Args)-1; i > 0 ; i-- {
		argstr += os.Args[i] + "\n"
	}
	for _, r := range argstr {
		ft.PrintRune(r)
	}
}
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

func Swap(a *string, b *string) {
	if *a > *b {
		tmp := *a
		*a = *b
		*b = tmp
	}
}

func Sortparams() {
	argCount := Len(os.Args) - 1
	var argstr [1000]string 
	for i := 0; i < argCount; i++ {
		argstr[i] = os.Args[i + 1]
	}
	for i := 0; i < argCount - 1; i++ {
		for j := 0; j < argCount - 1 - i; j++ {
			if argstr[j] > argstr[j + 1] {
				Swap(&argstr[j], &argstr[j + 1])
			}
		}
	}
	for i := 0; i < argCount; i++ {
		for _, r := range argstr[i] {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}

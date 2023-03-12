package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	if len(a) <= 0 {
		z01.PrintRune(' ')
	} else {
		for i, v := range a {
			tabrune := []rune(v)
			for _, d := range tabrune {
				z01.PrintRune(d)
			}
			if len(a) > i {
				z01.PrintRune(10)
			}
		}
	}
}

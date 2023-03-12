package piscine

func ShoppingSummaryCounter() {
	menu := map[string]int{
		"Burger": 2,
		"Water":  4,
		"Carrot": 4,
		"Coffee": 1,
		"Chips":  1,
	}
	for item, value := range menu {
		Printf("%s => %d$\n", item, value)
	}
}

func Printf(s, item string, value int) {
	panic("unimplemented")
}

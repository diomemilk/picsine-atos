package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(oder string) int {
	burger := food{15}
	chips := food{10}
	nuggets := food{12}
	switch oder {
	case "burger":
		return burger.preptime
	case "chips":
		return chips.preptime
	case "nuggets":
		return nuggets.preptime
	default:
		return 404
	}
}

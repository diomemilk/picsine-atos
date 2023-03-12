package piscine

func ShoppingListSort(slice []string) []string {
	lenSlice := 0
	for i := range slice {
		lenSlice = i + 1
	}
	for i := 0; i < lenSlice-1; i++ {
		for j := i + 1; j < lenSlice; j++ {
			if len(slice[i]) > len(slice[j]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

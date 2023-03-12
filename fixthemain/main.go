package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
func CloseDoor(ptrDoor Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = "CLOSE"
	return true
}
func IsDoorOpen(ptrDoor Door) bool {
	PrintStr("is the Door opened ?")
	return ptrDoor.state == "OPEN"
}
func IsDoorClose(ptrDoor Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == "CLOSE"
}
func OpenDoor(ptrDoor Door) {
	PrintStr("Door Opening...")
}
func main() {
	door := &Door{}
	door.state = "OPEN"
	OpenDoor(*door)
	if IsDoorClose(*door) {
		OpenDoor(*door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(*door)
	}
}

package piscine

func JumpOver(str string) string {
	length := len(str)
	retour := ""
	if length < 3 {
		return "\n"
	} else {
		for i := 2; i < length; i += 3 {
			retour += string(str[i])
		}
	}

	return retour + "\n"
}

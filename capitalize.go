package piscine

func Capitalize(s string) string {
	tabS := []rune(s)

	lettre := true

	for i := 0; i < len(s); i++ {
		if CheckRune(tabS[i]) == true && lettre {
			if tabS[i] >= 'a' && tabS[i] <= 'z' {
				tabS[i] = tabS[i] - 32
			}
			lettre = false
		} else if tabS[i] >= 'A' && tabS[i] <= 'Z' {
			tabS[i] = tabS[i] + 32
		} else if CheckRune(tabS[i]) == false {
			lettre = true
		}
	}
	return string(tabS)
}

func CheckRune(a rune) bool {
	find := false
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		find = true
	}
	return find
}

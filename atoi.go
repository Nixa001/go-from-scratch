package piscine

func Atoi(s string) int {
	Xvalue := 0
	Zvalue := 1
	for i, n := range s {
		Yvalue := 0
		if n < '0' || n > '9' {
			if n == '-' || n == '+' {
				if i != 0 {
					return 0
				}
				if n == '-' {
					Zvalue = -1
				}
			} else {
				return 0
			}
		}
		for i := '1'; i <= n; i++ {
			Yvalue++
		}
		Xvalue = Xvalue*10 + Yvalue*Zvalue
	}
	return Xvalue
}

// func Atoi(s string) int {
// 	result := 0
// 	mult := 1
// 	i := 0

// 	if len(s) == 0 {
// 		return 0
// 	}
// 	if s[0] == '-' {
// 		mult = -1
// 		i++
// 	} else if s[0] == '+' {
// 		mult = 1
// 		i++
// 	}
// 	for i < len(s) {
// 		result = result*10 + int(s[i]) - 48

// 		if int(s[i])-48 > 9 || int(s[i])-48 < 0 {
// 			return 0
// 		}

// 		i++
// 	}
// 	return result * mult
// }

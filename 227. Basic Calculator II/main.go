package main

import "unicode"

func calculate(s string) int {
	stk := make([]int, 0)
	num := 0
	sign := '+'
	for i, r := range s {
		if unicode.IsDigit(r) {
			num = num*10 + int(r-'0')

			if i != len(s)-1 {
				continue
			}
		}

		if r == ' ' && i != len(s)-1 {
			continue
		}

		switch sign {
		case '+':
			stk = append(stk, num)
		case '-':
			stk = append(stk, -num)
		case '*':
			newNum := stk[len(stk)-1] * num
			stk[len(stk)-1] = newNum
		case '/':
			newNum := stk[len(stk)-1] / num
			stk[len(stk)-1] = newNum
		}

		num = 0
		sign = r
	}

	res := 0
	for _, el := range stk {
		res += el
	}
	return res
}
func main() {

}

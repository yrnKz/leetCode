package main

import "strconv"

func main() {
	println(solveEquation("0x+0x=0x"))
}

func solveEquation(equation string) string {
	ex := equation
	var flag = false // 判断是否过=
	numX, number := 0, 0
	for pos := 0; pos < len(ex); {
		//print(pos)
		//print("  " + string(ex[pos]) + "  ")
		if ex[pos] == 'x' {
			numX = f(flag, 1, numX, "+")
			pos++
		} else if ex[pos] == '=' {
			flag = true
			pos++
		} else if ex[pos] == '+' || ex[pos] == '-' {
			temp := ""
			i := pos + 1
			for ; i < len(ex); i++ {
				if ex[i] != '=' && ex[i] != '+' && ex[i] != '-' && ex[i] != 'x' {
					temp = temp + string(ex[i])
				} else {
					break
				}
			}
			n, _ := strconv.Atoi(temp)
			//println(n)
			if i < len(ex) && ex[i] == 'x' {
				if n == 0 && len(temp) == 0 {
					n = 1
				}
				numX = f(flag, n, numX, string(ex[pos]))
				pos = i + 1
			} else {
				number = f(flag, n, number, string(ex[pos]))
				pos = i
			}
		} else {
			temp := ""
			i := pos
			for ; i < len(ex); i++ {
				if ex[i] != '=' && ex[i] != '+' && ex[i] != '-' && ex[i] != 'x' {
					temp = temp + string(ex[i])
				} else {
					break
				}
			}
			n, _ := strconv.Atoi(temp)
			//println(n)
			if i < len(ex) && ex[i] == 'x' {
				numX = f(flag, n, numX, "+")
				pos = i + 1
			} else {
				number = f(flag, n, number, "+")
				pos = i
			}
		}
	}
	//println("----------")
	//println(number)
	//println(numX)

	number = -number
	if numX == 0 && number != 0 {
		return "No solution"
	}
	if numX == 0 && number == 0 {
		return "Infinite solutions"
	}
	r := number / numX
	r1 := strconv.Itoa(r)
	return "x=" + r1
}
func f(flag bool, number, result int, Conform string) int {
	if !flag {
		if Conform == "+" {
			result += number
			return result
		} else {
			result -= number
		}
	} else {
		if Conform == "+" {
			result -= number
			return result
		} else {
			result += number
		}
	}
	return result
}

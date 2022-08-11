package main

func main() {

}

func reformat(s string) string {
	var number, letter, r string
	for i := 0; i < len(s); i++ {
		if s[i] <= '9' && s[i] >= '0' {
			number += string(s[i])
		} else {
			letter += string(s[i])
		}
	}

	if len(number)-1 == len(letter) {
		for i := 0; i < len(letter); i++ {
			r += string(number[i]) + string(letter[i])
		}
		r += string(number[len(number)-1])
		return r
	} else if len(number)+1 == len(letter) {
		for i := 0; i < len(number); i++ {
			r += string(letter[i]) + string(number[i])
		}
		r += string(letter[len(letter)-1])
		return r
	} else if len(number) == len(letter) {
		for i := 0; i < len(number); i++ {
			r += string(letter[i]) + string(number[i])
		}
		return r
	}
	return r
}

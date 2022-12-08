package determine_color_of_a_chessboard_square

func squareIsWhite(coordinates string) bool {
	if len(coordinates) != 2 {
		return false
	}
	r0 := byte(coordinates[0])
	r1 := byte(coordinates[1]) + 1
	return (r1+r0)%2 == 0
}

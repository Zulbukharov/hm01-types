package chess

import (
	"fmt"
	"strconv"
)

func CanKnightAttack(white, black string) (bool, error) {
	var a, b, c, d int
	a, b = conv_pos(white)
	c, d = conv_pos(black)

	if a == c && b == d {
		return false, fmt.Errorf("equal pos on board")
	}

	if a > 8 || b > 8 || c > 8 || d > 8 || a < 1 || b < 1 || c < 1 || d < 1 {
		return false, fmt.Errorf("equal pos on board")
	}

	if a+1 == c && b+2 == d {
		return true, nil
	} else if a+1 == c && b-2 == d {
		return true, nil
	} else if a-1 == c && b+2 == d {
		return true, nil
	} else if a-1 == c && b-2 == d {
		return true, nil
	}

	if a+2 == c && b+1 == d {
		return true, nil
	} else if a+2 == c && b-1 == d {
		return true, nil
	} else if a-2 == c && b+1 == d {
		return true, nil
	} else if a-2 == c && b-1 == d {
		return true, nil
	}

	return false, nil
}

func conv_pos(pos string) (int, int) {
	err1 := "not found on board"

	if len(pos) != 2 {
		fmt.Println(err1)
		return 0, 0
	}
	temp, err := strconv.Atoi(string(pos[1]))
	if err != nil {
		fmt.Println(err1)
	}
	if temp < 1 || temp > 8 {
		return 0, 0
	}

	switch pos[0] {
	case 'a':
		return temp, 1
	case 'b':
		return temp, 2
	case 'c':
		return temp, 3
	case 'd':
		return temp, 4
	case 'e':
		return temp, 5
	case 'f':
		return temp, 6
	case 'g':
		return temp, 7
	case 'h':
		return temp, 8
	default:
		return 0, 0
	}
}

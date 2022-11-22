package chess

import (
	"errors"
)

// CanKnightAttack checks if two knights on specified positions can attack each other
func CanKnightAttack(white, black string) (bool, error) {
	if white == black { // positions are the same
		return false, errors.New("error1")
	}
	if white[0] < 'a' || white[0] > 'h' || white[1] < '1' || white[1] > '8' { // invalid symbol and range of position
		return false, errors.New("error2")
	}
	if black[0] < 'a' || black[0] > 'h' || black[1] < '1' || black[1] > '8' {
		return false, errors.New("error3")
	}

	w1, w2 := int(white[0]-'a'), int(white[1]) // converting a-h positions using ascii logic to int, and range to int

	b1, b2 := int(black[0]-'a'), int(black[1])

	switch { // check for variation of white and black attacking each other by dislocation
	case w1+2 == b1 && w2+1 == b2:
		return true, nil
	case w1+2 == b1 && w2-1 == b2:
		return true, nil
	case w1-2 == b1 && w2+1 == b2:
		return true, nil
	case w1-2 == b1 && w2-1 == b2:
		return true, nil
	case w1+1 == b1 && w2+2 == b2:
		return true, nil
	case w1+1 == b1 && w2-2 == b2:
		return true, nil
	case w1-1 == b1 && w2+2 == b2:
		return true, nil
	case w1-1 == b1 && w2-2 == b2:
		return true, nil
	}

	return false, nil
}

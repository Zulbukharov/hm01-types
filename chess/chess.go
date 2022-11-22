package chess

import (
	"errors"
)

type Attack struct {
	m     int
	n     int
	i     int
	j     int
	a     string
	board [][]string
}

// CanKnightAttack checks if two knights on specified positions can attack each other

func CanKnightAttack(white, black string) (bool, error) {
	check, err := checkingError(white, black)
	if check == false {
		return false, err
	}
	checkAttack, err := checkAttackPosition(white, black)
	return checkAttack, err
}

func checkingError(white, black string) (bool, error) {
	if white == black {
		return false, errors.New("cannot be same")
	}
	if white[0] < 'a' || white[0] > 'h' || white[1] < '1' || white[1] > '8' {
		return false, errors.New("out of range")
	}
	if black[0] < 'a' || black[0] > 'h' || black[1] < '1' || black[1] > '8' {
		return false, errors.New("out of range")
	}

	return true, nil
}

func checkAttackPosition(white, black string) (bool, error) {
	white1, white2 := int(white[0]-'a'), int(white[1]) // 2
	black1, black2 := int(black[0]-'a'), int(black[1])
	if white1+1 == black1 && white2-2 == black2 {
		return true, nil
	} else if white1+2 == black1 && white2-1 == black2 {
		return true, nil
	} else if white1+2 == black1 && white2+1 == black2 {
		return true, nil
	} else if white1+1 == black1 && white2+2 == black2 {
		return true, nil
	} else if white1-1 == black1 && white2+2 == black2 {
		return true, nil
	} else if white1-2 == black1 && white2+1 == black2 {
		return true, nil
	} else if white1-2 == black1 && white2-1 == black2 {
		return true, nil
	} else if white1-1 == black1 && white2-2 == black2 {
		return true, nil
	} else {
		return false, nil
	}
}

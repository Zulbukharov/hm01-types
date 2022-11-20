package chess

import "errors"

// CanKnightAttack checks if two knights on specified positions can attack each other

func CanKnightAttack(white, black string) (bool, error) {
	white1, white2, err := Devideposition(white)
	if err != nil {
		return false, err
	}
	black1, black2, err := Devideposition(black)
	if err != nil {
		return false, err
	}
	if white1 == black1 && white2 == black2 {
		err := errors.New("Invalid position")
		return false, err
	}
	if white1-2 == black1 && white2+1 == black2 ||
		white1-1 == black1 && white2-2 == black2 ||
		white1-1 == black1 && white2+2 == black2 ||
		white1+1 == black1 && white2-2 == black2 ||
		white1+1 == black1 && white2+2 == black2 ||
		white1+2 == black1 && white2-1 == black2 ||
		white1+2 == black1 && white2+1 == black2 {
		return true, nil
	}

	return false, nil
}
func Devideposition(position string) (int, int, error) {
	if len([]rune(position)) != 2 {
		err := errors.New("Invalid position")
		return 0, 0, err
	}
	res := []rune(position)
	firstPos := int(res[0]-'a') + 1
	secondPos := int(res[1] - '0')
	if firstPos < 1 || firstPos > 8 || secondPos < 1 || secondPos > 8 {
		err := errors.New("Invalid Position")
		return 0, 0, err
	}

	return firstPos, secondPos, nil
}

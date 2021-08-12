package hasface

const (
	AllFace = iota
	NonFace
	HasFace
)

func CheckHasFace(hasFaceArr []int) int {
	if len(hasFaceArr) == 0 {
		return AllFace
	}

	tmpNonFace := false
	tmpHasFace := false

	for _, v := range hasFaceArr {
		switch v {
		case 0:
			return AllFace
		case 1:
			if tmpNonFace {
				return AllFace
			}
			tmpHasFace = true
		case 2:
			if tmpHasFace {
				return AllFace
			}
			tmpNonFace = true
		}
	}
	if tmpNonFace {
		return NonFace
	}
	return HasFace
}

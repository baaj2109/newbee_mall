package enum

type GoodsCategoryLevel int

const (
	Default    GoodsCategoryLevel = 0
	LevelOne   GoodsCategoryLevel = 1
	LevelTwo   GoodsCategoryLevel = 2
	LevelThree GoodsCategoryLevel = 3
)

func (g GoodsCategoryLevel) Code() int {
	switch g {
	case LevelOne:
		return 1
	case LevelTwo:
		return 2
	case LevelThree:
		return 3
	default:
		return 0
	}
}

package utils

type direction [2]int

func GetAllCellDirections() []direction{
	return []direction{{-1,-1}, {-1,0}, {-1,1}, {0,1}, {1,1}, {1,0}, {1,-1}, {0,-1}}
}

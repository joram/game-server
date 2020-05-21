package utils

var currID = -1

func NextID() int {
	currID += 1
	return currID
}

package ids

var currIDs = map[string]int{}

func NextID(t string) int {
	if _, ok := currIDs[t]; !ok {
		currIDs[t] = -1
	}
	currIDs[t] += 1
	return currIDs[t]
}

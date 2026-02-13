func finalPositionOfSnake(n int, commands []string) int {
	mymap := map[string]int{
		"RIGHT": 1,
		"LEFT": -1,
		"DOWN": n,
		"UP": -n,
	}

	var out int
	for _, val := range commands {
		out += mymap[val]
	}
	return out
}
package towerhopper

func isHoppable(towers []int) bool {
	if len(towers) == 0 {
		return false
	}
	return helper(0, towers[0], towers)
}

func helper(currentPosition, currentHeight int, towers []int) bool {
	// fmt.Println(currentPosition, currentHeight, towers)
	if (currentPosition+1)+currentHeight > len(towers) {
		return true
	}

	// loop through each possibility
	for i := 1; i <= currentHeight; i++ {
		// fmt.Println("i:", i)
		nextPosition := currentPosition + i
		// fmt.Println("next position:", nextPosition)
		if helper(nextPosition, towers[nextPosition], towers) {
			return true
		}
	}

	return false
}

package fibonacci

func FibbRecursion(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return FibbRecursion(x-1) + FibbRecursion(x-2)
}

func FibbMapCalc(x int) int {
	fibbMap := make(map[int]int, x+1)

	fibbMap[0] = 0
	fibbMap[1] = 1

	for i := 2; i < x+1; i++ {
		fibbMap[i] = fibbMap[i-1] + fibbMap[i-2]
	}

	if x == 0 {
		return fibbMap[0]
	} else if x == 1 {
		return fibbMap[1]
	}

	return fibbMap[x]
}

func FibbMapRecursion(numMap map[int]int, x int) (map[int]int, int) {

	numMap[0] = 0
	numMap[1] = 1

	length := len(numMap)

	numMap[length] = numMap[length-1] + numMap[length-2]

	if len(numMap) == x+1 {
		return numMap, x
	}
	if x == 1 {
		return numMap, x
	}

	return FibbMapRecursion(numMap, x)
}

func FibbMapRecursion2(numMap map[int]int) map[int]int {
	length := len(numMap)
	numMap[length] = numMap[length-1] + numMap[length-2]

	if length == numMap[0] {
		return numMap
	}
	if numMap[0] == 1 || numMap[0] == 2 {
		return numMap
	}

	return FibbMapRecursion2(numMap)
}

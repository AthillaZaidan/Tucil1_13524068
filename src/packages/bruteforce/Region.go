package bruteforce

func countRegion(arr [][]byte, row, col int) int {
	var color [26]byte
	for i := 0; i < 26; i++ {
		color[i] = 0
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			color[arr[i][j]-'A']++
		}
	}

	var regionNum int = 0
	for i := 0; i < 26; i++ {
		if color[i] != 0 {
			regionNum++
		}
	}
	return regionNum
}

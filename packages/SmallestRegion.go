package utils

func SmallestRegion(arr [][]byte, row, col int) int {
	var color [26]byte
	for i := 0; i < 26; i++ {
		color[i] = 0
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++{
			if (arr[i][j] == '.') {
				continue
			} else {
				color[arr[i][j] - 'A']++
			}
		}
	}
	min := color[0]
	minPos := 0
	for i := 1; i < 26; i++{
		if color[i] < min {
			minPos = i
			min = color[i]
		} 
	}
	return minPos
}


func DeleteRegion(arr [][]byte, row, col int, colorDel byte){
	for i := 0; i < row; i++{
		for j := 0; j < col; j++{
			if arr[i][j] == colorDel {
				arr[i][j] = '.'
			}
		}
	}
}

func countRegion(arr [][]byte, row, col int) int{
	var color [26]byte
	for i := 0; i < 26; i++ {
		color[i] = 0
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++{
			color[arr[i][j] - 'A']++
		}
	}

	var regionNum int = 0;
	for i := 0; i < 26; i++{
		if color[i] != 0 {
			regionNum ++;
		}
	}
	return regionNum
}


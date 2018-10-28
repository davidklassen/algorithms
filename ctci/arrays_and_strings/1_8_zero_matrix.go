package arrays_and_strings

func ZeroMatrix(matrix [][]int, w, h int) [][]int {
	xs := make(map[int]struct{})
	ys := make(map[int]struct{})
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if matrix[y][x] == 0 {
				xs[x] = struct{}{}
				ys[y] = struct{}{}
			}
		}
	}
	for x := range xs {
		for y := 0; y < h; y++ {
			matrix[y][x] = 0
		}
	}
	for y := range ys {
		for x := 0; x < w; x++ {
			matrix[y][x] = 0
		}
	}
	return matrix
}

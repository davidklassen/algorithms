package arrays_and_strings

func getOffset(x, y, size int) int {
	return 4 * (x + y*size)
}

func copyPixel(src, dst []byte, from, to int) {
	dst[to] = src[from]
	dst[to+1] = src[from+1]
	dst[to+2] = src[from+2]
	dst[to+3] = src[from+3]
}

func RotateMatrix(img []byte, size int) []byte {
	for k := 0; k < size/2; k++ {
		for i := 0; i < size-2*k-1; i++ {
			tmp := make([]byte, 4)
			max := size - 1
			copyPixel(img, tmp, getOffset(i+k, k, size), 0)
			copyPixel(img, img, getOffset(k, max-i-k, size), getOffset(i+k, k, size))
			copyPixel(img, img, getOffset(max-i-k, max-k, size), getOffset(k, max-i-k, size))
			copyPixel(img, img, getOffset(max-k, i+k, size), getOffset(max-i-k, max-k, size))
			copyPixel(tmp, img, 0, getOffset(max-k, i+k, size))
		}
	}
	return img
}

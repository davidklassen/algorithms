package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func filter(s []int, fn func(int) bool) []int {
	p := []int{}
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func qsort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	pix := rand.Intn(len(arr))
	p := arr[pix]
	arr = append(arr[:pix], arr[pix+1:]...)
	left := filter(arr, func(i int) bool {
		return i < p
	})
	right := filter(arr, func(i int) bool {
		return i >= p
	})

	return append(qsort(left), append([]int{p}, qsort(right)...)...)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	arr := []int{}
	for _, v := range strings.Split(string(bytes), "\n") {
		n, err := strconv.Atoi(v)

		if err == nil {
			arr = append(arr, n)
		}
	}

	for _, n := range qsort(arr) {
		fmt.Println(n)
	}
}

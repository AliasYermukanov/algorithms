package main

import (
	"fmt"
	"math"
)

func show(x,y int)  {
	fmt.Println(x, "in", y, "index")
}

func Min(a, b int) int {
	if a > b {
		return b
	} else if b > a {
		return a
	} else {
		return a
	}
}

func main() {
	array := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	needed := 55

	show(LinearSearch(array, needed))
	show(BinarySearch(array, 0, len(array)-1, needed))
	show(JumpSearch(array, needed))
	show(InterpolationSearch(array,needed,0,len(array)-1))
}

func InterpolationSearch(ints []int, x int, start int, end int) (int, int) {
	pos:= start + ((end-start)/(ints[end]-int(start))*(x-ints[start]))
	fmt.Println(pos)

	if ints[pos] == x {
		return ints[pos], pos
	}else if ints[pos] > x {
		InterpolationSearch(ints,x,start,pos-1)
	}else if ints[pos] < x {
		InterpolationSearch(ints,x,pos+1,end)
	}

	return -1, -1
}



func JumpSearch(ints []int, x int) (int, int) {
	step := int(math.Floor(math.Sqrt(float64(len(ints)))))
	prev := 0

	for ints[Min(step, len(ints)-1)] < x {
		prev = step
		step += step
		if prev >= len(ints) {
			return -1, -1
		}
	}
	for ints[prev] < x {
		prev++

		if ints[prev] == x {
			return ints[prev], prev
		}
	}

	return -1,-1
}

func BinarySearch(ints []int, l int, r int, x int) (int, int) {
	if r >= 1 {
		mid := l + (r-l)/2

		if ints[mid] == x {
			return x, mid
		}

		if mid > x {
			return BinarySearch(ints, l, mid-1, x)
		}

		if mid < x {
			return BinarySearch(ints, mid, r-1, x)
		}
	}
	return -1, -1
}

func LinearSearch(ints []int, x int) (int, int) {
	for key, value := range ints {
		if value == x {
			return x, key
		}
	}
	return -1, -1
}

package main

import "fmt"

func main() {
	var slice []int = []int{3, 6, 7, 2, 9, 4}
	fmt.Println(slice)
	fmt.Println("===========================")
	qSort(&slice)
	fmt.Println(slice)
}

func qSort(slice *[]int) {
	initiateQuickSort(slice, 0, len(*slice)-1)
}

func initiateQuickSort(slicep *[]int, first int, last int) {
	if first >= last {
		return
	}

	var slice = *slicep
	var wall = first

	for i := first; i < last; i++ {
		if slice[i] < slice[last] {
			slice[wall], slice[i] = slice[i], slice[wall]
			wall++
		}
	}
	slice[wall], slice[last] = slice[last], slice[wall]
	if wall > 0 {
		initiateQuickSort(&slice, first, wall-1)
	}
	initiateQuickSort(&slice, wall+1, last)

}

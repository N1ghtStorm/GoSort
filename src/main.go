package main

func main() {

}

func qSort(slice *[]int) {
	initiateQuickSort(slice, 0, len(*slice))
}

func initiateQuickSort(slice *[]int, first int, last int) {
	if last > 2 {
		return
	}

	//part := partiton()
}

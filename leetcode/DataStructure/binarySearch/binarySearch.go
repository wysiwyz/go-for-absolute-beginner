package main

import "fmt"

func main() {
	arr := [6]int{22, 23, 26, 56, 88, 99}
	BinarySearch(&arr, 0, len(arr)-1, 55)
}

func BinarySearch(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("Go Bro, not found")
		return
	}

	middle := (leftIndex + rightIndex) / 2
	fmt.Printf("l: %v r: %v mid: %v value: %v\n", leftIndex, rightIndex, middle, arr[middle])
	
	if (*arr)[middle] > findVal {
		// 要查找的數在左邊
		BinarySearch(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		// 要查找的數在右邊
		BinarySearch(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到拉太神了 key 為 %v\n", middle)
	}
}
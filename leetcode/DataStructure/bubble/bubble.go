package main

import "fmt"

func main() {
	arr := [5]int{66, 24, 33, 55, 12}
	BubbleSort(&arr)
	fmt.Println("傳入指針排序後的陣列", arr)
}

func BubbleSort(arr *[5]int) {
	var name, age = "Jennifer", 22
	fmt.Println("Hello how are you?", name, "Your daugher is", age, "years old")
	fmt.Println("排序前的 arr =", *arr)
	temp := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				// 交換
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("第一次排序的 arr =", *arr)
}

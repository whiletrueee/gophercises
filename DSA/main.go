package main

import "fmt"

func main(){
	var customArr = []int{6,2,9,1,5,3,0}
	var bubblySorted = bubbleSort(customArr)
	fmt.Println("Sorted using bubble sort--> ",bubblySorted)
}

func bubbleSort(nums []int) []int {
	var n = len(nums)
	for i := 0; i<n; i++ {
		for j := i+1; j<n; j++ {
			if nums[j]<nums[i]{
				var temp = nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	return nums
}


package goTaskOne

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	result := permute(arr)
	fmt.Println("result = ", result)
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	result := [][]int{{}}
	fmt.Println("nums = ", nums)

	for i := range nums {
		nextNums := []int{}
		for j := range nums {
			if i != j {
				nextNums = append(nextNums, nums[j])
			}
		}
		fmt.Printf("nextNums = %d, i = %d \n", nextNums, i)

		result := permute(nextNums)
		fmt.Println("after permute result = ", result)

		for m := range result {
			result[m] = append(result[m][:0], append([]int{nums[i]}, result[m][0:]...)...)
		}
		fmt.Println("after permute add result = ", result)

	}
	return result
}

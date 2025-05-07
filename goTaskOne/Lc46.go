package goTaskOne

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	result := make([][]int, 0)
	for i := range nums {
		var nextNums []int
		for j := range nums {
			if i != j {
				nextNums = append(nextNums, nums[j])
			}
		}
		currResult := permute(nextNums)
		for m := range currResult {
			currResult[m] = append(currResult[m][:0], append([]int{nums[i]}, currResult[m][0:]...)...)
		}
		result = append(result, currResult...)
	}
	return result
}

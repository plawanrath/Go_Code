package main

// func main() {
// 	nums := []int{2, 7, 11, 15}
// 	target := 9
// 	result := twoSum(nums, target)
// 	fmt.Println(result)
// }

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if val, ok := m[complement]; ok {
			res = append(res, val)
			res = append(res, i)
		} else {
			m[nums[i]] = i
		}
	}
	return res
}

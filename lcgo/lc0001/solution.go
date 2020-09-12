package lc0001

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		diff := target - val
		if j, ok := m[diff]; ok {
			return []int{j, i}
		}
		m[val] = i
	}
	return nil
}
// TwoSum solution fo lc0001
func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}
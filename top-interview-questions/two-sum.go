// https://leetcode.com/problems/two-sum/
package topinterviewquestions

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) //key: element and value: index
	for i := 0; i < len(nums); i++ {
		if idx, ok := m[target-nums[i]]; ok {
			return []int{idx, i}
		}
		m[nums[i]] = i
	}
	return []int{0, 0}
}

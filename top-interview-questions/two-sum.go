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

// This is a Go programming language implementation of a function called twoSum. This function takes an array of integers nums and an integer target as input and returns an array of two integers that represent the indices of the two numbers in the input array that add up to the target.

// Here's how the function works:

// The function initializes a map called m to keep track of the elements in the input array and their indices.
// The function enters a loop that iterates over each element in the input array.
// For each element nums[i] in the input array, the function checks if the complement target-nums[i] is already in the m map. The complement is the number that, when added to nums[i], results in the target. The map is used to store the complement number as the key and the index of the corresponding element in the array as the value.
// If the complement is found in the map, the function returns an array of the indices of the two numbers that add up to the target, which are the value of the complement key in the map and the current index i.
// If the complement is not found in the map, the current element nums[i] is added to the map with its index i as the value.
// If the loop completes without finding a complement that adds up to the target, the function returns an array of two zeros.
// The use of the map m helps to improve the time complexity of the function from O(n^2) to O(n), since checking if an element is in a map takes constant time on average.

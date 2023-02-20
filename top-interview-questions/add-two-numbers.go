// https://leetcode.com/problems/add-two-numbers/
package topinterviewquestions

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	curr := dummyHead
	carry := 0

	for l1 != nil || l2 != nil {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := val1 + val2 + carry
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next
}

// This is a Go programming language implementation of a function called addTwoNumbers. This function takes two singly linked lists l1 and l2 as input and returns a singly linked list that represents the sum of the two input lists.

// The function works as follows:

// The function initializes a dummy head node dummyHead and a current node pointer curr to the dummy head node.
// It also initializes a variable carry to zero, which represents the carryover from the previous digit addition.
// The function enters a loop that continues until both input lists l1 and l2 are empty.
// For each iteration of the loop, the function extracts the current digit values
// from each input list by checking the Val field of the current nodes l1 and l2.
// If either input list has reached the end, the corresponding value is assumed to be zero.
// The function calculates the sum of the current digit values and the carryover carry. If the sum is greater than or equal to 10,
// the carryover is set to 1, and the sum modulo 10 is added to the new node's value.
// Otherwise, the carryover is set to 0, and the sum is added directly to the new node's value.
// The function creates a new node with the computed value and sets it as the next node in the output list, which is pointed to by curr. The curr pointer is then advanced to the next node.
//
// Once both input lists have been fully traversed,
// the function checks if there is a final carryover that needs to be added to the output list.
// If there is, it creates a new node with the carryover value and appends it to the output list.
// The function returns the output list, which is the list of nodes pointed to by the Next field of the dummy head node.
// This approach effectively performs digit-by-digit addition with carryover,
// which is a common approach to adding two numbers represented as linked lists.
// The use of a dummy head node allows for easier management of the output list without having
// to handle the special case of the head node separately.

package main

import (
	"fmt"

	"eyr.question.solving/pkg/ds"
	"eyr.question.solving/pkg/dt"
)

func main() {
	// https://leetcode.com/problems/merge-two-sorted-lists/

	listOne := ds.LinkedList[int]{}
	listOne.Insert(1)
	listOne.Insert(3)
	listOne.Insert(5)

	listTwo := ds.LinkedList[int]{}
	listTwo.Insert(2)
	listTwo.Insert(4)
	listTwo.Insert(6)

	iter := mergeTwoLists(listOne.Head, listTwo.Head)

	for iter != nil {
		fmt.Println(*iter.Val)

		iter = iter.Next
	}
}

func mergeTwoLists[T dt.Comparable](list1 *ds.LinkNode[T], list2 *ds.LinkNode[T]) *ds.LinkNode[T] {

	result := &ds.LinkNode[T]{}
	pointer := result

	for list1 != nil && list2 != nil {

		figure1 := *list1.Val
		figure2 := *list2.Val

		if figure1 < figure2 {
			pointer.Next = list1
			list1 = list1.Next
		} else {
			pointer.Next = list2
			list2 = list2.Next
		}

		pointer = pointer.Next
	}

	if list1 != nil {
		pointer.Next = list1
	} else if list2 != nil {
		pointer.Next = list2
	}

	return result.Next
}

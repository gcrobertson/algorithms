package main

import "fmt"

/*	given:
 *		2 sorted arrays, nums1 & nums2, m # elements in nums1, n # elements in nums2
 *
 *	func:
 *		merge nums2 into nums1
 *
 *	return:
 *		merged array
 *
 *	example:
 *		input:	nums1 := [1,2,3,nil,nil,nil]	// m := 3 number of elements, size / length is 6
 *				nums2 := [2,5,6,]
 *				m := 	 3
 *				n := 	 3
 *		return:			 [1,2,2,3,5,6]
 *
 */

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	fmt.Printf("before merge: %+v\n", nums1)
	merge(nums1, 3, nums2, len(nums2))
	fmt.Printf("after merge: %+v\n", nums1)
}

/*
 * O((n+m)n)
 */
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	var p1, p2 int                  // O(1)
// 	for p1 < len(nums1) && p2 < n { // O(1)
// 		if nums1[p1] > nums2[p2] { // O(1)
// 			copy(nums1[p1+1:], nums1[p1:]) // O(n)
// 			nums1[p1] = nums2[p2]          // O(1)
// 			p2++                           // O(1)
// 		} else { // O(1)
// 			p1++ // O(1)
// 		}
// 	}
// 	if p2 < n { // O(1)
// 		copy(nums1[m+p2:], nums2[p2:]) // O(n)
// 	}
// }

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	for i < m+n && j < n {
		if nums1[i] > nums2[j] {
			copy(nums1[i+1:], nums1[i:])
			nums1[i] = nums2[j]
			j++
			continue
		}
		i++
	}
	if j < n {
		copy(nums1[m+j:], nums2[j:])
	}
}

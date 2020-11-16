package main

import "fmt"

//Node ... recursive data structure
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func main() {

	tree := &Node{
		Key:   57,
		Left:  nil,
		Right: nil,
	}

	tree.Insert(32)
	tree.Insert(24)
	tree.Insert(48)
	tree.Insert(39)
	tree.Insert(35)
	tree.Insert(43)
	tree.Insert(85)
	tree.Insert(81)
	tree.Insert(96)
	tree.Insert(90)
	tree.Insert(78)

	fmt.Printf("the max value in tree is %v and the min value is %v", tree.Max(), tree.Min())
}

// Search ...
func (n *Node) Search(key int) bool {
	if n == nil {
		// if we find a nil node, the key does not exist
		return false
	}
	if n.Key > key {
		return n.Left.Search(key)
	} else if key > n.Key {
		return n.Right.Search(key)
	}
	return true
}

// Insert ... pointer receiver
func (n *Node) Insert(key int) {
	if n.Key > key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	} else if key > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	}
}

// Delete ...
func (n *Node) Delete(key int) *Node {

	if key < n.Key {
		// go to left
		n.Left = n.Left.Delete(key)
	} else if key > n.Key {
		// go to right
		n.Right = n.Right.Delete(key)
	} else {
		// found value to delete. this will point through...
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		// 2 children underneath. we need to find the next highest number to go to our current `delete` position.
		// to do this: find the minimum value of the right subtree.
		// explanation around 18:00 https://www.youtube.com/watch?v=4iYtR5pIMwA

		// once found ... if left or right node is empty, you return the opposite node. why?
		// point to the opposite node. if both nodes are empty we point to a nil deleting the entire subtree. what?
		min := n.Right.Min()
		n.Key = min
		n.Right = n.Right.Delete(min)

		// if there are 2 children.. we need to find the next highest number to go into the position.
		// find minimum value from right subtree.
		// to find minimum node, move to the left forever.
		// to find maximum node, move to the right forever.
	}
	return n
}

// Min ...
// min if you move to the left until you reach nil, you get the min
func (n *Node) Min() int {
	if n.Left == nil {
		return n.Key
	}
	return n.Left.Min()
}

// Max ...
// max if you move to the right until you reach nil, you get the max
func (n *Node) Max() int {
	if n.Right == nil {
		return n.Key
	}
	return n.Right.Max()
}

// https://www.youtube.com/watch?v=4iYtR5pIMwA

/*
 *
 *	"Guess the number between 1 and 100".
 *
 *	"No, its lower than 57..."
 *
 *	"No, its higher than 32..."
 *
 *	This is a binary search algorithm.
 *
 *
 *	3 main aspects to a binary tree:
 *	- Search
 * 	- Insert
 *  - Delete
 *
 */

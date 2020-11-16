package main

func main() {

}

/*	https://medium.com/backendarmy/linked-lists-in-go-f7a7b27a03b9
 *
 *	Circular LinkedList
 *
 *	A `circular linked list` is one which can loop back from teh last node to the first node and vice versa.
 *
 *	You can make the singly or doubly linked lists circular by linking the last node back to the first node.
 *
 *	This can make our music playlist loop back to the first song after pressing next track while playing the last song on the playlist.
 *
 *	Implementation:
 *
 *	root node:
 *	-	Previous would point to the tail node
 *
 *	tail node:
 *	-	Next would point to the root node
 *
 *
 *	Operational Cost:
 *
 *	When we implement any data structure in our code, it is important to know the operational cost.
 *
 *	Linked List Time Complexities:
 *
 *	-	Access: O(n)
 *	-	Search: O(n)
 *	-	Insert: O(1)
 *	-	Delete: O(1)
 *
 *
 *
 *	Linked List v Arrays
 *
 *	Array
 *	-	Access: O(1) if we know index
 *	-	Insert: O(n)	[??]
 *	-	Append to end is easy, Append to middle means moving all elements after by 1 step.
 *
 *	If you need faster insertions int he middle [priority queue] then a linked list may be the better choice.
 *
 *	Linked Lists are heavier in space complexity:
 *	-	Linked List holds data and pointer to next / previous node
 *	-	Array holds only data
 */

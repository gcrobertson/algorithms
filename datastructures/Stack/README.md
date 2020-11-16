# Stacks

An array is a random access data structure.

A linked list is a sequential access data structure.

A subcase of sequential data structures, so-called limited access data structures.

Stacks

A stack is a container of objects that are inserted and removed according to the last-in first-out principle.

Pushdown Stacks have only two operations:
-	Push the item into the stack
-	Pop the item out of the stack

A stack is a limited access data structure - elements can be added and removed from the stack only at the top.
-	Push adds an item to the top of the stack
-	Pop removes the item from the top of the stack

A stack is a recursive data structure. 
-	A stack is either empty or
-	A stack consists of a top and the rest which is a stack

Applications for Stacks
-	Reversing a word
-	"Undo" mechanism in text editors

Implementation:

In the standard library of classes, the data type stack is an `adapter` class.
-	A stack is built on top of other data structures
-	The underlying structure for a stack could be an array, linked list, or any other collection

The unique interface of a stack:
public interface StackInterface<AnyType>
{
	public void push(AnyType e);
	public AnyType pop();
	public AnyType peek();
	public boolean isEmpty();
}

-	Each of these operations must run in constant time O(1).

Array based Implementation:
-	[0,1,2,3,_,_]
-	capacity of 6
-	top is 3
-	when top is -1 stack is empty
-	when top is capacity -1 stack is full

Linked List-based Implementation:

-	Provides the best efficiency of dynamic stack implementation
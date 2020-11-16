// https://github.com/ljun20160606/leetcode


Queues

	A queue is a linear collection [container] of objects that are inserted and removed according to first-in first-out principle.

	NOTE: Stacks are LIFO, Queues are FIFO.

	In the queue only 2 operations are allowed:
	-	Enqueue
		-	insert an item into the back of the queue
	-	Dequeue
		-	remove the front item

	NOTE: The difference between `stack` and `queue` is removing.
	-	In a stack we remove the most recently added item
	-	In a queue we remove the least recently added item

	Queue Implementation

	Queue is also an `adapter` class, meaning that a queue is built on top of other data structures.

	The data structure could be anything like an Array or LinkedList.

	Unique interface of a Queue

	interface QueueInterface<AnyType>
	{
		public boolean isEmpty();
		public AnyType getFront();
		public AnyType dequeue();
		public void enqueue(AnyType e);
		public void clear();
	}

	Each of the above basic operations must run at constant time O(1).

	***

	Circular Queue

	Given:
	-	Array A of size >= 1
	-	Reference back -1
	-	Reference front 0

	etc. Not going to devote time to this now.  It would be great to look up after I know the more basic stuff better.


	***
	***
	***

	Applications

	The simplest two search techniques
	-	Depth-First Search		[DFS]
	-	Breadth-First Search	[BSF]

	These searches describe how looking at the search tree [all possible paths from start] will be traversed.


	Depth-First Search [DFS] with a Stack

	In depth-first search we go down a path until we get to a dead end;
		-	then we `backtrack` or back up [by popping a stack] to get an alternative path.

	-	Create a stack
	-	Create a new choice point
	-	Push the choice point ontot he stack
	-	While (not found and stack is not empty)
		-	Pop the stack
		-	Find all possible choices after the last one tried
		-	Push these choices onto the stack
	-	Return


	Breadth-First Search [BFS] with a Queue

	In breadth-first search we explore all the nearest possibilities by finding all possible successors and enqueue them to a queue.

	-	Create a queue
	-	Create a new choice point
	-	Enqueue the choice point onto the queue
	-	While (not found and queue is not empty)
		-	Dequeue the queue
		-	Find all possible choices after the last one tried
		-	Enqueue these choices ontot he queue
	-	Return






// https://yourbasic.org/golang/implement-fifo-queue/

A simple way to introduce a queue data structure in Go is to use a slice:
- to enqueue use the built-in append
- to dequeue slice off the first element


// https://github.com/ljun20160606/leetcode
package notes 

type  queueInterface interface {
   isEmpty() bool              // O(1)
   peek() interface{}          // O(1)
   dequeue() interface{}       // O(1)
   enqueue(e interface{})      // O(1)
   clear()                     // O(1)
   print()                     // O(n)
}

type  stackInterface interface {
   isEmpty() bool              // O(1)
   peek() interface{}          // O(1)
   pop() interface{}           // O(1)
   push(e interface{})         // O(1)
   clear()                     // O(1)
   print()                     // O(n)
}

type linkedList interface {
   access()    // O(n) 
   search()    // O(n)
   insert()    // O(1)
   delete()    // O(1)
}

type array interface {
   access()    // O(1) 
   search()    // O(n)
   insert()    // O(n)
   delete()    // O(n)
}

# `Stack` [LIFO]
├──   `Array`
└──   `Linked List`

# `Queue` [FIFO]
├──   `Array`
└──   `Linked List`

# Linear data structures
├──   `Array`
├──   `Linked List`
├──   `Stack`
└──   `Queue`

# Non-linear data structures
├──   `Tree`
└──   `Graph`

# Tree Traversal
├── `BFS` 
│   └── `Queue`
└── `DFS` 
    └── `Stack`
        ├── `Preorder`  [root, left, right]
        ├── `Inorder`   [left, root, right]
        └── `Postorder` [left, right, root]
        
# Depth-First Tree Traversal [Stack]

# Breadth-First Tree Traversal []





# Terminology
`Queue` 	a linear collection of objects inserted and removed according to FIFO.
`Stack` 	a linear collection of objects inserted and removed according to LIFO.
`Enqueue`	adds an item to the back of `queue`
`Dequeue`	removes the front item of `queue`
`Push` 		adds an item to the top of `stack`
`Pop` 		removes the item from the top of `stack`

# Queue 
## Queue with Array Data Type
`Enqueue`	append to slice
`Dequeue`	queue = queue[1:]

# Stack
A `stack` is a recursive data structure. 
-	A `stack` is either empty or
-	A `stack` consists of a `top` and the rest which is a `stack`
## Stack Applications 
-	Reversing a word
-	"Undo" mechanism in text editors

### Queue Applications: Bread-First Search [BFS]
In breadth-first search we explore all the nearest possibilities by 
finding all possible successors and enqueuing them.
-	Create a queue
-	Create a new choice point
-	Enqueue the choice point onto the queue
-	While (not found and queue is not empty)
	-	Dequeue the queue
	-	Find all possible choices after the last one tried
	-	Enqueue these choices onto the queue
-	Return

### Stack Applications: Depth-First Search [DFS]
In depth-first search we go down a path until we get to a dead end;
then we `backtrack` or back up [by popping a stack] to get an alternative path.
-	Create a stack
-	Create a new choice point
-	Push the choice point onto the stack
-	While (not found and stack is not empty)
	-	Pop the stack
	-	Find all possible choices after the last one tried
	-	Push these choices onto the stack
-	Return
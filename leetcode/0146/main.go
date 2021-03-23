package main

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {

	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // cache is {1=1}
	lRUCache.Put(2, 2) // cache is {1=1, 2=2}
	lRUCache.Get(1)    // return 1
	lRUCache.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	lRUCache.Get(2)    // returns -1 (not found)
	lRUCache.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	lRUCache.Get(1)    // return -1 (not found)
	lRUCache.Get(3)    // return 3
	lRUCache.Get(4)    // return 4
}

type doubleLinkList struct {
	head *linkedList
	tail *linkedList
}

type linkedList struct {
	val  int
	next *linkedList
	prev *linkedList
}

// LRUCache ...
type LRUCache struct {
	capacity int
	hash     map[int]int
	list     *doubleLinkList
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		hash:     make(map[int]int, capacity),
		list:     &doubleLinkList{},
	}
}

// Get ...
func (t *LRUCache) Get(key int) int {
	if val, ok := t.hash[key]; ok {

		// d

		return val
	}
	return -1
}

// Put ...
func (t *LRUCache) Put(key int, value int) {

	// O(1) map lookup
	if _, ok := t.hash[key]; ok {
		t.hash[key] = value
		return
	}

	new := &linkedList{
		val: value,
	}

	if len(t.hash) < t.capacity {

		// add to map
		t.hash[key] = value

		// add to double linked list
		if t.list.head == nil {
			t.list.head = new
			t.list.tail = new
		} else {
			o := t.list.head
			t.list.head = new
			new.next = o
		}
		return
	}

	// add to map
	t.hash[key] = value

	// remove last element form double linked list
	old := t.list.tail
	t.list.tail = old.prev
	old = nil

	// add to front of linked list

}

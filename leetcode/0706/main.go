package main

func main() {

}

// MyHashMap ...
type MyHashMap struct {
	indexes []int // [1, 1000000]
}

// Constructor ...
/** Initialize your data structure here. */
func Constructor() MyHashMap {

	indexes := make([]int, 10000001)

	// O(n)
	for i := 0; i < 10000001; i++ {
		indexes[i] = -1
	}

	mhm := MyHashMap{
		indexes: indexes,
	}

	return mhm
}

// Put ...
/** value will always be non-negative. */
func (t *MyHashMap) Put(key int, value int) {
	t.indexes[key] = value
}

// Get ...
/** Returns the value to which the specified key is mapped, or -1 if t map contains no mapping for the key */
func (t *MyHashMap) Get(key int) int {
	if t.indexes[key] == -1 {
		return -1
	}
	return t.indexes[key]
}

// Remove ...
/** Removes the mapping of the specified value key if t map contains a mapping for the key */
func (t *MyHashMap) Remove(key int) {
	t.indexes[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

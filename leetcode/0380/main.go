package main

import "math/rand"

// RandomizedSet ...
type RandomizedSet struct {
	hash map[int]int
	x    []int
}

// Constructor Initialize your data structure here.
func Constructor() RandomizedSet {
	return RandomizedSet{
		x:    []int{},
		hash: make(map[int]int),
	}
}

// Insert ... Inserts a value to the set. Returns true if the set did not already contain the specified element.
func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.hash[val]; ok {
		return false
	}

	rs.hash[val] = len(rs.x)
	rs.x = append(rs.x, val)
	return true
}

// Remove Removes a value from the set. Returns true if the set contained the specified element.
func (rs *RandomizedSet) Remove(val int) bool {
	arrKey, ok := rs.hash[val]
	if !ok {
		return false
	}

	last := rs.x[len(rs.x)-1]
	if val == last {
		// if we are removing the last element in the array,
		// we don't need to swap and update index in map
		rs.x = rs.x[:len(rs.x)-1]
	} else {
		// swap and remove the last element makes the delete of the element O(1)
		rs.x[arrKey] = last
		rs.x = rs.x[:len(rs.x)-1]
		// note: we also need to update the index of the last element in the map
		rs.hash[rs.x[arrKey]] = arrKey
	}
	delete(rs.hash, val)
	return true
}

// GetRandom Get a random element from the set.
func (rs *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(rs.x))
	return rs.x[index]
}

func main() {

}

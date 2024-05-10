package main

import "slices"

// https://leetcode.com/problems/design-hashset/
type MyHashSetV1 struct {
	items []int
}

func ConstructorV1_705() MyHashSetV1 {
	return MyHashSetV1{}
}

func (this *MyHashSetV1) AddV1(key int) {
	if !this.ContainsV1(key) {
		this.items = append(this.items, key)
	}
}

func (this *MyHashSetV1) RemoveV1(key int) {
	slices.DeleteFunc(this.items, func(item int) bool {
		return item == key
	})
}

func (this *MyHashSetV1) ContainsV1(key int) bool {
	return slices.Contains(this.items, key)
}

type MyHashSetV2 struct {
	items []int
}

func ConstructorV2_705() MyHashSetV2 {
	return MyHashSetV2{make([]int, 1000001)}
}

func (this *MyHashSetV2) Add(key int) {
	this.items[key] = 1
}

func (this *MyHashSetV2) Remove(key int) {
	this.items[key] = 0
}

func (this *MyHashSetV2) Contains(key int) bool {
	return this.items[key] == 1
}

type MyHashSetV3 struct {
	items [1_000_001]bool
}

func ConstructorV3_705() MyHashSetV3 {
	return MyHashSetV3{}
}

func (this *MyHashSetV3) Add(key int) {
	this.items[key] = true
}

func (this *MyHashSetV3) Remove(key int) {
	this.items[key] = false
}

func (this *MyHashSetV3) Contains(key int) bool {
	return this.items[key]
}

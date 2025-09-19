package main

func removeElFromSlice(sl []int, idx int) []int {
	copy(sl[idx:], sl[idx+1:])
	return sl[:len(sl)-1]
}
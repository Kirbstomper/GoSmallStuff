package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {
	println(isIsomorphic("badc",
		"baba"))

	
	a := inc()

	a()
	a()
	println(a())
}

func inc() func() int{
	i := 0
	return func() int{
		i++
		return i
	}
}
func isIsomorphic(s string, t string) bool {
	rt := []rune(t)
	rs := []rune(s)
	m := make(map[rune]rune)
	vals := make(map[rune]int)

	for i := range rs {
		_, ok := m[rt[i]]
		if !ok {
			m[rt[i]] = rs[i]
		}
		if m[rt[i]] != rs[i] {
			return false
		}
	}
	for _,v := range m{
		vals[v] ++
		if vals[v] == 2{
			return false
		}
	}
	return true

}

	

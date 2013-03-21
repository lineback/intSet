package intSet

import (
	"bytes"
	"strconv"
	"strings"
)

/*
 IntSet - an integer set
*/
type IntSet map[int]struct{}

/*
 returns a new IntSet
*/
func NewIntSet() IntSet {
	return make(IntSet)
}
/*
 adds an int to IntSet
*/	
func (s IntSet) Add(a int) {
	s[a] = struct{}{}
}
/*
 removes an int from IntSet
*/
func (s IntSet) Remove(a int){
	delete(s, a)
}

/*
 returns true if s contains a
*/
func (s IntSet) Contains(a int) bool{
	_, ok := s[a]
	return ok
}

/*
 returns the union of s and r
*/
func (s IntSet) Union(r IntSet) IntSet {
	q := NewIntSet()
	for k,_ := range s {
		q[k] = struct{}{}
	}
	for k,_ := range r {
		q[k] = struct{}{}
	}
	return q
}

/*
 returns the intersection of s and r
*/
func (s IntSet) Intersect(r IntSet) IntSet {
	q:= NewIntSet()
	for k, _ := range s{
		if r.Contains(k){
			q[k] = struct{}{}
		}
	}
	return q
}

/*
 returns the set difference of s and r
*/
func (s IntSet) Diff(r IntSet) IntSet {
	q := NewIntSet()
	for k, _ := range s{
		if !r.Contains(k){
			q[k] = struct{}{}
		}
	}
	return q
}

/*
 returns the symmetric set difference of s and r
*/
func (s IntSet) SymmDiff(r IntSet) IntSet {
	q := NewIntSet()
	for k, _ := range s{
		if !r.Contains(k){
			q[k] = struct{}{}
		}
	}
	for k, _ := range r{
		if !s.Contains(k){
			q[k] = struct{}{}
		}
	}
	return q
}
/*
 stringer function for an IntSet
*/
func (s IntSet) String() string {
	var buff bytes.Buffer
	buff.WriteString("{")
	for k, _ := range s{
		buff.WriteString(strconv.Itoa(k) + ", ")
	}
	buffString := buff.String()
	buffString = strings.TrimRight(buffString, ", ")
	return buffString + "}"
}
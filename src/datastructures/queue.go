package datastructures

type Queue interface {
	Push(value int) int
	Pop() (int, bool)
}
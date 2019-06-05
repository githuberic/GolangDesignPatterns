package IteratorPattern

type SomeSlice interface{}

type Iterator interface {
	HasNext() bool
	Next() SomeSlice
}

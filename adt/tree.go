package main

import (
	"fmt"
	"log"
	"math"
)

// Based on Haskell definition:
//
// data Tree = Empty
//      | Leaf Int
//      | Node Tree Tree
//
// depth :: Tree -> Int
// depth Empty = 0
// depth (Leaf n) = 1
//  depth (Node l r) = 1 + max (depth l) (depth r)
//

// Initially Tree was just:
// type Tree interface{}
//
// This is kind of like Java Object
// a type that matches any type. So
//
// we add a method to the Tree
//
type Tree interface {
	isTree()
}

type Empty struct{}
type Leaf struct {
	v int
}
type Node struct {
	left, right Tree
}

// So now for these to be seen as a type you
// implement the methods.
func (_ Empty) isTree() {}
func (_ Leaf) isTree()  {}
func (_ Node) isTree()  {}

// This allows us to implment a Tree method
// that accepts all of them.
func depth(t Tree) int {
	switch nt := t.(type) {
	case Empty:
		return 0
	case Leaf:
		return 1
	case Node:
		return 1 + int(math.Max(float64(depth(nt.left)), float64(depth(nt.right))))
	default:
		log.Fatalf("unexpected type %T", nt)
	}

	return 0
}

func main() {
	root := Node{
		left: Node{
			left: Empty{},
			right: Node{
				left:  Leaf{v: 16},
				right: Leaf{v: 101},
			},
		},
		right: Node{
			left: Node{
				left:  Leaf{v: 100},
				right: Node{left: Empty{}, right: Leaf{v: 6666}},
			},
			right: Node{
				left:  Empty{},
				right: Leaf{v: 56},
			},
		},
	}

	fmt.Printf("The tree depth %v \n\n %+v", depth(root), root)
}

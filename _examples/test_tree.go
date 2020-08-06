package main

import (
	"fmt"
	"github.com/shengjh/keydb"
)

func main()  {

	t := &keydb.Tree{}
	t.Insert([]byte("1"), []byte("1"))
	t.Insert([]byte("2"), []byte("2"))
	t.Insert([]byte("3"), []byte("3"))
	t.Insert([]byte("0"), []byte("0"))
	t.Insert([]byte("4"), []byte("4"))

	re := t.FindNodes([]byte("0"), []byte("3"))
	for _, p := range re{
		fmt.Printf("%s:%s ", p.Key, p.Value)
	}

	print("\n")
	re = t.FindNodesReverse([]byte("0"), []byte("3"))
	for _, p := range re{
		fmt.Printf("%s:%s ", p.Key, p.Value)
	}
}

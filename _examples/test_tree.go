package main

import "github.com/shengjh/keydb"

func main()  {

	t := &keydb.Tree{}
	t.Insert([]byte("1"), []byte("1"))
	t.Insert([]byte("2"), []byte("2"))
	t.Insert([]byte("3"), []byte("3"))
	t.Insert([]byte("0"), []byte("0"))
	t.Insert([]byte("4"), []byte("4"))

	t.BfsDump()
	t.BfsDump1()
}

package keydb

import (
	"fmt"
	"testing"
)

func TestMultiSegment(t *testing.T) {
	m1 := newMemorySegment(DefaultKeyCompare{})
	for i := 0; i < 100000; i++ {
		m1.Put([]byte(fmt.Sprint("mykey", i)), []byte(fmt.Sprint("myvalue", i)))
	}
	m2 := newMemorySegment(DefaultKeyCompare{})
	for i := 100000; i < 150000; i++ {
		m2.Put([]byte(fmt.Sprint("mykey", i)), []byte(fmt.Sprint("myvalue", i)))
	}

	ms := newMultiSegment([]segment{m1, m2}, nil, m1.getKeyCompare())
	itr, err := ms.Lookup(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for {
		_, _, err := itr.Next()
		if err != nil {
			break
		}
		count++
	}
	if count != 150000 {
		t.Fatal("incorrect count", count)
	}

}

func TestMultiSegment2(t *testing.T) {
	m1 := newMemorySegment(DefaultKeyCompare{})
	for i := 0; i < 1; i++ {
		m1.Put([]byte(fmt.Sprint("mykey", i)), []byte(fmt.Sprint("myvalue", i)))
	}
	m2 := newMemorySegment(DefaultKeyCompare{})
	for i := 1; i < 150000; i++ {
		m2.Put([]byte(fmt.Sprint("mykey", i)), []byte(fmt.Sprint("myvalue", i)))
	}

	ms := newMultiSegment([]segment{m1, m2}, nil, m1.getKeyCompare())
	itr, err := ms.Lookup(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for {
		_, _, err := itr.Next()
		if err != nil {
			break
		}
		count++
	}
	if count != 150000 {
		t.Fatal("incorrect count", count)
	}

}
package ds

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

/*
	functional test
*/

func TestSkipList_Get(t *testing.T) {
	sl := NewSkipList()
	n := sl.Get([]byte("name"))
	assert.Nil(t, nil, n)
}

func TestSkipList_Put(t *testing.T) {
	sl := NewSkipList()
	sl.Put([]byte("name"), "zhang san")
}

func TestSkipList_Remove(t *testing.T) {
	sl := NewSkipList()
	sl.Remove([]byte("name"))
}

func TestSkipList_1(t *testing.T) {
	sl := NewSkipList()

	sl.Put([]byte("name"), "zhang san")
	v := sl.Get([]byte("name"))

	assert.Equal(t, "zhang san", v)
}

func TestSkipList_2(t *testing.T) {
	sl := NewSkipList()

	sl.Put([]byte("name"), "zhang san")
	sl.Put([]byte("name"), "li si")
	v := sl.Get([]byte("name"))

	assert.Equal(t, "li si", v)
}

func TestSkipList_3(t *testing.T) {
	sl := NewSkipList()

	sl.Put([]byte("name"), "zhang san")
	sl.Remove([]byte("name"))
	v := sl.Get([]byte("name"))

	assert.Nil(t, nil, v)
}

// exception test
func TestSkipList_4(t *testing.T) {
	sl := NewSkipList()

	sl.Put([]byte("a"), "1")
	sl.Put([]byte("b"), "2")
	sl.Put([]byte("d"), "4")
	v := sl.Get([]byte("c"))

	assert.Nil(t, nil, v)
}

// test random option
func TestSkipList_5(t *testing.T) {
	sl := NewSkipList()

	// store expect value
	mp := make(map[string]int)

	// random option
	put := 0
	rem := 1

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5000; i++ {
		opt := rand.Intn(2)
		k := rand.Intn(100)
		v := rand.Intn(100)
		switch opt {
		case put:
			sl.Put([]byte(strconv.Itoa(k)), v)
			mp[strconv.Itoa(k)] = v
		case rem:
			sl.Remove([]byte(strconv.Itoa(k)))
			delete(mp, strconv.Itoa(k))
		}
	}

	// check actual value
	for k, v := range mp {
		assert.Equal(t, sl.Get([]byte(k)), v)
	}
}

func TestSkipListUse(t *testing.T) {
	sl := NewSkipList()
	sl.Put([]byte("1"), "a")
	sl.Put([]byte("2"), "b")
	sl.Put([]byte("3"), "c")
	sl.Put([]byte("4"), "d")
	sl.Put([]byte("5"), "e")
	PrintSkipList(sl)
}

func PrintSkipList(sl *SkipList) {
	for i := sl.level - 1; i >= 0; i-- {
		x := sl.header
		for x != nil {
			x = x.next[i]
			if x != nil {
				fmt.Printf("%v-%v ", string(x.key), x.value)
			}
		}
		fmt.Println()
	}
}

/*
	benchmark test
*/

func PrepareSL() *SkipList {
	sl := NewSkipList()

	put := 0
	rem := 1

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5000; i++ {
		opt := rand.Intn(2)
		k := rand.Intn(100)
		v := rand.Intn(100)
		switch opt {
		case put:
			sl.Put([]byte(strconv.Itoa(k)), v)
		case rem:
			sl.Remove([]byte(strconv.Itoa(k)))
		}
	}

	return sl
}

//goos: darwin
//goarch: arm64
//pkg: CaskDB/ds
//BenchmarkSkipList_Get-8         18686070                61.38 ns/op            7 B/op          0 allocs/op
//PASS
//ok      CaskDB/ds       1.476s

func BenchmarkSkipList_Get(b *testing.B) {
	sl := PrepareSL()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sl.Get([]byte(strconv.Itoa(i)))
	}
}

//goos: darwin
//goarch: arm64
//pkg: CaskDB/ds
//BenchmarkSkipList_Put-8            66636             17928 ns/op             101 B/op          4 allocs/op
//PASS
//ok      CaskDB/ds       1.613s

func BenchmarkSkipList_Put(b *testing.B) {
	sl := PrepareSL()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sl.Put([]byte(strconv.Itoa(i)), i)
	}
}

//goos: darwin
//goarch: arm64
//pkg: CaskDB/ds
//BenchmarkSkipList_Remove-8      39297234                31.07 ns/op            7 B/op          0 allocs/op
//PASS
//ok      CaskDB/ds       1.511s

func BenchmarkSkipList_Remove(b *testing.B) {
	sl := PrepareSL()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sl.Remove([]byte(strconv.Itoa(i)))
	}
}

package linkedlist

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	list := New(1, 2, 3, 4, nil, 123, "哈哈")
	fmt.Println(list)
}

func TestList_Add(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	list.Add(6)
	list.Add(6)
	list.Add(6)
	fmt.Println(list)
	list = New()
	list.Add(1, 2, 3, 4, 5)
	fmt.Println(list)
}

func TestList_Prepend(t *testing.T) {
	list := New(1)
	list.Prepend(-2, -1, 0)
	list.Add(2, 3)
	list.Prepend(-3)
	fmt.Println(list)
}

func TestList_Iterator(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	//list := New()
	iterator := list.Iterator()
	for iterator.Next() {
		fmt.Println(iterator.Value())
	}
	for iterator.Prev() {
		fmt.Println(iterator.Value())
	}
}

func TestList_Map(t *testing.T) {
	list := New(1, 2, 3, "4", nil)
	list.Map(func(v interface{}) interface{} {
		switch v.(type) {
		case int:
			return v.(int) + 1
		case string:
			return fmt.Sprintf("%s1", v.(string))
		case nil:
			return v
		default:
			return v
		}
	})
	fmt.Println(list)
}

func TestList_Set(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	//list := New()
	list.Set(4, 100)
	fmt.Println(list)
	list.Set(4, 10)
	fmt.Println(list)
	list.Set(4, 1)
	fmt.Println(list)
}

func TestList_Insert(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	//list := New()
	list.Insert(3, 100)
	list.Insert(3, 10)
	list.Insert(3, 1)
	fmt.Println(list)
}

func TestList_Get(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	fmt.Println(list.Get(-1))
	fmt.Println(list.Get(1))
	fmt.Println(list.Get(2))
	fmt.Println(list.Get(5))
}

func TestList_Remove(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	list.Remove(2)
	list.Remove(2)
	list.Remove(2)
	list.Remove(0)
	list.Prepend(11, 12, 13, 14, 15)
	fmt.Println(list, list.Size())
}

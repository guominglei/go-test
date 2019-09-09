//package main
package lru

import "fmt"
import "math/rand"

// data

type Item struct {
	pre       *Item
	data      int
	cache_key string
	after     *Item
}

type Link struct {
	// 头指针
	head *Item
	// 尾指针
	tail *Item
	// 当前数量
	num int
	// 容量
	total int
	// 字典
	dict map[string]Item
}

type LinkInterface interface {
	add()
	deleted()
}

func (link *Link) add(item *Item) {
	if link.head == nil {
		link.head = item
		link.tail = item
	} else {
		link.tail.after = item
		item.pre = link.tail
		link.tail = item
	}
	link.num += 1
}

func (link *Link) deleted(item *Item) {
	pre_item := item.pre
	af_item := item.after
	if pre_item != nil {
		if af_item != nil {
			pre_item.after = af_item
			af_item.pre = pre_item
		} else {
			pre_item.after = nil
			link.tail = pre_item
		}
	} else {
		if af_item != nil {
			link.head = af_item
			af_item.pre = nil
		} else {
			link.head = nil
			link.tail = nil
		}
	}
	link.num -= 1
}

type LinkLruInterface interface {
	get_item()
}

func (lru *Link) get_item(key string) (item *Item) {
	cache_item, is_exists := lru.dict[key]
	if !is_exists {
		/*
			未命中
			1、获取新数据。
			2、判断是否有空间。
				有空间直接添加
				没有空间删除头，添加新的
		*/
		value_random := rand.Intn(100)
		new_item := Item{data: value_random, cache_key: key}
		if lru.num < lru.total {
			// 有空间
			fmt.Print("未命中 有空间", "\n")
			lru.add(&new_item)
			lru.dict[key] = new_item
		} else {
			// 没有空间
			fmt.Print("未命中 没有空间", "\n")
			head_cache_key := lru.head.cache_key
			lru.deleted(lru.head)
			delete(lru.dict, head_cache_key)
			lru.add(&new_item)
			lru.dict[key] = new_item
		}
		return &new_item

	} else {
		/*
			命中
			删除链接
			添加到队尾
		*/
		fmt.Print("命中", "\n")
		lru.deleted(&cache_item)
		lru.add(&cache_item)
		return &cache_item
	}
}

func main_() {

	link := Link{}

	item1 := Item{data: 1}
	link.add(&item1)
	fmt.Printf("%d \n", link.num)
	fmt.Print(&link.head, &link.tail, "\n")
	item2 := Item{data: 2}
	link.add(&item2)
	fmt.Printf("%d \n", link.num)
	item3 := Item{data: 3}
	link.add(&item3)
	fmt.Printf("%d \n", link.num)
	item4 := Item{data: 4}
	link.add(&item4)
	fmt.Printf("%d \n", link.num)

	// 遍历1
	dt := link.head
	for {
		fmt.Printf("for %d \n", dt.data)
		dt = dt.after
		if dt == nil {
			break
		}
	}
	// 遍历2
	for it := link.head; it != nil; it = it.after {
		fmt.Println(it.data, it.cache_key)
	}
	new_link := Link{total: 3, dict: make(map[string]Item)}

	//new_link := Link{total: 3, dict: map[string]Item{}}
	cache_1 := new_link.get_item("1")
	fmt.Print("1\t", cache_1.data, "\t", new_link.num, "\n")
	cache_2 := new_link.get_item("2")
	fmt.Print("2\t", cache_2.data, "\t", new_link.num, "\n")
	cache_3 := new_link.get_item("3")
	fmt.Print("3\t", cache_3.data, "\t", new_link.num, "\n")
	cache_4 := new_link.get_item("4")
	fmt.Print("4\t", cache_4.data, "\t", new_link.num, "\n")
	cache_5 := new_link.get_item("3")
	fmt.Print("3\t", cache_5.data, "\t", new_link.num, "\n")

	cache_6 := new_link.get_item("1")
	fmt.Print("1\t", cache_6.data, "\t", new_link.num, "\n")
}

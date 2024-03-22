//map是一种无序的基于key-value的数据结构，是引用类型，必须初始化才能使用

//map[KeyType]ValueType

package main

import (
	//"crypto/rand"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func eachMap() {
	scoreMap := make(map[string]int)
	scoreMap["zhangsan"] = 90
	scoreMap["lisi"] = 100
	scoreMap["wangwu"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}

func eachMapK() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	scoreMap["王五"] = 60
	for k := range scoreMap {
		fmt.Println(k)
	}
}

//delte(map,key)

func delMapForKey() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	scoreMap["王五"] = 602
	delete(scoreMap, "张三")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

}

// 按照指定顺序遍历map

func sortEachMap() {
	rand.Seed(time.Now().UnixNano()) // Initialize random number seed

	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

}

// 元素为map类型的切片
func mapSlice() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

}

func mapSlice2() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)

}

func main() {

	userInfo := map[string]string{
		"username": "derick",
		"password": "123456",
	}
	fmt.Println(userInfo)

	scoreMap := make(map[string]int, 8)
	scoreMap["zhangsan"] = 90
	scoreMap["xiaoming"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["zhangsan"])
	fmt.Println("type of a:%T\n", scoreMap)

	v, ok := scoreMap["zhangsan"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	mapSlice2()
	mapSlice()
	sortEachMap()
	delMapForKey()
	eachMapK()
	eachMap()

}

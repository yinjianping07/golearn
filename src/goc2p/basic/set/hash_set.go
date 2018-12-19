package set

import (
	"bytes"
	"fmt"
)

//利用map实现Hashset

type  HashSet struct {
	//利用map m中的key值来存储HashSet的元素值，选用了值占用空间小的bool类型来作为map m的value有以下三个原因：
	//1、bool类型值占用空间最小（之一），只占用一个字节
	//2、bool类型只有两个true/false，只两个值都是预定义常量
	//3、更加有利于判断map m中是否存在某一键。
	m map[interface{}]bool
}

//初始化HashSet类型的函数
func NewHashSet() *HashSet{
	//make分配内存空间，因为直接new一个HashSet类的对象的话，map的零值为nil
	//返回值类型为*HashSet，是因为在后面的方法在调用时，指针类型占用空间小
	return &HashSet{m: make(map[interface{}]bool)}
}

//Add添加元素
func (set *HashSet) Add(e interface{}) bool{
	//判断map中是否存在
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

//Remove删除函数
func (set *HashSet) Remove(e interface{}){
	delete(set.m,e)
}

//Clear清楚所有
//如果迭代去删除所有明显是不可行的，所以直接赋值一个空的map
//如果在这里接收者类型是 HashSet，那么在这里操作的是当前值的某个复制品中的字段m而已
func (set *HashSet) Clear(){
	set.m = make(map[interface{}]bool)
}

//Contains函数
func (set *HashSet) Contains(e interface{}) bool{
	return set.m[e]
}

//Len函数
func (set *HashSet) Len()int{
	return len(set.m)
}

//Same判断两个HashSet是否相同
func (set *HashSet) Same(other *HashSet) bool{
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m{
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

//Elements方法，返回一个元素数组
func (set *HashSet) Elements() []interface{}{
	initialLen := len(set.m)
	snapshot := make([]interface{},initialLen)
	actualLen := 0
	for key := range set.m{
		if actualLen<initialLen {
			snapshot[actualLen] = key
		}else {
			snapshot = append(snapshot,key)
		}
		actualLen++
	}
	if actualLen<initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

//String方法
func (set *HashSet) String()string{
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v",key))
	}
	buf.WriteString("}")
	return buf.String()
}

//高级功能
//判断set A 是否真包含set B --真集
func (set *HashSet) IsSuperSet(other *HashSet) bool{
	if other == nil {
		return false
	}
	oneLen := set.Len()
	otherLen := other.Len()
	if oneLen == 0 || oneLen == otherLen {
		return false
	}
	for _,v := range other.Elements(){
		if !set.Contains(v){
			return false
		}
	}
	return true
}

//并集
func (set *HashSet) Union(other *HashSet)(unionSet *HashSet){
	if other != nil {
		for _,v := range other.Elements(){
			set.Add(v)
		}
	}
	return set
}

//求交集
func (set *HashSet) Intersect(other *HashSet) *HashSet{
	if other == nil {
		return other
	}
	temp := set.Elements()
	tempMap := NewHashSet()
	for _,v := range temp{
		if other.Contains(v){
			tempMap.Add(v)
		}
	}
	return tempMap
}

//补集，属于A，但不属于B的
func (set *HashSet) Difference(other *HashSet)*HashSet{
	if set.Intersect(other) == nil{ //当other为空或者set和other无交集，返回set
		return set
	}
	//tempMap := set  //这里直接是指针上来，还是会修改set的值
	tempMap := NewHashSet()//复制一份set到tempMap
	for _,v := range set.Elements(){
		tempMap.Add(v)
	}
	for _,v := range set.Intersect(other).Elements() {
		if tempMap.Contains(v) {
			tempMap.Remove(v)
		}
	}
	return tempMap
}

//对称差集
func (set *HashSet) SymmetricDifference(other *HashSet)*HashSet{
	if set == nil && other == nil {
		return other
	}
	return set.Difference(other).Union(other.Difference(set))
}
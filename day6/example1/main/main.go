package main

import (
	"sort"
	"math/rand"
	"fmt"
)

type Student struct{
	Name string
	Age uint16
	Score float32
}

/*
type Interface interface {
        // Len is the number of elements in the collection.
        Len() int
        // Less reports whether the element with
        // index i should sort before the element with index j.
        Less(i, j int) bool
        // Swap swaps the elements with indexes i and j.
        Swap(i, j int)
}*/

type StudentArray []Student

func (stus StudentArray) Len() int{
	return len(stus)
}


func (stus StudentArray) Less(i, j int) bool{
	//根据分数从大到小
	return stus[i].Age>stus[j].Age
}


func (stus StudentArray) Swap(i, j int){
	stus[i],stus[j] = stus[j],stus[i] 
}

func main(){
	var stus StudentArray
	for i := 0; i < 10; i++ {
		var stu Student = Student{
			Name:fmt.Sprintf("stu%d", rand.Intn(100)),
			Age:uint16(rand.Intn(99)),
			Score:rand.Float32(),
		}
		stus = append(stus, stu)
	}
	for _,v:=range stus{
		fmt.Println(v)
	}
	sort.Sort(stus)
	fmt.Printf("\n\n")
	for _,v:=range stus{
		fmt.Println(v)
	}
}
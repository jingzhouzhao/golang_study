package main

import (
	"fmt"
)
//冒泡排序（从小到大）
func BubbleSort(arr []int) []int{
	fmt.Println("排序前：",arr)
	for i := 0; i < len(arr)-1; i++ {
		fmt.Println("i:",i)
		for j := 0; j < len(arr)-i-1 ; j++ {
			fmt.Printf("\t (%d)-(%d) | [%d] - [%d] \n",j,j+1,arr[j],arr[j+1])
			if arr[j] > arr[j+1]{
				arr[j+1],arr[j]=arr[j],arr[j+1]
			}
		}
		fmt.Printf("第(%d)轮结果:%v \n",i+1,arr)
	}
	fmt.Println("排序后：",arr)
	return arr
}

//选择排序（从小到大）（核心思想：每次选择一个最大或者最小的值）
func SelectionSort(arr []int) []int{
	for i := 0; i < len(arr)-1; i++ {
		var min int=i
		for j := i+1; j < len(arr); j++ {
			if arr[min]> arr[j]{
				min = j
			}
		}
		arr[min],arr[i] =arr[i],arr[min]
	}
	return arr
}

//插入排序（重新到大）（核心思想，与已排序的序列进行对比，默认第一个元素为已排序）
func InsertionSort(arr []int) []int{
	for i := 1; i < len(arr); i++ {
		for j := i-1; j >=0; j-- {
			if arr[j]>arr[j+1] {
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}else{
				break;
			}
		}
	}
	return arr
}

//快速排序(核心思想：分而治之)
func QuickSort(arr []int,left int, right int) []int{
	if left>=right{
		return arr
	}
	val:=arr[left]
	k:=left
	for i := left+1; i <= right; i++ {
		if arr[i]<val{
			arr[k],arr[i] = arr[i],arr[k+1]
			k++
		}
	}
	arr[k] = val
	QuickSort(arr,left,k-1)
	QuickSort(arr,k+1,right)
	return arr
}

func main(){
	//var arr = [...]int{6,1,5,3,4,7,9,8,0}
	//BubbleSort(arr[:])
	//fmt.Println(QuickSort(arr[:],0,len(arr)-1))
	test()
}

func test(){

	//1 1 2 3 5 8 13 21 34 55

	a:=1
	b:=1
	for i := 0; i < 10; i++ {
		fmt.Println(a)
		a,b=b,a+b
	}
}
package test

import(
	t2 "go_dev/day5/7_struct_method/test2"
	"fmt"
)

func Test(){
	//所有的type访问都需要加上包名，例如以下代码无法直接访问到
	//var stu Student
	//使用默认包名访问
	//var stu test2.Student
	//也可以给导入的包取别名
	var stu t2.Student
	stu.SetName1("只有大写开头的方法可以访问，说明golang通过首字母大小写来区分访问权限")
	fmt.Println(stu)
	
}
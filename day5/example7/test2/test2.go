package test2

type Student struct{
	name string
}

func (stu *Student) setName(name string){
	stu.name = name
}
func (stu *Student) SetName1(name string){
	stu.name = name
}
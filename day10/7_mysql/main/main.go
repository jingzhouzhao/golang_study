package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

type Test struct{
	Id int `db:"id"`
	UserName string `db:"user_name"`
	Age int `db:"age"`
	Email string `db:"email"`
}

func main(){
	db,_:=sqlx.Connect("mysql", "root:123456@tcp(localhost:3306)/go_test")
	tx:=db.MustBegin()
	tx.MustExec("insert into test (user_name,age,email) values (?,?,?)", "张三",18,"zhangsan@163.com")
	tx.MustExec("insert into test (user_name,age,email) values (?,?,?)", "李四",20,"lisi@163.com")
	tx.NamedExec("insert into test (user_name,age,email) values (:user_name,:age,:email)",&Test{UserName:"王五",Age:22,Email:"wangwu@qq.com"})
	tx.Commit()
	test:=[]Test{}
	db.Select(&test, "select * from test order by user_name")
	fmt.Printf("%#v\n",test)

	zhangsan:=Test{}
	db.Get(&zhangsan, "select * from test where user_name=?", "张三")
	fmt.Printf("%#v\n", zhangsan)
	r,err:=db.NamedExec("update test set age=:age where user_name=:user_name", Test{UserName:"张三",Age:38})
	if err!=nil{
		fmt.Println("updated failed:",err)
		return
	}
	ra,_:=r.RowsAffected()
	fmt.Printf("update affected Rows:%d\n",ra)
	r,err=db.Exec("delete from test where user_name =?", "王五")
	if err!=nil{
		fmt.Println("deleted failed:",err)
		return
	}
	ra,_=r.RowsAffected()
	fmt.Printf("delete affected Rows:%d",ra)
}
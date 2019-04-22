package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	//sql.DB是作为长连接使用，不要频繁open，close
	//如果需要短连接，可以将DB传入function
	//db, err := sql.Open("mysql", "用户名:密码@tcp(IP:端口)/数据库?charset=utf8")
	//db, err := sql.Open("mysql", "root:09274137@tcp(127.0.0.1:3306)/test?charset=utf8")
	db, err := sql.Open("mysql", "root:09274137@/test")
	checkerr(err)
	defer db.Close()
	//验证连接
	err = db.Ping()
	checkerr(err)
	//插入
	stmt, err := db.Prepare("insert into student(name) values(?) where id=?;")
	result, err := stmt.Exec("wa",3)
	//可以直接使用exec
	//result, err := db.Exec("insert into student(name) values(?) where id=?;","wa",3)
	checkerr(err)
	defer stmt.Close()
	//获取插入id
	lastinsertid, err := result.LastInsertId()
	checkerr(err)
	fmt.Println(lastinsertid)
	//获取影响数据库的行数，可以根据该数值判断是否插入或删除或修改成功
	//count, err := result.RowAffected()

}

func checkerr(err error){
	if err != nil{
		panic(err.Error())
	}
}
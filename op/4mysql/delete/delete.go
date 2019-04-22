package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql","root:09274137@/test")
	checkerr(err)
	defer db.Close()
	err = db.Ping()
	checkerr(err)
	//删除
	stmt, err := db.Prepare("delete from student where id=?;")
	result, err := stmt.Exec(3)
	checkerr(err)
	defer stmt.Close()
	//获取删除的记录数量
	rowsaffect, err := result.RowsAffected()
	checkerr(err)
	fmt.Println(rowsaffect)
}

func checkerr(err error){
	if err != nil{
		panic(err.Error())
	}
}
package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db,err := sql.Open("mysql", "root:09274137@/test")
	checkerr(err)
	//err = db.Ping()
	//checkerr(err)

	//更新
	stmt, err := db.Prepare("update student set sname=? where id=?")
	result, err := stmt.Exec("aw",2)
	//result, err := db.Exec("update student set name=? where id=?;", "aw", 1)
	checkerr(err)
	defer stmt.Close()
	//获取更新的记录数量
	rowsaffect, err := result.RowsAffected()
	checkerr(err)
	fmt.Println(rowsaffect)
}
func checkerr(err error){
	if err != nil{
		panic(err.Error())
	}
}
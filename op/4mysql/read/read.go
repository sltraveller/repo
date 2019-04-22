package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:09274137@tcp(127.0.0.1:3306)/test")
	checkerr(err)
	defer db.Close()

	err = db.Ping()
	checkerr(err)

	//查询
	rows, err := db.Query("select * from student;")
	checkerr(err)
	//columns, err := rows.Columns()
	//checkerr(err)
	//fmt.Println(columns)

	for rows.Next(){
		var name string
		var id int
		//和
		err = rows.Scan(&id, &name)
		checkerr(err)
		fmt.Println(name)
	}
	defer rows.Close()
}
/*
func main(){
	db, err := sql.Open("mysql", "root:09274137@/test")
	checkerr(err)
	defer db.Close()

	err = db.Ping()
	checkerr(err)

	//查询
	rows, err := db.Query("select * from student;")
	checkerr(err)
	columns, err := rows.Columns()
	checkerr(err)
	fmt.Println(columns)
	scanargs := make([]interface{}, len(columns))
	values := make([]interface{},len(columns))
	for i := range values{
		scanargs[i] = &values[i]
	}

	for rows.Next(){
		err = rows.Scan(scanargs...)
		checkerr(err)
		record := make(map[string]string)
		for i,col := range values{
			if col != nil{
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

func main(){
	db, err := sql.Open("mysql", "root:09274137@/test")
	checkerr(err)
	err = db.Ping()
	checkerr(err)
	stmt, err := db.Prepare("select * from student where id=?;")
	checkerr(err)
	var id int
	var name string
	err = stmt.QueryRow(3).Scan(&id,&name)
	checkerr(err)
	fmt.Println(id,name)
}
*/
func checkerr(err error){
	if err != nil{
		panic(err.Error())
	}
}
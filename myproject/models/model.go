package models
import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
外键始终在子表上
使用标签`orm:"column(id)"`对属性进行标注，用于解析
自动建表不会生成外键
`orm:"rel(one)"` 表示one2one
`orm:"rel(fk)"`  表示one2many
`orm:"rel(m2m)"` 表示many2many
`orm:"reverse(one)"` 表示one2one反向关系
`orm:"reverse(many)"`表示one2many反向关系
*/

//User is
type User struct {
	Id	int
	Name	string
	Profile	*Profile `orm:"rel(one)"`
	Post	[]*Post	`orm:"reverse(many)"`
}

//Profile is
type Profile struct {
	Id	int
	Age	int16
	User *User `orm:"reverse(one)"`
}

//Post is
type Post struct {
	Id int
	Title string
	User *User `orm:"rel(fk)"`
	Tags []*Tag `orm:"rel(m2m)"`
}

//Tag is
type Tag struct {
	Id int
	Name string
	Post []*Post `orm:"reverse(many)"`
}


//Up is
type Up struct {
	Id int
	Name string
}
func init(){
	//设置数据库信息
	orm.RegisterDataBase("default", "mysql", "root:09274137@/test")
	//最大空闲连接
	//maxIdle := 30
	//orm.SetMaxIdleConns("default", 30)
	//最大数据库连接
	//maxConn := 30
	//orm.SetMaxOpenConns("default", 30)
	//orm.RegisterDataBase("default", "mysql", "root:09274137@/test", maxIdle, maxConn)


	//映射model
	orm.RegisterModel(new(User), new(Profile), new(Post), new(Tag), new(Up))
	//使用前缀，创建后为prefix_user
	//orm.RegisterModelWithPrefix("prefix_", new(User))


	//生成表 如果创建表有错需要更改，false->true， 第三个参数是显示创建表过程
	orm.RunSyncdb("default", false, true)
}
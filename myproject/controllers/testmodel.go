package controllers

import (
	"github.com/astaxie/beego/logs"
	"myproject/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)


//ModelController is
type ModelController struct {
	beego.Controller
}

/*
//CreateModel is
func (c *ModelController) CreateModel() {
	//orm.RunSyncdb("default", false, true)
	datainit()
}

func datainit() {
	o := orm.NewOrm()
	//rel 自动生成外键为表明_id

	sql1 := "insert into user (name,profile_id) values ('hanru',1),('ruby',2),('王二狗',3);"
	sql2 := "insert into profile (age) values (20),(19),(21);"
    sql3 := "insert into tag (name) values ('offical'),('beta'),('dev');"
    sql4 := "insert into post (title,user_id) values ('paper1',1),('paper2',1),('paper3',2),('paper4',3),('paper5',3);"
    // m2m 生成的 表名：子表_主表s  主键自增
    sql5 := "insert into post_tags (tag_id, post_id) values (1,1),(1,3),(2,2),(3,3),(2,4),(3,4),(3,5); "

    //使用Raw（）.Exec（）执行sql
    o.Raw(sql1).Exec()
    o.Raw(sql2).Exec()
    o.Raw(sql3).Exec()
    o.Raw(sql4).Exec()
    o.Raw(sql5).Exec()
}
*/

//Get is
func (c *ModelController) Get() {
	o := orm.NewOrm()
	/*
	//插入
	//orm对象
	o := orm.NewOrm()
	//一个要插入数据的结构体对象
	user := models.User{}
	//对结构体赋值
	user.Name="111"
	user.Pwd="222"
	//插入
	_, err := o.Insert(&user)
	if err != nil{
		logs.Info("插入失败", err)
	}
	*/
	// up := models.Up{}
	// up.Name = "qw"
	// id, err := o.Insert(&up)
	// if err != nil {
	// 	logs.Info(err)
	// }
	

	/*
	//查询
	//orm对象
	o := orm.NewOrm()
	//查询对象
	user := models.User{}
	//指定查询对象字段值
	user.Id = 1
	user.Name = "111"
	//查询
	err = o.Read(&user, "Name")
	if err != nil {
		logs.Info("查询失败", err)
		return
	}
	logs.Info("查询成功", user)
	*/
	// up := models.Up{Id:6}
	// //也可以在这里指定up.Id = 6
	// err := o.Read(&up)
	// if err != nil {
	// 	logs.Info("1")
	// }
	// c.Ctx.WriteString(up.Name)

	 

	/*
	//更新
	//orm对象
	o := orm.NewOrm()
	//查找需要更新的对象
	user := models.User{}
	//重新赋值
	user.Id = 1
	err := o.Read(&user)
	//更新
	if err == nil {
		user.Name = "222"
		user.Pwd = "333"
		_, err = o.Update(&user)
		if err != nil {
			logs.Info("更新失败",err)
			return
		}
	}
	*/
	// id, err := o.Update(&models.Up{Name:"wa"})
	// if err != nil {
	// 	logs.Info(err)
	// }
	

	/*
	//删除
	//orm对象
	o := orm.NewOrm()
	//删除的对象
	user := models.User{}
	//指定数据
	user.Id = 5
	//删除
	_, err := o.Delete(&user)
	if err != nil {
		logs.Info("删除错误", err)
		return
	}
	*/
	// id, err := o.Delete(&models.Up{Id:5})
	// if err != nil {
	// 	logs.Info(err)
	// }
	

	//如果不存在则创建
	if created, id, err := o.ReadOrCreate(&models.Up{Name:"dd"}, "Name"); err ==nil {
		if created {
			c.Ctx.WriteString(fmt.Sprintf("created:%t,id:%d",created,id))
		}else{
			c.Ctx.WriteString("you")
		}
	}else{
		logs.Info("q")
	}

	//c.Ctx.WriteString(fmt.Sprintf("id:%d",id))
	
	//获取信息
	// dr := o.Driver()
	// fmt.Println(dr.Name() == "default")
	// fmt.Println(dr.Type() == orm.DRMySQL)
}
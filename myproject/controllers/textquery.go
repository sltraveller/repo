package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"myproject/models"
)


//Query is
func (c *ModelController) Query() {
	o := orm.NewOrm()
	//获取queryseter对象
	q := o.QueryTable("user")
	//也可以这样获取
	//user := new(User)
	//q = o.QueryTable(user)
	var users []*models.User
	num, err:=q.All(&users)
	if err != nil {
        // 处理err
    }
    c.Ctx.WriteString("<html>" + fmt.Sprintf("共查询了 num:%d 条数据。。", num) + "<br/><br/>")
    c.Ctx.WriteString("<table border='1' width='50%' cellspacing='0'>")
    c.Ctx.WriteString("<th>Id</th><th>Name</th><th>profile_id</th>")

    for _, user := range users {
        c.Ctx.WriteString("<tr>" +
            "<td>" + fmt.Sprintf("%v", user.Id) + "</td>" +
            "<td>" + fmt.Sprintf("%v", user.Name) + "</td>" +
            "<td>" + fmt.Sprintf("%v", user.Profile.Id) + "</td>" +
            "</tr>")

    }
    c.Ctx.WriteString("</table></html>")
	
}
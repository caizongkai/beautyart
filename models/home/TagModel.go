package home

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "github.com/hunterhug/beautyart/lib"
	"strconv"
	"strings"
)

//标签表
type Tag struct {
	Id    int64
	Name  string `orm:"size(20);index"`
	Count int64
}

func init() {
	orm.RegisterModelWithPrefix(beego.AppConfig.String("db_prefix"), new(Tag))
}

func (m *Tag) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

//删除标签，先将文章标签清空，再删除文章-标签记录，最后删除标签
func (m *Tag) Delete() error {
	var list []*TagPost
	table := new(Post).TableN()
	new(TagPost).Query().Filter("tagid", m.Id).All(&list)
	if len(list) > 0 {
		ids := make([]string, 0, len(list))
		for _, v := range list {
			ids = append(ids, strconv.FormatInt(v.Postid, 10))
		}
		orm.NewOrm().Raw("UPDATE "+table+" SET tags = REPLACE(tags, ?,',') WHERE id IN ("+strings.Join(ids, ",")+")", ","+m.Name+",").Exec()
		new(TagPost).Query().Filter("tagid", m.Id).Delete()
	}
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

//表查询
func (m *Tag) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

//标签连接
func (m *Tag) Link() string {
	return fmt.Sprintf("<a class=\"category\" href=\"/category/%s\">%s</a>", Rawurlencode(m.Name), m.Name)
}

//更新统计
func (m *Tag) UpCount() {
	m.Count, _ = new(TagPost).Query().Filter("tagid", m.Id).Count()
	m.Update("count")
}

//合并到另一个标签
func (m *Tag) MergeTo(to *Tag) {
	var list []*TagPost
	var tp TagPost
	tp.Query().Filter("tagid", m.Id).All(&list)
	if len(list) > 0 {
		ids := make([]string, 0, len(list))
		for _, v := range list {
			ids = append(ids, strconv.FormatInt(v.Postid, 10))
		}
		tp.Query().Filter("tagid", m.Id).Update(orm.Params{"tagid": to.Id})
		orm.NewOrm().Raw("UPDATE "+new(Post).TableN()+" SET tags = REPLACE(tags, ?, ?) WHERE id IN ("+strings.Join(ids, ",")+")", ","+m.Name+",", ","+to.Name+",").Exec()
	}
}

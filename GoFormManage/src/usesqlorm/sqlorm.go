package main

import (
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
	"time"
	"fmt"
)

type Userinfo struct {
	Uid     int `PK` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username    string
	Departname  string
	Created     time.Time
}

type test_time struct {
	time time.Time
	id int32
}

func (t * test_time)String() string {
	return t.time.String()
}

func main(){

	//db, err := sql.Open("mysql", "root:2737353904@tcp(localhost:3306)/test?charset=utf8")
	db, err := sql.Open("mymysql", "test/root/2737353904")
	if err != nil {
		panic(err)
	}

	beedb.OnDebug=true

	orm := beedb.New(db)

	var t test_time

	orm.Where("id=?", 1).Find(&t)

	fmt.Println(t.time)

	//addslice := make([]map[string]interface{}, 0)
	//add := make(map[string]interface{})
	//add2 := make(map[string]interface{})
	//add["username"] = "astaxie"
	//add["departname"] = "cloud develop"
	//add["created"] = "2012-12-02"
	//add2["username"] = "astaxie2"
	//add2["departname"] = "cloud develop2"
	//add2["created"] = "2012-12-02"
	//addslice =append(addslice, add, add2)
	//fmt.Println(orm.SetTable("userinfo").InsertBatch(addslice))

	//var saveone Userinfo
	//saveone.Username = "Test Add User"
	//saveone.Departname = "Test Add Departname"
	//fmt.Println(time.Now())
	//saveone.Created = time.Now()
	//fmt.Println("Time:", saveone.Created)
	//err = orm.Save(&saveone)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(saveone.Uid)
}
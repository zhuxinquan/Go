package main

import (
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
	"time"
	"fmt"
	"strconv"
	"reflect"
	"strings"
	"unsafe"
	"util"
)

type Userinfo struct {
	Uid     int `PK` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username    string
	Departname  string
	Created     time.Time
}

func (u Userinfo)String() string{
	return strconv.Itoa(u.Uid)
}

type Tt struct {
	Id int `PK`
	Time time.Time
}

func String(b []byte) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}


func main(){

	//db, err := sql.Open("mysql", "root:2737353904@tcp(localhost:3306)/test?charset=utf8")
	db, err := sql.Open("mymysql", "test/root/2737353904")
	if err != nil {
		panic(err)
	}

	//beedb.OnDebug=true

	orm := beedb.New(db)

	var user Tt
	//var userinfo Userinfo
	orm.Where("id=?", 1).Find(&user)
	a, _ := orm.SetTable("userinfo").SetPK("uid").Select("uid,username").FindMap()
	for _, v := range a {
		for k, value := range v {
			if(strings.EqualFold(k, "uid")) {
				fmt.Println(k, ":", util.ByteToInt(value))
			} else {
				fmt.Println(k, ":", string(value))
			}
		}
	}
}
package Db

import (
	"fmt"
	"hxsg/orm"
)

type UserDb struct {
	Uid        int    `db:uid`
	OpenId     string `db:"openId"`
	Password   string `db:"password"`
	ServerId   int    `db:"serverId"`
	Name       string `db:"name"`
	Level      int    `db:"level"`
	Avatar     string `db:"avatar"`
	Sex        int    `db:"sex"`
	Vip        int    `db:"vip"`
	Medal      int    `db:"medal"`
	JobId      int    `db:"jobId"`
	UpdateTime string `db:"updateTime"`
	CreateTime string `db:"createTime"`
}

var UserObjs = &UserDb{}

func (this *UserDb) CheckUserName(userName string) error {
	var u UserDb
	err := orm.Orm.Where("name", userName).Find(&u)
	if err != nil {
		fmt.Println("scan uid,failed , err:%v \n", err)
		return err
	}
	fmt.Printf("id:%d name:%s ", u.Uid, u.Name)
	return nil
}

func (this *UserDb) CheckOpenId(openId string) {

}

package services

import (
	"context"
	"fmt"
	"grpcServer/dbs"
	"strings"
	"time"
)

const version = "v 0.0.3"

type UserService struct {
}

func (this *UserService) GetUserList(ctx context.Context, req *UserRequest) (res *UserResponseList, err error) {
	res = new(UserResponseList)
	res.Version = version
	if req.Search == "" {
		c := &struct {
			Count int `gorm:"column:c"`
		}{}
		sqlStr := "select count(*) as c from users "
		dbs.Orm.Raw(sqlStr).Find(&c)
		fmt.Println(c.Count)
		res.Total = int32(c.Count)
		dbs.Orm.Raw("select * from users order by id limit ?,?", (req.Page-1)*req.Size, req.Size).Find(&res.Users)
	} else {
		c := &struct {
			Count int `gorm:"column:c"`
		}{}
		split := strings.Split(req.Search, ",")
		sqlStr := "select count(*) as c  from users where  "
		queryStr := ""
		for i := 0; i < len(split); i++ {
			queryStr += "name like '%" + split[i] + "%' "
			if i < len(split)-1 {
				queryStr += " or "
			}
		}
		sqlStr += queryStr + "order by id"
		dbs.Orm.Raw(sqlStr).Find(&c)
		res.Total = int32(c.Count)
		dbs.Orm.Raw("select * from users where "+queryStr+" order by id limit ?,?", (req.Page-1)*req.Size, req.Size).Find(&res.Users)
	}

	return
}

func (this *UserService) AddUsers(ctx context.Context, in *UsersInputRequest) (*Result, error) {

	us := in.Users
	for i := 0; i < len(us); i++ {

		m := us[i]
		if m.Id != 0 {
			_ = dbs.Orm.Exec("update users set name = ?, password = ? ,telephone = ?, email =? where id = ?", m.Username, m.Password, m.Tel, m.Email, m.Id).Error
		} else {
			_ = dbs.Orm.Exec("insert into users (name,password,telephone,email,create_time) values(?,?,?,?,?)", m.Username, m.Password, m.Tel, m.Email, time.Now()).Error
		}
		//if err != nil {
		//	return &Result{Success: false}, err
		//}
	}
	return &Result{Success: true, Version: version}, nil
}

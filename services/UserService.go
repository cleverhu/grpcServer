package services

import (
	"context"
	"fmt"
	"grpcServer/dbs"
	"strings"
)

type UserService struct {
}

func (this *UserService) GetUserList(ctx context.Context, req *UserRequest) (res *UserResponseList, err error) {
	res = new(UserResponseList)
	if req.Search == "" {
		c := &struct {
			Count int `gorm:"column:c"`
		}{}
		sqlStr := "select count(*) as c from user "
		dbs.Orm.Raw(sqlStr).Find(&c)
		fmt.Println(c.Count)
		res.Total = int32(c.Count)
		dbs.Orm.Raw("select * from user order by id limit ?,?", (req.CurrentPage-1)*req.PageSize, req.PageSize).Find(&res.Users)

		return
	} else {
		c := &struct {
			Count int `gorm:"column:c"`
		}{}
		split := strings.Split(req.Search, ",")
		//fmt.Println(split)
		sqlStr := "select count(*) as c  from user where  "
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
		//fmt.Println(count)
		dbs.Orm.Raw("select * from user where "+queryStr+" order by id limit ?,?", (req.CurrentPage-1)*req.PageSize, req.PageSize).Find(&res.Users)

	}


	return
}

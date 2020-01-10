package main

type Goods struct {
	Id     int64
	Name   string
	NeedJf int64
}

type UserInfo struct {
	Nickname string
	Jf       int64
	Goods    []int64
}
type AdminInfo struct {
	Nickname string
}
type QusetTwo struct {
	GoodsList []Goods
	Users     []UserInfo
	Admin     []AdminInfo
	Nowjs     int64 //1管理员2普通用户
}

func SetQtInit() *QusetTwo {
	this := new(QusetTwo)
	this.Admin = []AdminInfo{
		AdminInfo{
			Nickname: "admin001",
		},
	}
	this.Users = []UserInfo{
		UserInfo{
			Nickname: "user001",
			Jf:       100,
			Goods:    []int64{},
		},
	}
	this.GoodsList = []Goods{
		Goods{
			Id:     1,
			Name:   "cs001",
			NeedJf: 10,
		},
	}
	return this

}

//选择角色
func (qt *QusetTwo) selectjs(t int64) {
	qt.Nowjs = t
}

//查看用户信息
func (qt *QusetTwo) getusers() ([]UserInfo, string) {
	if qt.Nowjs != 1 {
		return nil, "没有权限"
	}
	return qt.Users, ""
}

//查看所以商品
func (qt *QusetTwo) getgoods() []Goods {
	return qt.GoodsList
}

//设置奖品内容和积分
func (qt *QusetTwo) setgoods(goodsinfo Goods) {
	idx := len(qt.GoodsList)
	goodsinfo.Id = int64(idx) + 1
	qt.GoodsList = append(qt.GoodsList, goodsinfo)
}

//兑换商品
func (qt *QusetTwo) getgoodbyid(id int64) string {
	goodinfo := Goods{}
	for _, g := range qt.GoodsList {
		if g.Id == id {
			goodinfo = g
			break
		}
	}
	if qt.Users[0].Jf < goodinfo.NeedJf {
		return "积分不足"
	}
	qt.Users[0].Jf -= goodinfo.NeedJf
	qt.Users[0].Goods = append(qt.Users[0].Goods, goodinfo.Id)
	return ""
}

//签到
func (qt *QusetTwo) qd() {
	qt.Users[0].Jf += 10
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//input := bufio.NewScanner(os.Stdin)
	//input.Scan()
	//fmt.Println("你输入的是：", input.Text())
	f := true
	for {
		if f {
			f = false
			fmt.Print("请选择题目,a：题目3，b:题目4:\n")
			str1 := getInput()
			if str1 == "a" {
				fmt.Println("题目3开始")
				q1 := SetInit()
				q1.gameStart()
				f = true
			} else if str1 == "b" {
				fmt.Println("题目4")
				q2 := SetQtInit()
				selectinfo(q2)
			} else {
				fmt.Println("请选择正确")
			}
		}

	}
}
func getInput() string {
	//使用os.Stdin开启输入流
	//函数原型 func NewReader(rd io.Reader) *Reader
	//NewReader创建一个具有默认大小缓冲、从r读取的*Reader 结构见官方文档
	in := bufio.NewReader(os.Stdin)
	//in.ReadLine函数具有三个返回值 []byte bool error
	//分别为读取到的信息 是否数据太长导致缓冲区溢出 是否读取失败
	str, _, err := in.ReadLine()
	if err != nil {
		return err.Error()
	}
	return string(str)
}

func selectinfo(q *QusetTwo) {
	fmt.Print("请选择角色,a:管理员,b:普通用户\n")
	str1 := getInput()
	if str1 == "a" {
		q.selectjs(1)
		selectinfot(q)
	} else if str1 == "b" {
		q.selectjs(1)
		userselect(q)
	} else {
		selectinfo(q)
	}
}

//玩家
func userselect(q *QusetTwo) {
	fmt.Print("请选择,a:签到,b:兑换商品,c:返回上一级\n")
	str1 := getInput()
	if str1 == "a" {
		fmt.Print("签到成功获得10积分\n")
		q.qd()
		userselect(q)
	} else if str1 == "b" {
		selectgoods(q)
	} else if str1 == "c" {
		userselect(q)
	}
}
func selectgoods(q *QusetTwo) {
	goodsinfo := q.getgoods()
	for _, g := range goodsinfo {
		fmt.Print("当前积分" + strconv.FormatInt(q.Users[0].Jf, 10) + "\n")
		fmt.Print("商品ID:" + strconv.FormatInt(g.Id, 10) + ",商品名称:" + g.Name + ",商品所需积分:" + strconv.FormatInt(g.NeedJf, 10) + "\n")
	}
	fmt.Print("请输入商品id来兑换或者输入c返回\n")
	goodidstr := getInput()
	if goodidstr == "c" {
		userselect(q)
	} else {
		good, _ := strconv.ParseInt(goodidstr, 10, 64)
		r := q.getgoodbyid(good)
		if r != "" {
			fmt.Print("积分不足")
		} else {
			fmt.Print("兑换成功")
		}
		selectgoods(q)
	}

}
func selectinfot(q *QusetTwo) {
	fmt.Print("请选择,a:查看用户积分,b:设置商品,c:返回上一级\n")
	str1 := getInput()
	if str1 == "a" {
		info, err := q.getusers()
		if err != "" {
			fmt.Print("你没有权限")
			selectinfot(q)
		}
		fmt.Print("用户名" + info[0].Nickname)
		fmt.Print("积分" + strconv.FormatInt(info[0].Jf, 10) + "\n")
		fmt.Print("按b返回\n")
		selectinfot(q)
	} else if str1 == "b" {
		ginfo := Goods{}
		fmt.Print("输入商品名\n")
		str1 := getInput()
		ginfo.Name = str1
		fmt.Print("输入所需积分\n")
		str2 := getInput()
		d, _ := strconv.ParseInt(str2, 10, 64)
		ginfo.NeedJf = d
		q.setgoods(ginfo)
		fmt.Print("新增成功\n")
		selectinfot(q)
	} else if str1 == "c" {
		selectinfo(q)
	}
}

package main

import (
	"fmt"
	"math/rand"
)

//作业1
type QuestOne struct {
	PlayerOne *PlayerInfo
	PlayerTwo *PlayerInfo
	OneWin int64
	TwoWin int64
}

//角色信息
type PlayerInfo struct{
	Hp int64
	Mp int64
	Act int64
	Skill int64
}

//初始化
func SetInit() *QuestOne{
	this:=new(QuestOne)
	this.PlayerOne=&PlayerInfo{
		Hp:100,
		Mp:0,
		Act:10,
		Skill:1,
	}
	this.PlayerTwo=&PlayerInfo{
		Hp:300,
		Mp:0,
		Act:20,
		Skill:2,
	}
	this.OneWin=0
	this.TwoWin=0
	return this
}
//游戏开始
func (qo *QuestOne)gameStart() {
	for i:=0;i<10;i++{
		fmt.Println("第")
		fmt.Println(i+1)
		fmt.Println("回合")
		if i<5{
			fmt.Println("A 先手攻击")
			for {
				qo.aact()
				r:=qo.iswin()
				if r==1{
					fmt.Println("A 获胜 ")
					qo.OneWin++
					break
				}else if r==2{
					fmt.Println("B 获胜 ")
					qo.TwoWin++
					break
				}
				qo.bact()
				if r==1{
					fmt.Println("A 获胜 ")
					qo.OneWin++
					break
				}else if r==2{
					fmt.Println("B 获胜 ")
					qo.TwoWin++
					break
				}

			}
		}else{
			fmt.Println("B 先手攻击 ")
			for {
				qo.bact()
				r:=qo.iswin()
				if r==1{
					fmt.Println("A 获胜 ")
					qo.OneWin++
					break
				}else if r==2{
					fmt.Println("B 获胜 ")
					qo.TwoWin++
					break
				}
				qo.aact()
				if r==1{
					fmt.Println("A 获胜 ")
					qo.OneWin++
					break
				}else if r==2{
					fmt.Println("B 获胜 ")
					qo.TwoWin++
					break
				}

			}
		}
		qo.reset()
	}
	fmt.Println("A 获胜次数")
	fmt.Println(qo.OneWin)
	fmt.Println("B 获胜次数")
	fmt.Println(qo.TwoWin)
}
//重置
func (qo *QuestOne)reset(){
	qo.PlayerOne=&PlayerInfo{
		Hp:100,
		Mp:0,
		Act:10,
		Skill:1,
	}
	qo.PlayerTwo=&PlayerInfo{
		Hp:300,
		Mp:0,
		Act:20,
		Skill:2,
	}
}
//A攻击
 func (qo *QuestOne)aact() error{
	 qo.PlayerTwo.Hp = qo.PlayerTwo.Hp-qo.PlayerOne.Act
	 fmt.Println("A 攻击 伤害：")
	 fmt.Println(qo.PlayerOne.Act)
	 fmt.Println("B 生命")
	 fmt.Println(qo.PlayerTwo.Hp)
	 r:=rand.Intn(2)
	 if r<1{
		 //连击
		 fmt.Println("A 发动技能再次攻击 伤害：")
		 fmt.Println(qo.PlayerOne.Act)
		 qo.PlayerTwo.Hp = qo.PlayerTwo.Hp-qo.PlayerOne.Act
		 fmt.Println("B 生命")
		 fmt.Println(qo.PlayerTwo.Hp)
	 }
	 return nil
 }
 //B攻击
func (qo *QuestOne)bact() error{
	qo.PlayerOne.Hp = qo.PlayerOne.Hp-qo.PlayerTwo.Act
	qo.PlayerTwo.Mp+=10
	fmt.Println("B 攻击 伤害")
	fmt.Println(qo.PlayerTwo.Act)
	fmt.Println("获得蓝量10 ")
	fmt.Println("A 生命")
	fmt.Println(qo.PlayerOne.Hp)
	fmt.Println("\n")
	if qo.PlayerTwo.Mp>=50{
		//发动技能
		qo.PlayerTwo.Mp=0
		qo.PlayerOne.Act = qo.PlayerOne.Act*90/100
		fmt.Println("B 发动技能")
	}
	return nil
}
//判断输赢
func (qo *QuestOne)iswin()int64{
	//
	if qo.PlayerOne.Hp==0{
		return 2
	}
	if qo.PlayerTwo.Hp==0{
		return 1
	}
	return 0
}
/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/07/19 11:37
 */

package main

import "fmt"

// 抽象类,制作饮料,包裹一个模板的全部步骤的实现
type Beverage interface {
	//煮开水
	BoilWater()

	//冲泡
	Brew()

	//导入杯中
	PourInCup()

	//添加蘸料
	AddThings()

	WantAddThings() bool //是否加入酌料Hook
}

// 封装一套流程模板,让具体的制作流程继承且实现
type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	if t.b.WantAddThings() {
		t.b.AddThings()

	}
}

// 具体的子类的模板 制作咖啡
type MakeCoffee struct {
	template //继承模板
}

func (mc *MakeCoffee) BoilWater() {
	fmt.Println("将水煮到100摄氏度")
}
func (mc *MakeCoffee) Brew() {
	fmt.Println("用水冲咖啡豆")
}
func (mc *MakeCoffee) PourInCup() {
	fmt.Println("将冲好的咖啡倒入瓷杯中")
}
func (mc *MakeCoffee) AddThings() {
	fmt.Println("加入糖和牛奶")
}

func (mc *MakeCoffee) WantAddThings() bool {
	return true
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)
	//b是Beverage,是MakeCoffee的接口,这里需要给接口赋值,让b
	makeCoffee.b = makeCoffee
	return makeCoffee
}

type MakeTea struct {
	template
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}

// 煮开水
func (mt *MakeTea) BoilWater() {
	fmt.Println("将水煮沸到80摄氏度")
}

// 冲泡
func (mt *MakeTea) Brew() {
	fmt.Println("用水冲泡茶叶")
}

// 导入杯中
func (mt *MakeTea) PourInCup() {
	fmt.Println("将充好的茶水倒入茶壶中")
}

// 添加蘸料
func (mt *MakeTea) AddThings() {
	fmt.Println("添加柠檬")
}

func (mt *MakeTea) WantAddThings() bool {
	return true
}

func main() {
	//制作一杯咖啡
	makeCoffee := NewMakeCoffee()
	makeCoffee.MakeBeverage()
	fmt.Println("------------------")
	makeTea := NewMakeTea()
	makeTea.MakeBeverage()
}

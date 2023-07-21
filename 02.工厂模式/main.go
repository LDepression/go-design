/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/07/16 16:07
 */

package main

import "fmt"

// ******抽象层********
type Fruit interface {
	Show()
}

type Factory interface {
	CreateFruit() Fruit
}

type apple struct{}

func (a *apple) Show() {
	fmt.Println("我是苹果")
}

type banana struct{}

func (b *banana) Show() {
	fmt.Println("我是香蕉")
}

type appleFac struct {
}

func (f appleFac) CreateFruit() Fruit {
	var a Fruit
	a = new(apple)
	return a
}

type bananaFac struct{}

func (f bananaFac) CreateFruit() Fruit {
	var b Fruit
	b = new(banana)
	return b
}

func main() {
	var fac Factory
	fac = new(appleFac)
	f := fac.CreateFruit()
	f.Show()
}

/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/07/16 16:56
 */

package main

import "fmt"

type Cat struct{}

func (c Cat) Eat() {
	fmt.Println("小猫吃饭")
}

// 给小猫B添加一个可以睡觉的方法
type CatB struct {
	Cat
}

// 给小猫添加一个 可以睡觉的方法 （使用组合的方式）
type CatC struct {
	C *Cat
}

func (cc *CatC) Sleep() {
	fmt.Println("小猫睡觉")
}
func (cb *CatB) Sleep() {
	fmt.Println("小猫睡觉")
}
func main() {
	//通过继承增加的功能，可以正常使用
	cb := new(CatB)
	cb.Eat()
	cb.Sleep()

	//通过组合增加的功能，可以正常使用
	cc := new(CatC)
	cc.C = new(Cat)
	cc.C.Eat()
	cc.Sleep()
}

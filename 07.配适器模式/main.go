/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/07/18 14:41
 */

package main

import "fmt"

// V5 适配的目标
type V5 interface {
	Use5V()
}

// 业务类
type phone struct {
	v V5
}

func NewPhone(v V5) *phone {
	return &phone{v: v}
}

func (p *phone) Charge() {
	fmt.Println("phone 进行充电....")
	p.v.Use5V()
}

// 被适配的对象,适配者
type V220 struct{}

func (V220) Use220V() {
	fmt.Println("使用220V的电压")
}

// 电源适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电")

	//调用适配者的方法
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

func main() {
	iphone := NewPhone(NewAdapter(new(V220)))
	iphone.Charge()
}

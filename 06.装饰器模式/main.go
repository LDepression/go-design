/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/07/17 10:08
 */

package main

import "fmt"

// ******抽象类*******
type Phone interface {
	Show()
}

type Decorator struct {
	phone Phone // 原本应该是一个接口的,但是我们还应该有成员
}

func (Decorator) Show() {
}

// ********实现类********
type Huawei struct{}

func (Huawei) Show() {
	fmt.Println("huawei yyds")
}

type xiaomi struct{}

func (xiaomi) Show() {
	fmt.Println("xiaomi yyds")
}

type MoDecorator struct {
	Decorator //继承Decorator
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

func (m MoDecorator) Show() {
	m.phone.Show()
	fmt.Println("进行了手机贴膜的相关操作")
}

func main() {
	var phone Phone
	phone = Huawei{}
	de := NewMoDecorator(phone)
	de.Show()

}

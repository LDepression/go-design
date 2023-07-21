/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/07/21 9:46
 */

package main

import "fmt"

// ----------抽象层---------
type car interface {
	Drive()
}

type driver interface {
	DriveCar(c car)
}

// ------实现层---------
type benci struct{}

func (b *benci) Drive() {
	fmt.Println("奔驰汽车在行驶")
}

type baoMa struct{}

func (b *baoMa) Drive() {
	fmt.Println("宝马汽车在行驶")
}

type zhang3 struct{}

func (d *zhang3) DriveCar(c car) {
	fmt.Println("张3在驾驶")
	c.Drive()
}

type li4 struct{}

func (d *li4) DriveCar(c car) {
	fmt.Println("李4在驾驶")
	c.Drive()
}

func main() {
	var d driver
	d = &zhang3{}
	b := new(baoMa)
	d.DriveCar(b)
}

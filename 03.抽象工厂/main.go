/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/07/16 15:26
 */

package main

import "fmt"

/*
	中国,日本,美国各自有苹果,香蕉,梨
	梨,香蕉,苹果各自有展示的方法
*/

// *****抽象的产品*****
type AbstractApple interface {
	showApple()
}
type AbstractPear interface {
	showPear()
}

type AbstractBanana interface {
	showBanana()
}

// *****抽象工厂******
type abstractFactory interface {
	createApple() AbstractApple
	createPear() AbstractPear
	createBanana() AbstractBanana
}

// *****实现层********
type ChinaApple struct {
}

func (a ChinaApple) showApple() {
	fmt.Println("中国苹果")
}

type ChinaBanana struct{}

func (b ChinaBanana) showBanana() {
	fmt.Println("中国香蕉")
}

type ChinaPear struct{}

func (b ChinaPear) showPear() {
	fmt.Println("中国梨")
}

type ChinaFac struct{}

func (f *ChinaFac) createApple() AbstractApple {
	var a AbstractApple
	a = new(ChinaApple)
	return a
}
func (f *ChinaFac) createPear() AbstractPear {
	var a AbstractPear
	a = new(ChinaPear)
	return a
}
func (f *ChinaFac) createBanana() AbstractBanana {
	var a AbstractBanana
	a = new(ChinaBanana)
	return a
}

//*****日本产品族*****

type japanApple struct{}

func (a japanApple) showApple() {
	fmt.Println("日本梨")
}

type japanBanana struct{}

func (b japanBanana) showBanana() {
	fmt.Println("日本香蕉")
}

type japanPear struct{}

func (j japanPear) showPear() {
	fmt.Println("日本梨")
}

type JapanFac struct{}

func (j *JapanFac) createApple() AbstractApple {
	var a AbstractApple
	a = new(japanApple)
	return a
}
func (j *JapanFac) createBanana() AbstractBanana {
	var a AbstractBanana
	a = new(japanBanana)
	return a
}
func (j *JapanFac) createPear() AbstractPear {
	var a AbstractPear
	a = new(japanPear)
	return a
}

func main() {
	//****实现中国苹果
	var fc abstractFactory
	fc = new(ChinaFac)
	apple := fc.createApple()
	apple.showApple()

	//***实现日本香蕉

	var jf abstractFactory
	jf = new(JapanFac)
	b := jf.createBanana()
	b.showBanana()

}

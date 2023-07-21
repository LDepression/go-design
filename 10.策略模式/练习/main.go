/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/07/21 8:58
 */

package main

import "fmt"

/*
	商场促销有策略A(8折) 策略B(消费满200,返现100),用策略模式来模拟场景
*/

type sellStrategy interface {
	//实现一个原价与实际价格的方法
	OriginPrice2Price() float64
}
type strategy1 struct {
	price float64
}

func (s *strategy1) OriginPrice2Price() float64 {
	return s.price * 0.8
}

type strategy2 struct {
	price float64
}

func (s *strategy2) OriginPrice2Price() float64 {
	if s.price > 200 {
		return s.price - 100
	}
	return s.price
}

type customer struct {
	s sellStrategy
}

func (c *customer) CustomerWithNewStrategy(strategy sellStrategy) {
	c.s = strategy
}
func (c *customer) Buy() float64 {
	return c.s.OriginPrice2Price()
}

func main() {
	cus := customer{}
	s := new(strategy2)
	s.price = 150
	cus.CustomerWithNewStrategy(s)
	price := cus.Buy()
	fmt.Println(price)
}

/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/07/17 8:41
 */

package main

import "fmt"

type Goods struct {
	Kind string
	Fact bool
}

// ******抽象层*******
type Shopping interface {
	Buy(goods *Goods)
}

type ChinaShopping struct {
}

func (s *ChinaShopping) Buy(good *Goods) {
	fmt.Println("去韩国购物,购买了:", good.Kind)
}

type AmericaShopping struct{}

func (s *AmericaShopping) Buy(goods *Goods) {
	fmt.Println("去美国购物,购买了:", goods.Kind)
}

// ****代理层******

type ShopProxy struct {
	shopping Shopping //为了让原代理调用,这里也就是代理着某个主题
}

func NewProxy(shopping Shopping) Shopping {
	return &ShopProxy{shopping}
}
func (sp *ShopProxy) Distinguish(good *Goods) bool {
	fmt.Println("对[", good.Kind, "]进行了辨别真伪")
	if good.Fact == false {
		fmt.Println("发现假货", good.Kind, "不应该购买")
	}
	return good.Fact
}
func (sp *ShopProxy) Check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "] 进行了海关检查， 成功的带回祖国")
}

func (sp *ShopProxy) Buy(goods *Goods) {
	if sp.Distinguish(goods) {
		sp.shopping.Buy(goods) //调用原代理的具体主题任务
		sp.Check(goods)
	}
}

func main() {
	good1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}
	good2 := Goods{
		Kind: "CTE4",
		Fact: false,
	}
	var shopping Shopping
	shopping = new(ChinaShopping)
	proxy := NewProxy(shopping)
	proxy.Buy(&good1)
	proxy.Buy(&good2)
}

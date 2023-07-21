/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/07/16 16:30
 */

package main

import "fmt"

/*
对于一个银行转账的业务员,如果我们是采用平铺式设计的话,
那么特别容易造成的代码的臃肿.

此时我们可以将职责分出去
比如:转账业务员,存款业务员,支付业务员...
*/
type AbstractBanker interface {
	DoBusi() //抽象逻辑层的处理
}
type saveBanker struct {
}

func (b saveBanker) DoBusi() {
	fmt.Println("进行了存款")
}

type transferBanker struct {
}

func (t transferBanker) DoBusi() {
	fmt.Println("进行了存款")
}

type payBanker struct {
}

func (t payBanker) DoBusi() {
	fmt.Println("进行了支付")
}

func main() {
	//进行存款
	s := &saveBanker{}
	s.DoBusi()

	//进行转账
	tb := &transferBanker{}
	tb.DoBusi()
}

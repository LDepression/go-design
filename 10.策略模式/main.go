/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/07/21 8:48
 */

package main

import "fmt"

type WeaponStrategy interface {
	UseWeapon() //使用武器
}

type Ak47 struct {
}

func (a Ak47) UseWeapon() {
	fmt.Println("使用Ak47 去战斗")
}

type Knife struct{}

func (k Knife) UseWeapon() {
	fmt.Println("使用knife 去战斗")
}

type Hero struct {
	strategy WeaponStrategy //拥有一个抽象的策略
}

func NewStrategy(strategy WeaponStrategy) *Hero {
	hero := &Hero{}
	hero.strategy = strategy
	return hero
}
func (h *Hero) Fight() {
	h.strategy.UseWeapon()
}

func main() {
	var s WeaponStrategy
	s = Ak47{}
	hero := NewStrategy(s)
	hero.Fight()
}

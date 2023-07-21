/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/07/21 9:21
 */

package main

import (
	"fmt"
	"time"
)

// ---------抽象层--------------
type Listener interface {
	OnTeacherComming() //观察者得到通知后要触发的具体动作
	DoBadThing()       //观察者得到通知后要触发的具体动作
}

type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	NotifyTeacher()
	NotifyNoTeacher()
}

//---------实现层---------

type StuZhang3 struct {
	BadThing string
}

func (z StuZhang3) OnTeacherComming() {
	fmt.Println("张3 停止 ", z.BadThing)
}
func (z StuZhang3) DoBadThing() {
	fmt.Println("张3 可以 ", z.BadThing)
}

type StuLi4 struct {
	BadThing string
}

func (l StuLi4) OnTeacherComming() {
	fmt.Println("李4 停止 ", l.BadThing)
}
func (z StuLi4) DoBadThing() {
	fmt.Println("李4 可以 ", z.BadThing)
}

type StuWang5 struct {
	BadThing string
}

func (w StuWang5) OnTeacherComming() {
	fmt.Println("王5 停止 ", w.BadThing)
}
func (z StuWang5) DoBadThing() {
	fmt.Println("王5 可以 ", z.BadThing)
}

// 通知者
type ClassMonitor struct {
	listenerList []Listener
}

func (m *ClassMonitor) AddListener(listener Listener) {
	m.listenerList = append(m.listenerList, listener)
}
func (m *ClassMonitor) RemoveListener(listener Listener) {
	for index, l := range m.listenerList {
		if l == listener {
			m.listenerList = append(m.listenerList[:index], m.listenerList[index+1:]...)
			break
		}
	}
}
func (m *ClassMonitor) NotifyTeacher() {
	for _, lis := range m.listenerList {
		lis.OnTeacherComming()
	}
}

func (m *ClassMonitor) NotifyNoTeacher() {
	for _, lis := range m.listenerList {
		lis.DoBadThing()
	}
}

/*

	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
*/

func main() {
	s1 := &StuZhang3{BadThing: "抄作业"}
	s2 := &StuLi4{BadThing: "玩王者荣耀"}
	s3 := &StuWang5{BadThing: "看赵4玩王者荣耀"}

	ClassMonitor := new(ClassMonitor)
	ClassMonitor.AddListener(s1)
	ClassMonitor.AddListener(s2)
	ClassMonitor.AddListener(s3)

	fmt.Println("上课了,老师没有来")
	ClassMonitor.NotifyNoTeacher()

	time.Sleep(5 * time.Second)
	fmt.Println("----------------现在老师来了--------------")
	ClassMonitor.NotifyTeacher()
}

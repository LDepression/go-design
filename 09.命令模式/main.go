/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/07/20 9:52
 */

package main

import "fmt"

type Doctor struct{}

func (d Doctor) treatEye() {
	fmt.Println("治疗眼睛")
}
func (d Doctor) treatNose() {
	fmt.Println("治疗鼻子")
}

// 抽象的命令
type Command interface {
	Treat()
}

// 治疗眼睛的病单
type CommandTreatEye struct {
	doctor *Doctor
}

func (cmd *CommandTreatEye) Treat() {
	cmd.doctor.treatEye()
}

// 治疗鼻子的病单
type CommandTreatNose struct {
	doctor *Doctor
}

func (cmd *CommandTreatNose) Treat() {
	cmd.doctor.treatNose()
}

type Nurse struct {
	CmdList []Command
}

func (n Nurse) Notify() {
	if n.CmdList == nil {
		return
	}
	for _, cmd := range n.CmdList {
		cmd.Treat()
	}
}

func main() {
	doctor := new(Doctor)
	cmdEye := CommandTreatEye{doctor: doctor}
	cmdNose := CommandTreatNose{doctor: doctor}

	nurse := new(Nurse)
	nurse.CmdList = append(nurse.CmdList, &cmdEye)
	nurse.CmdList = append(nurse.CmdList, &cmdNose)

	nurse.Notify()
}

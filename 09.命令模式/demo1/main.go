/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/07/20 9:40
 */

package main

import "fmt"

type Doctor struct{}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}
func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// 治疗眼睛的病单
type CommandTreatEye struct {
	doctor *Doctor
}

func (e *CommandTreatEye) Treat() {
	e.doctor.treatEye()
}

// 治疗眼睛的病单
type CommandTreatNose struct {
	doctor *Doctor
}

func (e *CommandTreatNose) Treat() {
	e.doctor.treatNose()
}

func main() {

}

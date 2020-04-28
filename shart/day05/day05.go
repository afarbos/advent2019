package day05

import(
	"github.com/sean-hart/advent2019/shart/libs"
)

// RunProgram will 
func RunProgram(inputInt int, memory []int, nextPointer int) (int, []int, int){
	output := 0
	output, memory, nextPointer = libs.RunComputer(inputInt, memory, nextPointer)
	return output, memory, nextPointer
}

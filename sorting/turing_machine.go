package main

import (
	"fmt"
)

var (
	a     [30000]byte
	prog  = "++++++++++[>++++++++++<-]>++++.+."
	p, pc int
)

func loop(inc int) {
	for i := inc; i != 0; pc += inc {
		print("_",pc+inc)
		switch prog[pc+inc] {
		case '[':
			print("[")
			i++
		case ']':
			print("]")
			i--
		}
	}
}

func main() {
	for {
		print("pc:",pc)
		switch prog[pc] {
		case '>':
			print(">")
			p++
			print("p:",p)
		case '<':
			print("<")
			p--
			print("p:",p)
		case '+':
			print("+")
			a[p]++
			print("a[",p,"]:",a[p])
		case '-':
			print("-")
			a[p]--
			print("a[",p,"]:",a[p])
		case '.':
			print(".")
			fmt.Println(string(a[p]))
		case '[':
			print("[")
			if a[p] == 0 {
				print("_a[",p,"]==0,jumptocorresponding]_")
				loop(1)
			} else {
				print("_a[",p,"]!=0,next_")
			}
		case ']':
			print("]")
			if a[p] != 0 {
				print("_a[",p,"]!=0,gotocorresponding[_")
				loop(-1)
			} else {
				print("_a[",p,"]==0,next_")
			}
		default:
			fmt.Println("Illegal instruction")
		}
		pc++
		if pc == len(prog) {
			return
		}
		println("\n")
	}
}

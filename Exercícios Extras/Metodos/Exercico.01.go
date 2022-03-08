//Realize a implementação de uma pilha e seus dois principais métodos, push() e pop().

package main

import "fmt"

type Stack struct {
	Tail *Node
}

func (s *Stack) Push(value string) {
	node := Node{Value: value}
	if s.Tail != nil {
		node.Next = s.Tail
	}

	s.Tail = &node
}
func (s *Stack) Pop() string {
	node := s.Tail
	s.Tail = node.Next

	return node.Value
}

type Node struct {
	Value string
	Next  *Node
}

func main() {
	stack := Stack{}

	stack.Push("Erika")
	stack.Push("João")
	stack.Push("Camila")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}

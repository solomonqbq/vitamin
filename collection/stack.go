package common

import "container/list"

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (this *Stack) Push(value interface{}) {
	this.list.PushBack(value)
}

func (this *Stack) Pop() interface{} {
	e := this.list.Back()
	if e != nil {
		this.list.Remove(e)
		return e.Value
	} else {
		return nil
	}
}

func (this *Stack) Peak() interface{} {
	e := this.list.Back()
	if e != nil {
		return e.Value
	} else {
		return nil
	}
}

func (this *Stack) Len() int {
	return this.list.Len()
}

func (this *Stack) Empty() bool {
	return this.list.Len() == 0
}

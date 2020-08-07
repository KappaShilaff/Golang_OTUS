package main

import "fmt"

type List struct {
	data int
	next *List
	prev *List
}

func (l *List)First()*List{
	for l.prev != nil {
		l = l.prev
	}
	return l
}

func (l *List)Last()*List{
	for l.next != nil {
		l = l.next
	}
	return l
}

func (l *List) PushFrontList(v *List){
	if l.next == nil {
		l.next = v
	} else {
		temp := l.next
		l.next = v
		temp.prev = v
		v.next = temp
	}
	v.prev = l
}

func (l *List) PushBackList(v *List){
	if l.prev == nil {
		l.prev = v
	} else {
		temp := l.prev
		l.prev = v
		temp.next = v
		v.prev = temp
	}
	v.next = l
}

func (l *List)Len()int {
	i := 0
	l = l.First()
	for l != nil {
		l = l.next
		i++
	}
	return i
}

func (l *List)Remove()*List {
	if l.next != nil && l.prev != nil {
		l.next.prev = l.prev
		l.prev.next = l.next
		return l.prev
	} else if l.prev != nil {
		l.prev.next = nil
		return l.prev
	} else if l.next != nil {
		l.next.prev = nil
		return l.next
	}
	return nil
}

func (l *List)Nex()*List {
	if l.next != nil {
		return l.next
	}
	return l
}

func (l *List)Prev() *List{
	if l.prev != nil {
		return l.prev
	}
	return l
}

func (l *List)PrintAll(){
	l = l.First()
	for l != nil {
		fmt.Printf("%d\n", l.data)
		l = l.next
	}
}

func (l *List)PrintAllReverse(){
	l = l.Last()
	for l != nil {
		fmt.Printf("%d\n", l.data)
		l = l.prev
	}
}

func (l *List) GetI(i int)(*List, error){
	l = l.First()
	n := 0
	for n = 0; n < i && l.next != nil; n++ {
		l = l.next
	}
	if  n != i {
		return l, fmt.Errorf("[GetI.Arguments:%d]", i)
	}
	return l, nil
}

func (l *List) RemoveI(i int)(*List, error){
	l, err := l.GetI(i)
	if err != nil {
		return l, fmt.Errorf("[RemoveI.Arguments:%d]%v", i, err)
	}
	return l.Remove(), nil
}

func (l *List) PushFront(item int){
	l.PushFrontList(&List{item, nil, nil})
}

func (l *List) PushBack(item int){
	l.PushBackList(&List{item, nil, nil})
}

func main() {
	kek := &List{228, nil, nil}
	for i := 0; i < 7; i++ {
		kek.PushFront(i)
		fmt.Print("insert front\n")
		kek.PrintAll()
	}
	for i := 6; i >= 0; i-- {
		kek.PushBack(i)
		fmt.Print("insert back\n")
		kek.PrintAll()
	}
	kek = kek.Last()
	fmt.Printf("Last: %d\n", kek.data)
	kek = kek.First()
	fmt.Printf("First: %d\n", kek.data)
	fmt.Printf("Len: %d\n", kek.Len())
	kek.PrintAll()
	kek, err := kek.RemoveI(-1)
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR%v", err))
	}
	fmt.Printf("Last: %d\n", kek.data)
	kek.PrintAll()
	kek.PrintAllReverse()
}

package main

import (
	"fmt"
)

type EventTime struct {
	Start int
	End   int
	Name  string
}

type EventDate struct {
	Age   int
	Month int
	Day   int
}

type CrStruct struct {
	Date map[EventDate][]EventTime
}

func IsValid(time EventTime, baseOfTime []EventTime) bool {
	for _, chunk := range baseOfTime {
		if (chunk.Start >= time.Start && time.Start <= chunk.End) || (chunk.Start >= time.End && time.End <= chunk.End) {
			return false
		}
	}
	return true
}

func (t EventDate) ReturnData() EventDate {
	return t
}

func (t EventTime) ReturnTime() EventTime {
	return t
}

type InsertData interface {
	ReturnData() EventDate
}

type InsertTime interface {
	ReturnTime() EventTime
}

func (t *CrStruct) EventInsert(d InsertData, ti InsertTime) error {
	date := d.ReturnData()
	time := ti.ReturnTime()
	if t.Date == nil {
		t.Date = map[EventDate][]EventTime{}
	}
	baseOfTime, ok := t.Date[date]
	if ok != true {
		t.Date[date] = []EventTime{time}
	} else {
		if IsValid(time, baseOfTime) == true {
			t.Date[date] = append(t.Date[date], time)
		} else {
			return fmt.Errorf("[ERROR] Time is reserved")
		}
	}
	return nil
}

func (t *CrStruct) EventAllNameRemove(name string) {
	for i, day := range t.Date {
		for k, times := range day {
			if times.Name == name {
				if k < len(day)-1 {
					t.Date[i] = append(t.Date[i][:k], t.Date[i][k+1:]...)
				} else if len(day) > 1 {
					t.Date[i] = t.Date[i][:k-1]
				}
				if len(day) == 1 {
					delete(t.Date, i)
				} else if len(t.Date[i]) == 0 {
					delete(t.Date, i)
				}
			}
		}
	}
}

func (t *CrStruct) PrintAllEvents() error{
	if t.Date == nil {
		return fmt.Errorf("[ERROR] nil map")
	}
	if len(t.Date) == 0 {
		fmt.Printf("Zero events!")
	}
	for day, date := range t.Date {
		s := fmt.Sprintf("%d.%02d.%02d	|", day.Age, day.Month, day.Day)
		for _, time := range date {
			s += fmt.Sprintf(" %02d:%02d-%02d:%02d %s|", time.Start / 100, time.Start % 100, time.End / 100, time.End % 100, time.Name)
		}
		fmt.Printf("%s\n", s)
	}
	return nil
}

//func (t *CrStruct) PrintAllEventsPrint() error{
//	if t.Date == nil {
//		return fmt.Errorf("[ERROR] nil map")
//	}
//	if len(t.Date) == 0 {
//		fmt.Printf("Zero events!")
//	}
//	for day, date := range t.Date {
//		fmt.Printf("%d.%02d.%02d	|", day.Age, day.Month, day.Day)
//		for _, time := range date {
//			fmt.Printf(" %02d:%02d-%02d:%02d %s|", time.Start / 100, time.Start % 100, time.End / 100, time.End % 100, time.Name)
//		}
//		fmt.Printf("\n")
//	}
//	return nil
//}

func (t *CrStruct) EventNameDateRemove(name string, d InsertData) {
	date := d.ReturnData()
	for i, time := range t.Date[date] {
		if time.Name == name {
			if i < len(t.Date[date])-1 {
				t.Date[date] = append(t.Date[date][:i], t.Date[date][i+1:]...)
			} else if len(t.Date[date]) > 1 {
				t.Date[date] = t.Date[date][:i-1]
			}
			if len(t.Date[date]) == 1 {
				delete(t.Date, date)
			} else if len(t.Date[date]) == 0 {
				delete(t.Date, date)
			}
		}
	}
}

func main() {
	kek := CrStruct{}
	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1930, 2000, "kekw"})
	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1900, 1940, "mda v prolete"})
	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{2010, 2040, "kekw"})
	_ = kek.EventInsert(EventDate{2021, 7, 30}, EventTime{2010, 2040, "kekw"})
	_ = kek.EventInsert(EventDate{2020, 7, 31}, EventTime{1900, 1940, "filmec"})
	_ = kek.PrintAllEvents()
	kek.EventNameDateRemove("filmec", EventDate{2020, 7, 31})
	kek.EventAllNameRemove("kekw")
	_ = kek.PrintAllEvents()
	//TryReAlloc()
}

//func memI() {
//	a := make([]EventTime, 1000)
//	for i := 0; i < len(a); i++ {
//		a[i].Name = "1"
//		a[i].Start = 2000
//		a[i].End = a[i].Start
//	}
//}
//
//func memRange() {
//	a := make([]EventTime, 1000)
//	for _, mem := range a {
//		mem.Name = "1"
//		mem.Start = 2000
//		mem.End = mem.Start
//	}
//}

//func TryReAlloc() {
//	a := make([]string, 2)
//	a[0] = "lol"
//	a[1] = "kek"
//	fmt.Printf("TYPE []string\n%v\n", a)
//	for i, str := range a {
//		str = "mdakers" // тутачки реалок
//		fmt.Printf("NEW %s OLD %s\n", str, a[i])
//	}
//	fmt.Printf("%v", a)
//	b := make([][]int, 2)
//	b[0] = append(b[0], 1, 2)
//	b[1] = append(b[1], 3, 4)
//	fmt.Printf("\nTYPE [][]int\n%v\n", b)
//	fmt.Printf("cap %v p %p\n", cap(b[0]), b[0])
//	for _, sl := range b {
//		fmt.Printf("SL cap %v p %p\n", cap(sl), sl)
//		sl[0] = 5 // тут нет реалока поэтому старое меняется тоже
//		sl[1] = 6
//	}
//	fmt.Printf("cap %v p %p\n", cap(b[0]), b[0])
//	fmt.Printf("%v\n", b)
//	for i, sl := range b {
//		sl = append(sl, 1) // тут тоже реалок
//		fmt.Printf("NEW %v\nOLD %v\n", sl, b[i])
//	}
//	b[0] = append(b[0], 0)
//	b[1] = append(b[1], 0)
//	fmt.Printf("cap %v p %p\n", cap(b[0]), b[0])
//	for i, sl := range b {
//		fmt.Printf("SL BEFORE cap %v len %v p %p\n", cap(sl), len(sl), sl)
//		sl = append(sl, 1)
//		fmt.Printf("SL AFTER cap %v len %v p %p\n", cap(sl), len(sl), sl)
//		fmt.Printf("NEW %v\nOLD %v\n", sl, b[i]) // а вот теперь нет реалока
//	}
//	fmt.Printf("cap %v p %p\n", cap(b[1]), b[1])
//	kek := b[0]
//	xx := reflect.ValueOf(kek)
//	fmt.Printf("REFLECT %v\n", xx.Slice(0,4))
//	fmt.Printf("%v len %d", b, len(b[0]))
//}

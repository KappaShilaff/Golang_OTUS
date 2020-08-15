package main

import (
	"fmt"
	"sort"
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

type DateSortStruct struct {
	sortInt int
	EDate EventDate
}

func ConvertMapToStruct(m *map[EventDate][]EventTime) ([]DateSortStruct, error) {
	if len(*m) == 0 {
		return []DateSortStruct{}, fmt.Errorf("[ERROR ConvertMtoS] Zero map")
	}
	d := make([]DateSortStruct, len(*m))
	i := 0
	for date := range *m {
		d[i].EDate = date
		d[i].sortInt = date.Day + date.Month * 100 + date.Age * 10000
		i++
	}
	sort.Slice(d, func(i, j int) bool {
		return d[i].sortInt < d[j].sortInt
	})
	return d, nil
}

func (t *CrStruct) PrintAllEvents() error {
	if t.Date == nil {
		return fmt.Errorf("[ERROR] nil map")
	}
	if len(t.Date) == 0 {
		fmt.Printf("Zero events!")
	}
	DateSl, _ := ConvertMapToStruct(&t.Date)
	for _, Date := range DateSl {
		day := Date.EDate
		s := fmt.Sprintf("%d.%02d.%02d	|", day.Age, day.Month, day.Day)
		for _, time := range t.Date[day] {
			s += fmt.Sprintf(" %02d:%02d-%02d:%02d %s|", time.Start/100, time.Start%100, time.End/100, time.End%100, time.Name)
		}
		fmt.Printf("%s\n", s)
	}
	return nil
}

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

type InserIventInterface interface {
	EventInsert(d InsertData, ti InsertTime) error
}

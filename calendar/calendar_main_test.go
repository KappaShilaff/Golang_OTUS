package main

import (
	"os"
	"testing"
)

func BenchmarkAlloc (b *testing.B) {
	os.Stdout = os.NewFile(0, os.DevNull)
	kek := CrStruct{}
	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1930, 2000, "kekw"})
	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1900, 1940, "mda v prolete"})
	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1010, 1040, "kekw"})
	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1110, 1240, "kekw"})
	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1310, 1340, "kekw"})
	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1410, 1440, "kekw"})
	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1510, 1540, "kekw"})
	_ = kek.EventInsert(EventDate{2021, 7, 30}, EventTime{1610, 1640, "kekw"})
	_ = kek.EventInsert(EventDate{2020, 7, 31}, EventTime{1900, 1940, "filmec"})
	for n := 0; n < b.N; n++ {
		_ = kek.PrintAllEvents()
	}
}

//func BenchmarkMemI (b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		memI()
//	}
//}
//
//func BenchmarkMemRange (b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		memRange()
//	}
//}

//func BenchmarkPrint(b *testing.B) {
//	os.Stdout = os.NewFile(0, os.DevNull)
//	kek := CrStruct{}
//	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1930, 2000, "kekw"})
//	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1900, 1940, "mda v prolete"})
//	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1010, 1040, "kekw"})
//	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1110, 1240, "kekw"})
//	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1310, 1340, "kekw"})
//	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1410, 1440, "kekw"})
//	_ = kek.EventInsert(EventDate{2020, 7, 30}, EventTime{1510, 1540, "kekw"})
//	_ = kek.EventInsert(EventDate{2021, 7, 30}, EventTime{1610, 1640, "kekw"})
//	_ = kek.EventInsert(EventDate{2020, 7, 31}, EventTime{1900, 1940, "filmec"})
//	for n := 0; n < b.N; n++ {
//		_ = kek.PrintAllEventsPrint()
//	}
//}

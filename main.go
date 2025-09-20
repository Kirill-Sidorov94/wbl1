package main

import(
	"fmt"
	// "time"
	// "math/rand"
	"os"
	"github.com/Kirill-Sidorov94/wbl1/lib"
)

func main() {
	//rand.Seed(time.Now().UnixNano())

	// l1.1 go run l1-1.go main.go
	// a := Action{Human: Human{Name: "action"}}
	//fmt.Println(a.GetName())
	
	// l1.2 go run l1-2.go main.go
	//CompetitiveSquaring()
	
	// l1.3 go run l1-3.go main.go
	//MultipleWorkersRunning(4)
	
	// l1.4 go run l1-4.go main.go
	//exitSigTerm()
	
	// l1.5 go run l1-5.go main.go
	//channelTimeout()
	
	// l1.6 go run l1-6.go main.go
	//stopGoRoutine()
	
	// l1.7 go run l1-7.go main.go
	//concurrentEntryInMap()
	
	// l1.8 go run l1-8.go main.go
	// num := int64(25)
	// replaceNum, err := replaceBit(num, 8, true)
	// if err != nil {
	// 	fmt.Printf("replaceBit(): %v", err)
	// }
	// fmt.Println(replaceNum)
	
	// l1.9 go run l1-9.go main.go
	// conveyorOfNumbers()
	
	// l1.10 go run l1-10.go main.go
	//temperatureGrouping()
	
	// l1.11 go run l1-11.go main.go
	// intersectionOfSets()
	
	// l1.12 go run l1-12.go main.go
	// setOfStrings()
	
	// l1.13 go run l1-13.go main.go
	// exchangeOfValues()
	
	// l1.14 go run l1-14.go main.go
	// num := 14
	// typeDefenition(num)
	// ch := make(chan any, 1)
	// typeDefenition(ch)
	
	// l1.16 go run l1-16.go main.go
	// unsortArr := make([]int, 10)
    // for i := 0; i < len(unsortArr); i++ {
    //     unsortArr[i] = rand.Intn(101)
    // }
    // sortArr := quickSort(unsortArr)
    // fmt.Println(sortArr)
    
    // l1.17 go run l1-16.go l1-17.go main.go
    // randIndex := rand.Intn(len(sortArr))
    // fmt.Println(binarySearch(sortArr, sortArr[randIndex]))
    // fmt.Println(binarySearch(sortArr, 101))
    
    // l1.18 go run l1-18.go main.go
    // competitiveCounter(20)
    
    // l1.19 go run l1-19.go main.go
    // fmt.Println(stringReversal("Hello world"))
    
    // l1.20 go run l1-20.go main.go
    // fmt.Println(revertWordInSentence("Hello world"))
    
    // l1.21 go run l1-21.go main.go
    // adapterDemonstaration()
    
    // l1.22 go run l1-22.go main.go
    // workingWithBigNums()
    
    // l1.23 go run l1-23.go main.go
    // sl := []int{1, 2, 3, 4, 5, 6}
    // newSl := removeElFromSlice(sl, 3)
    // fmt.Println(newSl)
    
    // l1.24 go run l1-24.go main.go
    // p1 := NewPoint(55.7558, 37.6173)
	// p2 := NewPoint(59.9343, 30.3351)
	// fmt.Println(p1.Distance(p2))
	
	// l1.25 go run l1-25.go main.go
	// customSleep(4 * time.Second)
	
	// l1.26 go run l1-26.go main.go
	// uniqStr := "absdefg"
	// unUniqStr := "abSsdefg"
	// fmt.Println(uniqSymbolsToStr(uniqStr))
	// fmt.Println(uniqSymbolsToStr(unUniqStr))
	
	// l2.8 go run l2-8.go main.go
	if time, err := lib.NtpTime(); err != nil {
		fmt.Fprintln(os.Stderr, "main: %v", err)
		os.Exit(1)
	}
	fmt.Println(time)
}
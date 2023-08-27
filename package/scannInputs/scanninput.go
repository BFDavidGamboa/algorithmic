package scanninput

import "fmt"

const InputDataMessage = "Please in the data type"

type I interface {
	ScanInput()
}

type TInt struct {
	t int
}

type TInt32 struct {
	t int32
}

type TInt64 struct {
	t int64
}

type TString struct {
	t string
}

type TBool struct {
	t bool
}

func (ST *TString) ScanInput() TString {
	fmt.Scanf("%s", ST.t)
	fmt.Println(InputDataMessage, " ", ST.t)
	return *ST
}

func (ST TInt) ScanInput() int {
	fmt.Scanf("%d", ST.t)
	fmt.Println(InputDataMessage, " ", ST.t)
	return ST.t
}

func (ST TInt32) ScanInput() int32 {
	fmt.Scanf("%d", ST.t)
	fmt.Println(InputDataMessage, " ", ST.t)
	return ST.t
}

func (ST TInt64) ScanInput() int64 {
	fmt.Scanf("%d", ST.t)
	fmt.Println(InputDataMessage, " ", ST.t)
	return ST.t
}

func (ST TBool) ScanInput() bool {
	fmt.Scanf("%t", ST.t)
	fmt.Println(InputDataMessage, " ", ST.t)
	return ST.t
}

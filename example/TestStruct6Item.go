/// -------------------------------------------------------------------------------
/// THIS FILE IS ORIGINALLY GENERATED BY redis2go.exe.
/// PLEASE DO NOT MODIFY THIS FILE.
/// -------------------------------------------------------------------------------
package main

import (
	cstruct "github.com/fananchong/cstruct-go"
)

type TestStruct6ItemData struct {
	Myb   bool
	Myf1  float32
	Myf2  float64
	Myi1  int8
	Myi2  int16
	Myi3  int32
	Myi4  int64
	Myi6  uint8
	Myi7  uint16
	Myi8  uint32
	Myi9  uint64
	Mys1  string
	Mys2  []byte
	Myst1 StrcutXX
	Myst2 StrcutYY
}

type TestStruct6Item struct {
	SubKey int32
	__data TestStruct6ItemData
	__root *TestStruct6
}

func NewTestStruct6Item(subKey int32, root *TestStruct6) *TestStruct6Item {
	return &TestStruct6Item{
		SubKey: subKey,
		__root: root,
	}
}

func (this *TestStruct6Item) GetMyb() bool {
	return this.__data.Myb
}

func (this *TestStruct6Item) GetMyf1() float32 {
	return this.__data.Myf1
}

func (this *TestStruct6Item) GetMyf2() float64 {
	return this.__data.Myf2
}

func (this *TestStruct6Item) GetMyi1() int8 {
	return this.__data.Myi1
}

func (this *TestStruct6Item) GetMyi2() int16 {
	return this.__data.Myi2
}

func (this *TestStruct6Item) GetMyi3() int32 {
	return this.__data.Myi3
}

func (this *TestStruct6Item) GetMyi4() int64 {
	return this.__data.Myi4
}

func (this *TestStruct6Item) GetMyi6() uint8 {
	return this.__data.Myi6
}

func (this *TestStruct6Item) GetMyi7() uint16 {
	return this.__data.Myi7
}

func (this *TestStruct6Item) GetMyi8() uint32 {
	return this.__data.Myi8
}

func (this *TestStruct6Item) GetMyi9() uint64 {
	return this.__data.Myi9
}

func (this *TestStruct6Item) GetMys1() string {
	return this.__data.Mys1
}

func (this *TestStruct6Item) GetMys2() []byte {
	return this.__data.Mys2
}

func (this *TestStruct6Item) GetMyst1(mutable bool) *StrcutXX {
	if mutable {
		this.__root.__dirtyData[this.SubKey] = 1
	}
	return &this.__data.Myst1
}

func (this *TestStruct6Item) GetMyst2(mutable bool) *StrcutYY {
	if mutable {
		this.__root.__dirtyData[this.SubKey] = 1
	}
	return &this.__data.Myst2
}

func (this *TestStruct6Item) SetMyb(value bool) {
	this.__data.Myb = value
	this.__root.__dirtyData[this.SubKey] = 1
}

func (this *TestStruct6Item) SetMyf1(value float32) {
	this.__data.Myf1 = value
	this.__root.__dirtyData[this.SubKey] = 1
}

func (this *TestStruct6Item) SetMyf2(value float64) {
	this.__data.Myf2 = value
	this.__root.__dirtyData[this.SubKey] = 1
}

func (this *TestStruct6Item) SetMyi1(value int8) {
	this.__data.Myi1 = value
	this.__root.__dirtyData[this.SubKey] = 1
}

func (this *TestStruct6Item) SetMyi2(value int16) {
	this.__data.Myi2 = value
	this.__root.__dirtyData[this.SubKey] = 1
}

func (this *TestStruct6Item) SetMyi3(value int32) {
	this.__data.Myi3 = value
	this.__root.__dirtyData[this.SubKey] = 1
}

func (this *TestStruct6Item) SetMyi4(value int64) {
	this.__data.Myi4 = value
	this.__root.__dirtyData[this.SubKey] = 1
}

func (this *TestStruct6Item) SetMyi6(value uint8) {
	this.__data.Myi6 = value
	this.__root.__dirtyData[this.SubKey] = 1
}

func (this *TestStruct6Item) SetMyi7(value uint16) {
	this.__data.Myi7 = value
	this.__root.__dirtyData[this.SubKey] = 1
}

func (this *TestStruct6Item) SetMyi8(value uint32) {
	this.__data.Myi8 = value
	this.__root.__dirtyData[this.SubKey] = 1
}

func (this *TestStruct6Item) SetMyi9(value uint64) {
	this.__data.Myi9 = value
	this.__root.__dirtyData[this.SubKey] = 1
}

func (this *TestStruct6Item) SetMys1(value string) {
	this.__data.Mys1 = value
	this.__root.__dirtyData[this.SubKey] = 1
}

func (this *TestStruct6Item) SetMys2(value []byte) {
	this.__data.Mys2 = value
	this.__root.__dirtyData[this.SubKey] = 1
}

func (this *TestStruct6Item) Unmarshal(data []byte) error {
	if err := cstruct.Unmarshal(data, &this.__data); err != nil {
		return err
	}
	return nil
}

func (this *TestStruct6Item) Marshal() ([]byte, error) {
	data, err := cstruct.Marshal(&this.__data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

package main

import "testing"

//テストできるが終了指示できないし、テストしにくい
//ポートも固定されている
func TestMainFunc(t *testing.T) {
	go main()
}

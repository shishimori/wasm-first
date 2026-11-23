//go:build js && wasm

package main

import "syscall/js"

// js.Valueに依存しない型だけで実装（Goのコードだけでテストがしやすくなる）
func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func main() {
	c := make(chan struct{})

	js.Global().Set("add", js.FuncOf(func(this js.Value, args []js.Value) any {
		// 検証のため、すべて Number 型が渡される前提で実装
		// Number 型以外が渡されると panic が発生
		var parsedArgs []int
		for _, arg := range args {
			parsedArgs = append(parsedArgs, arg.Int())
		}
		return add(parsedArgs...)
	}))

	<-c
}

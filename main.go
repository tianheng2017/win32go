package main

import (
	"fmt"
	"syscall"

	"github.com/tianheng2017/win"
)

func main() {
	var pid uint32
	var base uintptr = 0x0161B60
	var address1 uintptr
	var address2 uintptr
	var address3 uintptr
	var address4 uintptr
	var data uintptr
	var ret uintptr

	lpWindowName, err := syscall.UTF16PtrFromString("Tutorial-x86_64")
	if err != nil {
		panic(err.Error())
	}
	// 获取窗口句柄
	hmw := win.FindWindow(nil, lpWindowName)
	// 获取进程ID
	win.GetWindowThreadProcessId(hmw, &pid)
	// 获取进程句柄
	ph := win.OpenProcess(0xFFF, 0, pid)
	// 读取内存，基址+偏移
	win.ReadProcessMemory(ph, base+0x7C0, &address1, 4, &ret)
	win.ReadProcessMemory(ph, address1+0x7e0, &address2, 4, &ret)
	win.ReadProcessMemory(ph, address2+0x88, &address3, 4, &ret)
	win.ReadProcessMemory(ph, address3+0x20, &address4, 4, &ret)
	win.ReadProcessMemory(ph, address4+0x7F8, &data, 4, &ret)
	fmt.Println("当前健康值为: ", data)
	// 修改内存
	var p uintptr = 1000
	win.WriteProcessMemory(ph, address4+0x7F8, &p, 4, &ret)
	// 读取修改后的内存
	win.ReadProcessMemory(ph, address4+0x7F8, &data, 4, &ret)
	fmt.Println("修改后健康值为: ", data)
	fmt.Print("按任意键退出...")
	fmt.Scanln()
}

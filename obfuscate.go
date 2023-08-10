package main

import (
	"fmt"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

func encryptFunc(funcAddr uintptr, SleepTime uint32) error {

	fmt.Printf("[+] Encrypting function at address 0x%x\n", funcAddr)
	FuncLength := findAddrLength(funcAddr)

	var oldprotect uint32
	err := windows.VirtualProtect(funcAddr, uintptr(FuncLength), windows.PAGE_EXECUTE_READWRITE, &oldprotect)
	fmt.Printf("[+] Changing memory permissions to the memory region from 0x%x to 0x%x\n", oldprotect, windows.PAGE_EXECUTE_READWRITE)
	if err != nil {
		return fmt.Errorf("[Error] failed VirtualProtect: %v", err)
	}

	fmt.Println("[+] Encrypting memory region")
	xorFunc(funcAddr, FuncLength, 0xFF)

	fmt.Printf("[+] Sleeping for %d seconds\n", SleepTime)
	time.Sleep(time.Duration(SleepTime) * time.Second)

	fmt.Println("[+] Woke up")
	fmt.Println("[+] Decrypting memory region")
	xorFunc(funcAddr, FuncLength, 0xFF)

	fmt.Printf("[+] Restoring memory permissions of the memory region to 0x%x\n", oldprotect)
	err = windows.VirtualProtect(funcAddr, uintptr(FuncLength), oldprotect, &oldprotect)
	if err != nil {
		return fmt.Errorf("[Error] failed VirtualProtect: %v", err)
	}

	fmt.Println("[+] Success")
	//Success
	return nil
}

func findAddrLength(funcAddr uintptr) uint32 {
	var i uint32 = 0
	for {
		if *(*byte)(unsafe.Pointer(funcAddr + uintptr(i))) == 0xc3 {
			break
		}
		i++
	}
	return i
}

func xorFunc(funcAddr uintptr, funcLength uint32, key byte) { // Could be replaced with something fancy like SystemFunction032
	for i := uint32(0); i < funcLength; i++ {
		*(*byte)(unsafe.Pointer(funcAddr + uintptr(i))) = *(*byte)(unsafe.Pointer(funcAddr + uintptr(i))) ^ key
	}

}

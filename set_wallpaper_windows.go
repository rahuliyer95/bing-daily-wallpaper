package main

import (
	"log"
	"os"
	"syscall"
	"unsafe"
)

// GetWallpaperPath returns the path to store the wallpaper
func GetWallpaperPath() string {
	return os.Getenv("HOME") + "/Pictures/wallpaper.jpg"
}

// SetWallpaper sets the wallpaper from path for macOS
func SetWallpaper(path string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	systemParametersInfoW := user32.NewProc("SystemParametersInfoW")
	ret, _, _ := systemParametersInfoW.Call(
		20,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))),
		0x1)
	log.Printf("Wallpaper set with return code %d\n", ret)
}

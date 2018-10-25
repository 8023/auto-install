package main

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/go-vgo/robotgo"
	"golang.org/x/sys/windows/registry"
)

// DesktopIcon 枚举桌面图标类型
type DesktopIcon int

// WinrarVersion 枚举WinRAR版本
type WinrarVersion int

// OfficeVersion 枚举Office版本
type OfficeVersion int

// AdobepdfVersion 枚举AdobePDF版本
type AdobepdfVersion int

// 枚举类型
const (
	MyComputer DesktopIcon = iota
	ControlPanel
	Network
	Recycle
	Mydocument

	Winrar501x86 WinrarVersion = iota
	Winrar501x64

	Office2007x86 OfficeVersion = iota
	Office2010x86

	AcroRdrDC157 AdobepdfVersion = iota
)

func main() {
	// swpath, _ := filepath.Abs("./software")
	// installWinrar(swpath, Winrar501x64)
	// installAdobepdf(swpath, AcroRdrDC157)
	// showDesktopIcon()
	println(robotgo.Scale())
	println(robotgo.ScaleX())
	println(robotgo.ScaleY())
	bitmap := robotgo.OpenBitmap("test.png")
	// robotgo.BitmapClick(bitmap)
	x, y := robotgo.FindBitmap(bitmap)
	robotgo.MoveMouse(x, y)
	println(x, y)

}

func showDesktopIcon(icons ...DesktopIcon) {
	regname := map[DesktopIcon]string{
		MyComputer:   "{20D04FE0-3AEA-1069-A2D8-08002B30309D}",
		ControlPanel: "{5399E694-6CE5-4D6C-8FCE-1D8870FDCBA0}",
		Network:      "{F02C1A0D-BE21-4350-88B0-7367FC96EF3C}",
		Recycle:      "{645FF040-5081-101B-9F08-00AA002F954E}",
		Mydocument:   "{59031a47-3f72-44a7-89c5-5595fe6b30ee}",
	}

	reg, _ := registry.OpenKey(
		registry.CURRENT_USER,
		"Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\HideDesktopIcons\\NewStartPanel",
		registry.WRITE,
	)
	defer reg.Close()

	for _, icon := range icons {
		if name, ok := regname[icon]; ok {
			reg.SetDWordValue(name, 0)
		}
	}
}

func installWinrar(fpath string, version WinrarVersion) {
	fname := map[WinrarVersion]string{
		Winrar501x64: "WinRAR_5.01_x64_SC.exe",
		Winrar501x86: "WinRAR_5.01_x86_SC.exe",
	}
	abspath := filepath.Join(fpath, fname[version])

	cmd := new(exec.Cmd)
	switch version {
	case Winrar501x64, Winrar501x86:
		cmd = exec.Command(abspath, "/s")
	default:
		return
	}
	cmd.Run()
}

func installAdobepdf(fpath string, version AdobepdfVersion) {
	fname := map[AdobepdfVersion]string{
		AcroRdrDC157: "AcroRdrDC1500720033_zh_CN.exe",
	}
	abspath := filepath.Join(fpath, fname[version])

	cmd := new(exec.Cmd)
	switch version {
	case AcroRdrDC157:
		cmd = exec.Command(abspath, "/sALL", "/norestart", "/msi", "ALLUSERS=1", "EULA_ACCEPT=YES")
	default:
		return
	}
	cmd.Run()
}

func installOffice(fpath string, version OfficeVersion) {
	fname := map[OfficeVersion]string{
		Office2007x86: "office2007pro.chs\\setup.exe",
		Office2010x86: "office2010pro.chs\\setup.exe",
	}

	fmt.Println(fname[version])

}

func connectWIFI() {

}

func activateWindows() {

}

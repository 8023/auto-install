package main

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

// DesktopIcon 枚举桌面图标类型
type DesktopIcon int

// DesktopIcon 枚举分别代表桌面上的五个图标
const (
	MyComputer DesktopIcon = iota
	ControlPanel
	Network
	Recycle
	Mydocument
)

func main() {
	fmt.Println("test")
	showDesktopIcon(MyComputer, ControlPanel, Mydocument)
}

func showDesktopIcon(icons ...DesktopIcon) {
	regmap := map[DesktopIcon]string{
		MyComputer:   "{20D04FE0-3AEA-1069-A2D8-08002B30309D}",
		ControlPanel: "{5399E694-6CE5-4D6C-8FCE-1D8870FDCBA0}",
		Network:      "{F02C1A0D-BE21-4350-88B0-7367FC96EF3C}",
		Recycle:      "{645FF040-5081-101B-9F08-00AA002F954E}",
		Mydocument:   "{59031a47-3f72-44a7-89c5-5595fe6b30ee}",
	}

	reg, _ := registry.OpenKey(
		registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Explorer\HideDesktopIcons\NewStartPanel`,
		registry.WRITE,
	)

	for _, icon := range icons {
		if name, ok := regmap[icon]; ok {
			reg.SetDWordValue(name, 0)
		}
	}
}

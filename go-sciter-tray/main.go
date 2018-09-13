package main

import (
	"fmt"
	"os"
	"time"

	context "github.com/asmexie/golearn/go-sciter-tray/context"
	"github.com/asmexie/golearn/go-sciter-tray/tray"
	"github.com/fatih/color"
)

func settingFunc() {
	fmt.Println(color.GreenString("Setting clicked"))
}
func aboutFunc() {
	fmt.Println(color.GreenString("About clicked"))
}
func updateFunc() {
	fmt.Println(color.GreenString("Update clicked"))
}
func exitFunc() {
	fmt.Println(color.GreenString("Exit clieked"))
	os.Exit(1)
}
func gotray() {
	menu := context.Menu{
		Items: []context.MenuItem{
			{
				Text:          "Settings",
				ClickCallback: settingFunc,
			},
			{
				Text:          "About",
				ClickCallback: aboutFunc,
			},
			{
				Text:          "Check For Update",
				ClickCallback: updateFunc,
			},
			{
				Text:          "Exit",
				ClickCallback: exitFunc,
			},
		},
	}
	trayIcon := tray.ClickableIcon{
		IconData: context.IconData,
		ClickHandler: func(x, y int, rightClick bool) {
			fmt.Println("x", x, "y", y, "isRightClick", rightClick)
			if rightClick {
				menu.DisplayContextMenu(x, y, 100)
			}
		},
	}
	trayIcon.Initialise()
	fmt.Println("Clickable Icon initialised")
	for {
		time.Sleep(time.Second)
	}
}
func dowhile() {
	for {
		fmt.Println("dowhile running...")
		time.Sleep(time.Second * 5)
	}
}
func main() {
	go gotray()
	dowhile()
}

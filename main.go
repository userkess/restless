package main

import (
	"fmt"
	"log"
	"os"

	. "modernc.org/tk9.0"
)

func main() {
	status, displayStatus := getStatus()
	statusLabel := Label(Txt(displayStatus))
	statusButton := Button(Txt(status), Command(func() { changeStatus() }))
	Pack(statusLabel)
	Pack(statusButton)
	Pack(Button(Txt("Exit"), Command(func() { Destroy(App) })))
	App.Wait()
}

func getStatus() (bool, string) {
	content, err := os.ReadFile("restless.cfg")
	if err != nil {
		log.Fatal(err)
	}
	if string(content) == "true\n" {
		fmt.Print("its true")
		return true, "Restless is active"
	} else {
		fmt.Print("its false")
		return false, "Restless is not active"
	}
}

func changeStatus() (bool, string) {
	s, d := getStatus()
	var writeData string
	if !s {
		fmt.Println("changing to true")
		writeData = "true\n"
		writeStatus([]byte(writeData))
		//		sleepCmd := exec.Command("sudo", "echo", "pmset disablesleep 1")
		//		if err := sleepCmd.Run(); err != nil {
		//			panic(err)
		//		}
	} else {
		fmt.Println("changing to false")
		writeData = "false\n"
		writeStatus([]byte(writeData))
		//		sleepCmd := exec.Command("sudo", "echo", "pmset disablesleep 0")
		//		if err := sleepCmd.Run(); err != nil {
		//			panic(err)
	}
	return s, d
}

func writeStatus(data []byte) {

	file, err := os.OpenFile("restless.cfg", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}

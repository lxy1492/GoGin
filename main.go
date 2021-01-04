package main

import "GoGin/LxyWeb"
import "GoGin/Python"

func RunLxyWeb()  {
	LxyWeb.Index()
}

func RunPython()  {
	Python.CmdPython("avc", "xxx")
}

func main() {
	RunLxyWeb()
}

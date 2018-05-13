package main

import (
	//"fmt"
	"os"

	"github.com/therecipe/qt/widgets"
	"ui"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	pilot := ui.IntPilot()
	pilot.Start()
	widgets.QApplication_Exec()
	for pilot.Przelogowanie {
		pilot = ui.IntPilot()
		pilot.Start()
		widgets.QApplication_Exec()

	}
}

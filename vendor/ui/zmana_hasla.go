package ui

import (
	//"fmt"

	//"fmt"

	"encoding/json"
	"mstr"
	"restclient"
	"strconv"
	"strings"
	"tools"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func (p *Pilot) zmiana_hasla() {
	okno_zh := widgets.NewQDialog(p, 0)
	okno_zh.SetAttribute(core.Qt__WA_DeleteOnClose, true)
	okno_zh.SetWindowTitle("Zmiana Hasła Logowania")
	grid := widgets.NewQGridLayout(nil)
	layoutH := widgets.NewQHBoxLayout()
	layoutV := widgets.NewQVBoxLayout()
	okno_zh.SetLayout(layoutV)
	labelho := widgets.NewQLabel(okno_zh, 0)
	labelho.SetText("Stare Hasło: ")
	labelhn := widgets.NewQLabel(okno_zh, 0)
	editho := widgets.NewQLineEdit(okno_zh)
	editho.SetEchoMode(2)
	edithn := widgets.NewQLineEdit(okno_zh)
	edithn.SetEchoMode(2)
	labelhn.SetText("Nowe Hasło: ")
	grid.AddWidget(labelho, 0, 0, 0)
	grid.AddWidget(editho, 0, 1, 0)
	grid.AddWidget(labelhn, 1, 0, 0)
	grid.AddWidget(edithn, 1, 1, 0)

	buttonZmien := widgets.NewQPushButton2("Zmien Hasło Logowania", okno_zh)
	buttonZmien.ConnectClicked(func(checked bool) {
		//p.Db.Reconect()
		nowe_haslo := mstr.ZmianaHasla{}
		nowe_haslo.Stare = tools.Passhash(editho.Text())
		nowe_haslo.Nowe = tools.Passhash(edithn.Text())
		i, _ := strconv.Atoi(strings.Split(mstr.AktywnaSesja.Sessiontopgoid, ":")[0])

		nowe_haslo.Id = i
		jnowehaslo, _ := json.Marshal(nowe_haslo)
		odp, _ := restclient.ZmienHaslo("/passwd", string(jnowehaslo))
		//fmt.Println(odp)
		zwrotka := mstr.Returner{}
		json.Unmarshal([]byte(odp), &zwrotka)
		//:= Db.Zmien_haslo(p.UserZ.id, editho.Text(), edithn.Text())
		if zwrotka.Affected > 0 {
			//info := widgets.NewQMessageBox(okno_zh)
			//info.SetWindowTitle("Uwaga!!")
			//info.SetText("Hasło Zostało zmienione!")
			widgets.QMessageBox_Information(okno_zh, "Uwaga", "Hasło Zostało zmienione!",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			//info.Show()
			okno_zh.Close()
		} else {
			//info := widgets.NewQMessageBox(okno_zh)
			//info.SetWindowTitle("Uwaga!!")
			//info.SetText("Bład Zmiany Hasła!")
			//info.Show()
			widgets.QMessageBox_Warning(okno_zh, "Uwaga", "Bład Zmiany Hasła!",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

		}
	})
	buttonAnuluj := widgets.NewQPushButton2("Anuluj", okno_zh)
	buttonAnuluj.ConnectClicked(func(checked bool) {
		okno_zh.Close()
	})
	layoutH.AddWidget(buttonZmien, 0, 0)
	layoutH.AddWidget(buttonAnuluj, 0, 0)
	layoutV.AddLayout(grid, 0)
	layoutV.AddLayout(layoutH, 0)
	okno_zh.Exec()
}

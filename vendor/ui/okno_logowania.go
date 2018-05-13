package ui

import (
	//"github.com/therecipe/qt/core"
	//	"github.com/therecipe/qt/gui"
	"fmt"

	"encoding/json"
	"mstr"
	"restclient"
	"tools"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func (p *Pilot) initOknoLogowania() {
	p.okno_logowania = widgets.NewQDialog(nil, 0)
	p.okno_logowania.SetWindowIcon(gui.NewQIcon5(":/qml/haslo.png"))
	p.okno_logowania.SetAttribute(core.Qt__WA_DeleteOnClose, true)
	//var haslob int
	//var status bool
	login_textline := widgets.NewQLineEdit(p.okno_logowania)
	login_textline.SetText("admin")
	haslo_textline := widgets.NewQLineEdit(p.okno_logowania)
	haslo_textline.SetText("")
	grid := widgets.NewQGridLayout(nil)
	haslo_textline.SetEchoMode(2)
	p.okno_logowania.SetLayout(grid)
	login_label := widgets.NewQLabel(p.okno_logowania, 0)
	login_label.SetText("Login: ")
	haslo_label := widgets.NewQLabel(p.okno_logowania, 0)
	haslo_label.SetText("Hasło: ")

	grid.AddWidget(login_label, 0, 1, 0)
	grid.AddWidget(login_textline, 0, 2, 0)
	grid.AddWidget(haslo_label, 1, 1, 0)
	grid.AddWidget(haslo_textline, 1, 2, 0)
	haslo_textline.ConnectReturnPressed(func() {
		putDaneLogowania := mstr.SingleSessionUser{}
		putDaneLogowania.Login = login_textline.Text()
		putDaneLogowania.Haslo = tools.Passhash(haslo_textline.Text())
		//		fmt.Println(login_textline.Text(), haslo_textline.Text())
		put, _ := json.Marshal(putDaneLogowania)
		odp, err := restclient.Logon("/logon", string(put))
		//fmt.Println(string(put))
		//fmt.Println(odp)
		get := mstr.Session{}
		json.Unmarshal([]byte(odp), &get)
		if err != nil {
			widgets.QMessageBox_Warning(p.okno_logowania, "Uwaga Błąd!!", fmt.Sprintf("Bład: "+err.Error()),
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

		} else {

			if len(get.Sessiontopgoid) > 0 {
				mstr.AktywnaSesja = get
				p.okno_logowania.Close()
				//p.LoadUi()

				p.LoadUi()
			} else {

				widgets.QMessageBox_Warning(p.okno_logowania, "Uwaga !!", fmt.Sprintf("Błędne Hasło!"),
					widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			}

		}

	})
	//		status, p.UserZ.imie, p.UserZ.nazwisko, p.UserZ.id = Db.Sprawdz_haslo(login_textline.Text(), haslo_textline.Text())
	//		if status {
	//			fmt.Println(p.UserZ.imie, p.UserZ.nazwisko, p.UserZ.id)
	//			fmt.Println("zalogowany")

	//			p.okno_logowania.Close()
	//			p.Zaloguj()

	//		} else {
	//			haslob += 1

	//			widgets.QMessageBox_Warning(p.okno_logowania, "Uwaga!!", fmt.Sprintf("Błędne Hasło!\nPróba: %d.", haslob),
	//				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	//			if haslob >= 3 {
	//				p.okno_logowania.Close()
	//			}

	//		}

	//	})
	login_textline.ConnectReturnPressed(haslo_textline.SetFocus2)
	//buttonZaloguj := widgets.NewQPushButton2("Zaloguj", p.okno_logowania)
	//buttonZaloguj.ConnectClicked(func(checked bool) {
	//		haslo_textline.ReturnPressed()
	//	})

	//	buttonAnuluj := widgets.NewQPushButton2("Anuluj", p.okno_logowania)
	//	buttonAnuluj.ConnectClicked(func(checked bool) {
	//		p.okno_logowania.Close()
	//	})
	//grid.AddWidget(buttonAnuluj, 2, 1, 0)
	//grid.AddWidget(buttonZaloguj, 2, 2, 2)
	//layout.AddWidget(buttonL, 0, core.Qt__AlignCenter)
	p.okno_logowania.SetWindowTitle("Logowanie")
	p.okno_logowania.SetWindowIcon(gui.NewQIcon5(":/qml/haslo.png"))
	p.okno_logowania.SetMinimumSize2(280, 90)
	p.okno_logowania.Show()
}

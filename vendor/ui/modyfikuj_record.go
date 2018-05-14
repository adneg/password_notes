package ui

import (
	//"fmt"

	"encoding/json"
	"fmt"
	"mstr"
	"restclient"
	"strconv"
	"tools"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	//"gitlab.top/ksala/pamietnik_hasel/restclient"
)

func (p *Pilot) modyfikuj_record() {

	okno_mod_r := widgets.NewQDialog(p, 0)
	okno_mod_r.SetAttribute(core.Qt__WA_DeleteOnClose, true)
	okno_mod_r.SetWindowTitle("Okno modyfikacji")

	labelOpis := widgets.NewQLabel(okno_mod_r, 0)
	labelOpis.SetText("Opis: ")
	editOpis := widgets.NewQLineEdit(okno_mod_r)
	editOpis.SetText(p.treePassword.CurrentItem().Text(1))

	labelLogin := widgets.NewQLabel(okno_mod_r, 0)
	labelLogin.SetText("Login: ")
	editLogin := widgets.NewQLineEdit(okno_mod_r)
	l, _ := tools.Decrypt(tools.Key, p.treePassword.CurrentItem().Text(102))
	editLogin.SetText(l)

	labelHaslo := widgets.NewQLabel(okno_mod_r, 0)
	labelHaslo.SetText("Hasło: ")
	editHaslo := widgets.NewQLineEdit(okno_mod_r)
	h, _ := tools.Decrypt(tools.Key, p.treePassword.CurrentItem().Text(103))
	editHaslo.SetText(h)

	labelAdnotacje := widgets.NewQLabel(okno_mod_r, 0)
	labelAdnotacje.SetText("Adnotacje: ")
	editAdnotacje := widgets.NewQLineEdit(okno_mod_r)
	editAdnotacje.SetText(p.treePassword.CurrentItem().Text(4))
	//NewAdnotacje(p)
	//editAdnotacje.SetMax(500)
	//editAdnotacje.ConnectTextChanged(editAdnotacje.Ogranicz)

	buttonZapisz := widgets.NewQPushButton2("Zapisz", okno_mod_r)
	buttonZapisz.SetShortcut(gui.NewQKeySequence2("Ctrl+Z", gui.QKeySequence__NativeText))
	buttonZapisz.ConnectClicked(func(checked bool) {
		//p.Db.Reconect()
		nowy_record := mstr.PasswordRecord{}
		nowy_record.Id, _ = strconv.Atoi(p.treePassword.CurrentItem().Text(0))

		nowy_record.Opis = editOpis.Text()
		nowy_record.Login, _ = tools.Encrypt(tools.Key, editLogin.Text())
		nowy_record.Haslo, _ = tools.Encrypt(tools.Key, editHaslo.Text())
		//nowy_record.Login = editLogin.Text()
		//nowy_record.Haslo = editHaslo.Text()
		nowy_record.Adnotacje = editAdnotacje.Text()
		jnowy_record, _ := json.Marshal(nowy_record)
		//fmt.Println(string(jnowy_record))
		odp, err := restclient.DodajRecord("/ph/update", string(jnowy_record))
		if err != nil {
			widgets.QMessageBox_Warning(okno_mod_r, "Uwaga Błąd!!", fmt.Sprintf("Bład: "+err.Error()),
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

		} else {
			zwrotka := mstr.Returner{}
			//			fmt.Println(odp)
			json.Unmarshal([]byte(odp), &zwrotka)
			if zwrotka.Affected > 0 {
				nowy_record.Id = zwrotka.Id
				p.UpdateTreeItem(nowy_record)
				//p.AddTreeItem(nowy_record)
				okno_mod_r.Close()
				widgets.QMessageBox_Information(okno_mod_r, "Potwierdzenie!!", fmt.Sprintf("Info: Record ZOSTAŁ zmodyfikowany!"),
					widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				//fmt.Println("dodano_record")

			} else {
				widgets.QMessageBox_Warning(okno_mod_r, "Uwaga Błąd!!", fmt.Sprintf("Bład: Record NIE został zmodyfikowany!"),
					widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				//fmt.Println("nie dodano_record")
			}

		}
		//fmt.Println(odp)

	})

	//buttonZapisz.shortcu
	grid := widgets.NewQGridLayout(okno_mod_r)
	grid.AddWidget(labelOpis, 0, 0, 0)
	grid.AddWidget(editOpis, 0, 1, 0)

	grid.AddWidget(labelLogin, 1, 0, 0)
	grid.AddWidget(editLogin, 1, 1, 0)

	grid.AddWidget(labelHaslo, 2, 0, 0)
	grid.AddWidget(editHaslo, 2, 1, 0)

	grid.AddWidget(labelAdnotacje, 3, 0, 0)
	grid.AddWidget(editAdnotacje, 3, 1, 0)
	grid.AddWidget(buttonZapisz, 4, 1, 0)

	okno_mod_r.Show()
}

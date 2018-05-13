package ui

import (
	//"fmt"

	"encoding/json"
	"fmt"
	"mstr"
	"restclient"
	//TOOLS DO SZYFROWANIA
	"tools"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	//"gitlab.top/ksala/pamietnik_hasel/restclient"
)

func (p *Pilot) dodaj_record() {
	okno_dod_r := widgets.NewQDialog(p, 0)
	okno_dod_r.SetAttribute(core.Qt__WA_DeleteOnClose, true)
	okno_dod_r.SetWindowTitle("Okno Dodawania Haseł")

	labelOpis := widgets.NewQLabel(okno_dod_r, 0)
	labelOpis.SetText("Opis: ")
	editOpis := widgets.NewQLineEdit(okno_dod_r)

	labelLogin := widgets.NewQLabel(okno_dod_r, 0)
	labelLogin.SetText("Login: ")
	editLogin := widgets.NewQLineEdit(okno_dod_r)

	labelHaslo := widgets.NewQLabel(okno_dod_r, 0)
	labelHaslo.SetText("Hasło: ")
	editHaslo := widgets.NewQLineEdit(okno_dod_r)

	labelAdnotacje := widgets.NewQLabel(okno_dod_r, 0)
	labelAdnotacje.SetText("Adnotacje: ")
	editAdnotacje := widgets.NewQLineEdit(okno_dod_r)
	//editAdnotacje := NewAdnotacje(p)
	//editAdnotacje.SetMax(500)
	//editAdnotacje.ConnectTextChanged(editAdnotacje.Ogranicz)
	buttonZapisz := widgets.NewQPushButton2("Zapisz", okno_dod_r)
	buttonZapisz.SetShortcut(gui.NewQKeySequence2("Ctrl+Z", gui.QKeySequence__NativeText))
	buttonZapisz.ConnectClicked(func(checked bool) {
		//p.Db.Reconect()
		nowy_record := mstr.PasswordRecord{}
		nowy_record.Opis = editOpis.Text()
		//NIE SZYFROWANIE
		//		nowy_record.Login = editLogin.Text()
		//		nowy_record.Haslo = editHaslo.Text()
		//SZYFROWANIE
		nowy_record.Login, _ = tools.Encrypt(tools.Key, editLogin.Text())
		nowy_record.Haslo, _ = tools.Encrypt(tools.Key, editHaslo.Text())
		//		testD, _ := tools.Decrypt(tools.Key, nowy_record.Login)
		//		fmt.Println(nowy_record.Login, editLogin.Text(), testD)

		nowy_record.Adnotacje = editAdnotacje.Text()
		jnowy_record, _ := json.Marshal(nowy_record)
		//fmt.Println(string(jnowy_record))
		odp, err := restclient.DodajRecord("/ph/add", string(jnowy_record))
		if err != nil {
			widgets.QMessageBox_Warning(okno_dod_r, "Uwaga Błąd!!", fmt.Sprintf("Bład: "+err.Error()),
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

		} else {
			zwrotka := mstr.Returner{}
			//fmt.Println(odp)
			json.Unmarshal([]byte(odp), &zwrotka)
			if zwrotka.Affected > 0 {
				nowy_record.Id = zwrotka.Id
				//z := p.treePassword.InvisibleRootItem()
				p.AddTreeItem(nowy_record)
				//fmt.Println("to jest obecy zaznaczony item: ", *p.treePassword.CurrentItem())
				//				s := fmt.Sprintf("%s", *p.treePassword.CurrentItem())
				//				fmt.Println(s)
				//				if fmt.Sprintf("%s", p.treePassword.CurrentItem()) == "&{<nil>}" {
				//					fmt.Println("TO BY BYŁO TO?")
				//				}

				okno_dod_r.Close()
				widgets.QMessageBox_Information(okno_dod_r, "Potwierdzenie!!", fmt.Sprintf("Info: Record ZOSTAŁ dodany!"),
					widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				//fmt.Println("dodano_record")
				z := p.treePassword.InvisibleRootItem()
				if z.ChildCount() > 0 {
					p.klAction.SetDisabled(false)
					p.khAction.SetDisabled(false)
					p.ksAction.SetDisabled(false)
					p.usuAction.SetDisabled(false)
					p.modAction.SetDisabled(false)
				}
				if z.ChildCount() == 1 {
					p.treePassword.SelectAll()
				}
			} else {
				widgets.QMessageBox_Warning(okno_dod_r, "Uwaga Błąd!!", fmt.Sprintf("Bład: Record NIE został dodany!"),
					widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				//fmt.Println("nie dodano_record")
			}

		}
		//fmt.Println(odp)

	})

	//buttonZapisz.shortcu
	grid := widgets.NewQGridLayout(okno_dod_r)
	grid.AddWidget(labelOpis, 0, 0, 0)
	grid.AddWidget(editOpis, 0, 1, 0)

	grid.AddWidget(labelLogin, 1, 0, 0)
	grid.AddWidget(editLogin, 1, 1, 0)

	grid.AddWidget(labelHaslo, 2, 0, 0)
	grid.AddWidget(editHaslo, 2, 1, 0)

	grid.AddWidget(labelAdnotacje, 3, 0, 0)
	grid.AddWidget(editAdnotacje, 3, 1, 0)
	grid.AddWidget(buttonZapisz, 4, 1, 0)
	okno_dod_r.Exec()

}

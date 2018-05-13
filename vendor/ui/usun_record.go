package ui

import (
	"encoding/json"
	"fmt"
	"mstr"
	"restclient"
	"strconv"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	//"gitlab.top/ksala/pamietnik_hasel/restclient"
)

func (p *Pilot) usun_record() {
	okno_usu_r := widgets.NewQDialog(p, 0)
	okno_usu_r.SetAttribute(core.Qt__WA_DeleteOnClose, true)
	okno_usu_r.SetWindowTitle("Okno Usuwania")
	buttonTak := widgets.NewQPushButton2("Tak", okno_usu_r)
	buttonNie := widgets.NewQPushButton2("Nie", okno_usu_r)
	buttonNie.ConnectClicked(func(checked bool) {
		okno_usu_r.Close()
	})
	buttonTak.ConnectClicked(func(checked bool) {
		nowy_record := mstr.PasswordRecord{}
		nowy_record.Id, _ = strconv.Atoi(p.treePassword.CurrentItem().Text(0))
		nowy_record.Opis = p.treePassword.CurrentItem().Text(1)
		nowy_record.Login = p.treePassword.CurrentItem().Text(2)
		nowy_record.Haslo = p.treePassword.CurrentItem().Text(3)
		nowy_record.Adnotacje = p.treePassword.CurrentItem().Text(4)

		jnowy_record, _ := json.Marshal(nowy_record)
		odp, err := restclient.DodajRecord("/ph/remove", string(jnowy_record))
		if err != nil {
			widgets.QMessageBox_Warning(okno_usu_r, "Uwaga Błąd!!", fmt.Sprintf("Bład: "+err.Error()),
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

		} else {
			zwrotka := mstr.Returner{}
			//fmt.Println(odp)
			json.Unmarshal([]byte(odp), &zwrotka)
			if zwrotka.Affected > 0 {
				z := p.treePassword.InvisibleRootItem()

				if z.ChildCount() == 1 {
					//p.ksAction.SetChecked(false)

					p.khAction.SetDisabled(true)
					p.klAction.SetDisabled(true)
					p.ksAction.SetDisabled(true)
					p.usuAction.SetDisabled(true)
					p.modAction.SetDisabled(true)
				}
				z.RemoveChild(p.treePassword.CurrentItem())
				okno_usu_r.Close()

				//fmt.Println("to jest obecy zaznaczony item: ", p.treePassword.CurrentItem())
				widgets.QMessageBox_Information(okno_usu_r, "Potwierdzenie!!", fmt.Sprintf("Info: Record ZOSTAŁ usunienty!"),
					widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

			} else {
				widgets.QMessageBox_Warning(okno_usu_r, "Uwaga Błąd!!", fmt.Sprintf("Bład: Record NIE został USUNIENTY!"),
					widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
				//fmt.Println("nie dodano_record")
			}

		}
		//fmt.Println(odp, err)

	})
	layout := widgets.NewQHBoxLayout()
	layout.AddWidget(buttonTak, 0, 0)
	layout.AddWidget(buttonNie, 0, 0)
	okno_usu_r.SetLayout(layout)
	okno_usu_r.Exec()
}

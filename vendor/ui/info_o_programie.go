package ui

import (
	//"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	//"gitlab.top/ksala/pamietnik_hasel/restclient"
)

func (p *Pilot) info_o_programie() {

	okno_informacyjne := widgets.NewQDialog(p, 0)

	okno_informacyjne.SetAttribute(core.Qt__WA_DeleteOnClose, true)
	okno_informacyjne.SetWindowTitle("O Programie")
	okno_informacyjne.SetMinimumSize2(400, 400)
	//okno_informacyjne.SetMaximumSize2(400, 400)
	// sierodek
	texti := widgets.NewQLabel(okno_informacyjne, 0)
	texti.SetText("Notatnik Haseł")
	texti.SetStyleSheet("QLabel {color: black;font-size: 16pt; font-weight: bold;}")
	texti2 := widgets.NewQLabel(okno_informacyjne, 0)
	texti2.SetText("Aplikacja Do Notowania Haseł")
	//texti2.SetStyleSheet("QLabel {color: black;font-size: 16pt; font-weight: bold;}")
	layout := widgets.NewQVBoxLayout()

	//var layoutZas = widgets.NewVHBoxLayout()
	//var layoutLicencja = widgets.NewVHBoxLayout()
	okno_informacyjne.SetLayout(layout)
	//tab := widgets.NewQTabBar(okno_informacyjne)

	tab := widgets.NewQTabWidget(okno_informacyjne)
	//tab.SetTabPosition(2)
	info := widgets.NewQTabWidget(okno_informacyjne)
	layoutInfo := widgets.NewQVBoxLayout()
	info.SetLayout(layoutInfo)
	//info := widgets.NewQLabel(okno_informacyjne, 0)
	textE := widgets.NewQLabel(info, 0)

	textE.SetText("Elementy Systemu:")
	textE.SetStyleSheet("QLabel {color: black;font-size: 12pt; font-weight: bold;}")
	//	testEInfo := widgets.NewQPlainTextEdit(info)
	//	testEInfo.SetPlainText("Język Programowania: go1.9.4/nBaza Danych: sqlite3/nInterfejs:github.com/therecipe/qt /nInne: ")
	//	testEInfo.SetReadOnly(true)
	textJezyk := widgets.NewQLabel(info, 0)
	textJezyk.SetText("Język Programowania - go1.9.4")
	textDB := widgets.NewQLabel(info, 0)
	textDB.SetText("Baza Danych - sqlite3/restapi")
	textQt := widgets.NewQLabel(info, 0)
	textQt.SetText("Interfejs - github.com/therecipe/qt")
	textInne := widgets.NewQLabel(info, 0)
	textInne.SetText("Inne:")
	layoutInfo.AddWidget(textE, 0, 0)
	layoutInfo.AddWidget(textJezyk, 0, 0)
	layoutInfo.AddWidget(textDB, 0, 0)
	layoutInfo.AddWidget(textQt, 0, 0)
	layoutInfo.AddWidget(textInne, 0, 0)
	zaslugi := widgets.NewQTabWidget(okno_informacyjne)
	licencja := widgets.NewQTabWidget(okno_informacyjne)

	tab.AddTab(info, "Informacje")
	tab.AddTab(zaslugi, "Zaslugi")
	tab.AddTab(licencja, "Licencja")
	//tab.ShowMaximizedDefault()
	//tab.AddTab
	//tab.AddTab("test2")
	link := widgets.NewQLabel(okno_informacyjne, 0)
	link.SetText("<a href='mailto:kamil.sala@gmail.com'>kamil.sala@gmail.com</a>")
	link.SetOpenExternalLinks(true)
	layout.AddWidget(texti, 0, core.Qt__AlignCenter)
	layout.AddWidget(texti2, 0, core.Qt__AlignCenter)
	layout.AddWidget(tab, 0, 0)
	layout.AddWidget(link, 0, core.Qt__AlignCenter)
	okno_informacyjne.Exec()
}

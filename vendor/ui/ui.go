package ui

import (
	//"fmt"

	"restclient"

	"os"
	"tools"

	"github.com/therecipe/qt/core"
	//"encoding/json"
	"mstr"
	"strconv"
	"strings"

	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

//##################### PILOT I FUNKCJE INIT I WYLACZENIA
type Pilot struct {
	//*widgets.QMainWindow
	widgets.QMainWindow
	okno_logowania *widgets.QDialog
	menuAplikacja  *widgets.QMenu
	menuAkcja      *widgets.QMenu
	treePassword   *widgets.QTreeWidget
	//FontNew          *gui.QFont
	usuAction *widgets.QAction
	modAction *widgets.QAction
	khAction  *widgets.QAction
	klAction  *widgets.QAction
	ksAction  *widgets.QAction
	//dockAction       *widgets.QAction
	freshAction      *widgets.QAction
	Rdock            *widgets.QDockWidget
	TextWyszukiwania string
	Filter           *widgets.QLineEdit
	Przelogowanie    bool
	WybranaKolumna   int
	//Painter          *gui.QPainter
	//Schowek          *gui.QClipboard
}

func IntPilot() *Pilot {
	return NewPilot(nil, 0)
	//this := &Pilot{QMainWindow: widgets.NewQMainWindow(nil, 0)}
	//return this
}
func (p *Pilot) LoadUi() {
	//p.Painter = gui.NewQPainter()
	//p.Test.BeginNativePainting()
	//p.initRightDock()
	p.intmenubar()
	p.initTree()
	p.SetCentralWidget(p.treePassword)
	//p.FontNew = gui.NewQFont2("verdana", 12, 1, false)
	//p.FontNew.SetBold(true)
	//p.SetFont(p.FontNew)
	//p.Schowek = gui.QGuiApplication_Clipboard()
	//p.Schowek.SetText("test", 0)

	//p.treePassword.Key
	p.initFilterEdit()
	p.treePassword.ConnectKeyPressEvent(func(e *gui.QKeyEvent) {
		if e.Key() == 16777216 {
			p.treePassword.SetFocus2()
			p.Filter.SetText("")

			return
		}
		p.Filter.Show()
		p.Filter.KeyPressEvent(e)
	})
	p.ConnectResizeEvent(func(e *gui.QResizeEvent) {
		p.Filter.Move2(p.treePassword.Width()-120, p.treePassword.Height()-5)
		p.ResizeEventDefault(e)
	})

	p.ConnectMoveEvent(func(e *gui.QMoveEvent) {
		p.Filter.Move2(p.treePassword.Width()-120, p.treePassword.Height()-5)
		p.MoveEventDefault(e)

	})
	p.treePassword.ConnectKeyboardSearch(func(text string) {
		//fmt.Println(text)
	})
	z := p.treePassword.InvisibleRootItem()
	if z.ChildCount() == 0 {
		p.klAction.SetDisabled(true)
		p.khAction.SetDisabled(true)
		p.usuAction.SetDisabled(true)
		p.modAction.SetDisabled(true)
		p.ksAction.SetDisabled(true)
	}
	//p.initRightDock()
	//	p.treePassword.ConnectItemSelectionChanged(func() {
	//		fmt.Println(p.treePassword.CurrentItem().Text(0))
	//	})

	p.SetWindowTitle("Notatnik Haseł")
	p.SetWindowIcon(gui.NewQIcon5(":/qml/haslo.png"))
	//p.Test.End()
	p.ShowMaximized()

}
func (p *Pilot) Start() {
	if _, err := os.Stat("./settings.ini"); err == nil {
		setting := core.NewQSettings4("./settings.ini", core.QSettings__IniFormat, nil)
		host := setting.Value("host", core.NewQVariant()).ToString()
		port := setting.Value("port", core.NewQVariant()).ToString()
		crt := setting.Value("crt", core.NewQVariant()).ToString()

		//fmt.Println(host, port)
		if _, err := os.Stat(crt); os.IsNotExist(err) {
			widgets.QMessageBox_Critical(nil, "Uwaga!!", "Brak Pliku CRT.. lub nie wskazany! Przerywam!",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			p.Close()
			os.Exit(1)
		}
		key := setting.Value("key", core.NewQVariant()).ToString()
		if _, err := os.Stat(key); os.IsNotExist(err) {
			widgets.QMessageBox_Critical(nil, "Uwaga!!", "Brak Pliku KEY.. lub nie wskazany!  Przerywam!",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			p.Close()
			os.Exit(1)
		}
		ca := setting.Value("ca", core.NewQVariant()).ToString()
		if _, err := os.Stat(ca); os.IsNotExist(err) {
			widgets.QMessageBox_Critical(nil, "Uwaga!!", "Brak Pliku CA.. lub nie wskazany!  Przerywam!",
				widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			p.Close()
			os.Exit(1)
		}
		restclient.Adres = "https://" + host + ":" + port
		restclient.CreateClientTLS(ca, key, crt)
		p.initOknoLogowania()
	} else {
		widgets.QMessageBox_Critical(nil, "Uwaga!!", "Brak Pliku settings.ini! Przerywam!",
			widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		p.Close()
		os.Exit(1)
	}

}

//func (p *Pilot) initRightDock() {
//	p.Rdock = widgets.NewQDockWidget("Informacje", p, 0)
//	p.Rdock.SetAllowedAreas(core.Qt__RightDockWidgetArea)
//	p.AddDockWidget(core.Qt__RightDockWidgetArea, p.Rdock)
//	haslo_textline := widgets.NewQLineEdit(p)
//	haslo_textline.SetText("a")
//	grid := widgets.NewQGridLayout(p.Rdock)
//	grid.AddWidget(haslo_textline, 1, 2, 0)
//	p.Rdock.SetLayout(grid)
//	p.Rdock.Show()
//}
func (p *Pilot) intmenubar() {
	//Aplikacja
	p.initAplikacjaMenu()
	p.initAkcjaMenu()
	//p.initZarzadzanieMenu()
}

func (p *Pilot) initAkcjaMenu() {
	p.menuAkcja = p.MenuBar().AddMenu2("Akcja")
	//Modyfikuj
	//Dodaj
	//Usun
	p.modAction = p.menuAkcja.AddAction("Modyfikuj")
	p.modAction.SetShortcut(gui.NewQKeySequence2("Ctrl+M", gui.QKeySequence__NativeText))
	p.modAction.SetIcon(gui.NewQIcon5(":/qml/modyfikuj.png"))
	p.modAction.ConnectTriggered(
		func(checked bool) {
			if (p.treePassword.CurrentItem().IsHidden() == false) && (p.treePassword.CurrentItem().Type() > 0) {

				p.modyfikuj_record()
			}

			//p.zmiana_hasla()
		})
	dodAction := p.menuAkcja.AddAction("Dodaj")
	dodAction.SetShortcut(gui.NewQKeySequence2("Ctrl+D", gui.QKeySequence__NativeText))
	dodAction.SetIcon(gui.NewQIcon5(":/qml/dodaj.png"))
	dodAction.ConnectTriggered(
		func(checked bool) {
			p.dodaj_record()
			//p.zmiana_hasla()
		})
	p.usuAction = p.menuAkcja.AddAction("Usuń")
	p.usuAction.SetShortcut(gui.NewQKeySequence2("Ctrl+R", gui.QKeySequence__NativeText))
	p.usuAction.SetIcon(gui.NewQIcon5(":/qml/usun.png"))
	p.usuAction.ConnectTriggered(
		func(checked bool) {
			if (p.treePassword.CurrentItem().IsHidden() == false) && (p.treePassword.CurrentItem().Type() > 0) {

				p.usun_record()
			}
			//p.zmiana_hasla()
		})

	p.klAction = p.menuAkcja.AddAction("Kopiuj Login")
	p.klAction.SetShortcut(gui.NewQKeySequence2("Ctrl+L", gui.QKeySequence__NativeText))
	p.klAction.SetIcon(gui.NewQIcon5(":/qml/copy.png"))
	p.klAction.ConnectTriggered(
		func(checked bool) {
			if p.treePassword.CurrentItem().IsHidden() == false && (p.treePassword.CurrentItem().Type() > 0) {
				l, _ := tools.Decrypt(tools.Key, p.treePassword.CurrentItem().Text(102))
				//fmt.Println("kopiujemy do schowka:", l)
				//if err != nil {
				schowek := gui.QGuiApplication_Clipboard()
				schowek.SetText(l, 0)
				//}
			}
			//p.usun_record()
			//p.zmiana_hasla()
		})
	p.khAction = p.menuAkcja.AddAction("Kopiuj Hasło")
	p.khAction.SetShortcut(gui.NewQKeySequence2("Ctrl+H", gui.QKeySequence__NativeText))
	p.khAction.SetIcon(gui.NewQIcon5(":/qml/copy.png"))
	p.khAction.ConnectTriggered(
		func(checked bool) {
			if p.treePassword.CurrentItem().IsHidden() == false && (p.treePassword.CurrentItem().Type() > 0) {

				h, _ := tools.Decrypt(tools.Key, p.treePassword.CurrentItem().Text(103))
				//fmt.Println("kopiujemy do schowka:", h)
				//if err != nil {

				schowek := gui.QGuiApplication_Clipboard()
				schowek.SetText(h, 0)
				//}
			}
			//p.usun_record()
			//p.zmiana_hasla()
		})
	//	p.dockAction = p.menuAkcja.AddAction("Informacje dodatkowe")
	//	p.dockAction.SetCheckable(true)
	//	p.dockAction.SetShortcut(gui.NewQKeySequence2("Ctrl+I", gui.QKeySequence__NativeText))
	//	p.dockAction.SetIcon(gui.NewQIcon5(":/qml/szukaj.png"))
	p.ksAction.ConnectTriggered(
		func(checked bool) {
			if checked {
				p.Rdock.Show()

			} else {
				p.Rdock.Hide()
			}
		})
	p.ksAction = p.menuAkcja.AddAction("Odszyfruj")
	p.ksAction.SetCheckable(true)
	p.ksAction.SetShortcut(gui.NewQKeySequence2("Ctrl+O", gui.QKeySequence__NativeText))
	p.ksAction.SetIcon(gui.NewQIcon5(":/qml/oko.png"))
	p.ksAction.ConnectTriggered(
		func(checked bool) {
			p.treePassword.SetSortingEnabled(false)
			if checked {
				p.Odszyfruj(p.treePassword.InvisibleRootItem())
			} else {
				p.Zaszyfruj(p.treePassword.InvisibleRootItem())
			}
			p.treePassword.SetSortingEnabled(true)
		})
	p.freshAction = p.menuAkcja.AddAction("Odśwież")
	p.freshAction.SetShortcut(gui.NewQKeySequence2("Ctrl+F", gui.QKeySequence__NativeText))
	p.freshAction.SetIcon(gui.NewQIcon5(":/qml/fresh.png"))
	p.freshAction.ConnectTriggered(func(checked bool) {
		if p.Filter.Text() != "" {
			p.Filter.SetText("")
		}

		p.treePassword.Clear()
		daneSJ := []mstr.PasswordRecord{}
		daneSJ, _ = restclient.PobierzRekordy("/ph/getall")

		if len(daneSJ) > 0 {
			for _, v := range daneSJ {
				p.AddTreeItem(v)

			}
			p.klAction.SetDisabled(false)
			p.khAction.SetDisabled(false)
			p.usuAction.SetDisabled(false)
			p.modAction.SetDisabled(false)
			p.ksAction.SetDisabled(false)

		} else {
			p.klAction.SetDisabled(true)
			p.khAction.SetDisabled(true)
			p.usuAction.SetDisabled(true)
			p.modAction.SetDisabled(true)
			p.ksAction.SetDisabled(true)
		}

	})

}
func (p *Pilot) initAplikacjaMenu() {
	p.menuAplikacja = p.MenuBar().AddMenu2("Aplikacja")
	chpassAction := p.menuAplikacja.AddAction("Zmien Hasło")
	chpassAction.SetIcon(gui.NewQIcon5(":/qml/user.png"))
	chpassAction.ConnectTriggered(
		func(checked bool) {
			p.zmiana_hasla()
		})

	logoutAction := p.menuAplikacja.AddAction("Wyloguj")
	logoutAction.SetIcon(gui.NewQIcon5(":/qml/zmiana_usera.png"))
	logoutAction.ConnectTriggered(
		func(checked bool) {
			p.Wylacz(true)
		})
	closeAction := p.menuAplikacja.AddAction("Wyłącz")
	closeAction.SetShortcut(gui.NewQKeySequence2("Ctrl+Q", gui.QKeySequence__NativeText))
	closeAction.SetIcon(gui.NewQIcon5(":/qml/wyjdz.png"))

	closeAction.ConnectTriggered(
		func(checked bool) {
			p.Wylacz(false)
		})
	p.menuAplikacja.AddSeparator()
	oprogramieAction := p.menuAplikacja.AddAction("O Programie")
	oprogramieAction.SetIcon(gui.NewQIcon5(":/qml/haslo.png"))
	oprogramieAction.ConnectTriggered(
		func(checked bool) {
			p.info_o_programie()
		})
	//Widok
}

func (p *Pilot) Wylacz(flaga bool) {
	//fmt.Println("Akcja Wylacz Aplikacje")
	p.Przelogowanie = flaga
	//fmt.Println(p.Przelogowanie)
	p.Close()

}

func (p *Pilot) initTree() {

	p.treePassword = widgets.NewQTreeWidget(p)
	p.treePassword.SetRootIsDecorated(false)
	p.treePassword.Header().SetSectionsClickable(true)
	p.treePassword.Header().SetContextMenuPolicy(core.Qt__CustomContextMenu)
	p.treePassword.Header().ConnectSectionClicked(
		func(ind int) {

			//fmt.Println(ind)
		})
	p.treePassword.Header().ConnectCustomContextMenuRequested(
		func(ind *core.QPoint) {
			p.treePassword.Header().SetSectionResizeMode(2)
			p.Filter.SetText("")
			nrh := p.treePassword.IndexAtDefault(ind).Column()

			if nrh > -1 {
				p.SelectedColumn(nrh)

			}
			p.treePassword.Header().SetSectionResizeMode(3)
			//nrh := p.treePassword.IndexAt(ind).Column()

			//p.Filter.SetText("")
			//p.ShowAllItems(p.treePassword.InvisibleRootItem())
		})
	//p.treePassword.SetWordWrap(true)
	p.treePassword.SetSortingEnabled(true)
	daneSJ := []mstr.PasswordRecord{}
	daneSJ, _ = restclient.PobierzRekordy("/ph/getall")
	for _, v := range daneSJ {
		p.AddTreeItem(v)

	}

	//for i := 0; i < 4; i++ {
	//	p.treePassword.ResizeColumnToContents(i)
	//}

	p.treePassword.SetHeaderLabels([]string{"Id", "Opis", "Login", "Hasło", "Adnotacje"})

	//p.treePassword.Header().SetStretchLastSection(false)
	//p.treePassword.SetHorizontalScrollBarPolicy(2)
	//p.treePassword.SetVerticalScrollBarPolicy(2)
	p.SelectedColumn(0)
	//TEGO SZUKAŁEM  TEGO SZUKAŁEM  TEGO SZUKAŁEM  TEGO SZUKAŁEM
	p.treePassword.Header().SetSectionResizeMode(3)
	//p.treePassword.SetColumnWidth(0, 40)

}

func (p *Pilot) SetHeaderColor(color string, ico string) {
	hi := widgets.NewQTreeWidgetItem2([]string{"Id", "Opis", "Login", "Hasło", "Adnotacje"}, 0)
	qb := gui.NewQBrush()
	qb.SetColor(gui.NewQColor6(color))
	//hi.SetBackground(p.WybranaKolumna, qb)
	p.treePassword.SetHeaderItem(hi)
	p.treePassword.HeaderItem().SetBackground(p.WybranaKolumna, qb)
	p.treePassword.HeaderItem().SetIcon(p.WybranaKolumna, gui.NewQIcon5(ico))
	//p.treePassword.SetStyleSheet("QHeaderView::1 {background-color: rgb(28, 28, 28); color: rgb(215, 215, 215);}")
}
func (p *Pilot) SelectedColumn(nrh int) {

	p.WybranaKolumna = nrh
	p.SetHeaderColor("green", ":/qml/oko.png")

}
func (p *Pilot) initFilterEdit() {
	p.Filter = widgets.NewQLineEdit(p)
	p.Filter.SetMinimumSize2(115, 25)
	p.Filter.SetMaximumSize2(115, 25)
	//p.Filter.ConnectHideEvent()
	p.Filter.SetHidden(true)
	p.Filter.ConnectTextChanged(func(text string) {
		//fmt.Println(text)

		if text != "" {
			p.LookforItems(p.treePassword.InvisibleRootItem())
		}
		if text == "" {
			p.ShowAllItems(p.treePassword.InvisibleRootItem())
			p.SetHeaderColor("green", ":/qml/oko.png")
		}

	})
	p.Filter.ConnectKeyPressEvent(func(e *gui.QKeyEvent) {
		if e.Key() == 16777216 {
			p.Filter.SetText("")
			//p.ShowAllItems(p.treePassword.InvisibleRootItem())
			//p.ShowAllItems(p.treePassword.InvisibleRootItem())
			p.treePassword.SetFocus2()
			return
		}
		p.Filter.Show()
		p.Filter.SetFocus2()
		p.Filter.KeyPressEventDefault(e)
		p.treePassword.KeyPressEventDefault(e)

	})

	p.Filter.ConnectFocusOutEvent(func(e *gui.QFocusEvent) {
		p.Filter.SetHidden(true)

		p.Filter.FocusOutEventDefault(e)
	})
}

func (p *Pilot) Zaszyfruj(root *widgets.QTreeWidgetItem) {
	for i := 0; i < root.ChildCount(); i++ {
		item := root.Child(i)
		item.SetText(2, "**********")
		item.SetText(3, "**********")
		p.Zaszyfruj(item)
	}

}
func (p *Pilot) Odszyfruj(root *widgets.QTreeWidgetItem) {
	// SORTOWANIE PSUJE ODSZYFROWYWANIE JESLI WYBRANO KOLUMNE
	// KTÓRA MA BYC ODSZYFORWANA JAKO KOLUMNE DO SORTOWANIA
	//p.treePassword.SetSortingEnabled(false)
	for i := 0; i < root.ChildCount(); i++ {
		item := root.Child(i)
		l, _ := tools.Decrypt(tools.Key, item.Text(102))
		h, _ := tools.Decrypt(tools.Key, item.Text(103))
		item.SetText(2, l)
		item.SetText(3, h)

		p.Odszyfruj(item)
	}
	//p.treePassword.SetSortingEnabled(true)

}
func (p *Pilot) LookforItems(root *widgets.QTreeWidgetItem) {
	p.SetHeaderColor("red", ":/qml/okor.png")
	for i := 0; i < root.ChildCount(); i++ {
		item := root.Child(i)
		//fmt.Println("item: ", item.Text(p.WybranaKolumna), "  nr:", i)
		if strings.Contains(item.Text(p.WybranaKolumna), p.Filter.Text()) {
			item.SetHidden(false)
			qb := gui.NewQBrush()
			qb.SetColor(gui.NewQColor6("red"))
			item.SetForeground(p.WybranaKolumna, qb)
			item.SetBackground(p.WybranaKolumna, qb)
		} else {
			item.SetHidden(true)

		}

		p.LookforItems(item)
	}
}

func (p *Pilot) ShowAllItems(root *widgets.QTreeWidgetItem) {
	//child = root.ChildCount()
	for i := 0; i < root.ChildCount(); i++ {
		item := root.Child(i)

		item.SetHidden(false)
		qb := gui.NewQBrush()
		qb.SetColor(gui.NewQColor6("black"))
		//qb.SetColor(gui.NewQColor3(50, 205, 50, 5))
		item.SetForeground(p.WybranaKolumna, qb)
		item.SetBackground(p.WybranaKolumna, qb)
		p.ShowAllItems(item)
	}
	//fmt.Println(root)
	//fmt.Println("testS")
}
func (p *Pilot) AddTreeItem(v mstr.PasswordRecord) {
	//p.treePassword.ClearSelection()
	itemP := widgets.NewQTreeWidgetItem3(p.treePassword, 1)
	itemP.SetText(0, strconv.Itoa(v.Id))
	itemP.SetText(1, v.Opis)
	itemP.SetText(102, v.Login)
	itemP.SetText(103, v.Haslo)
	if p.ksAction.IsChecked() {

		l, _ := tools.Decrypt(tools.Key, itemP.Text(102))
		h, _ := tools.Decrypt(tools.Key, itemP.Text(103))
		itemP.SetText(2, l)
		itemP.SetText(3, h)

	} else {

		itemP.SetText(2, "**********")
		itemP.SetText(3, "**********")
	}
	itemP.SetText(4, v.Adnotacje)
	p.treePassword.SetCurrentItem(itemP)
	// trzeba sprawdzic czy sie da scrola zrobic gdyby bylo duzo hasel?
	//p.treePassword.ScrollToItem()

}

func (p *Pilot) UpdateTreeItem(v mstr.PasswordRecord) {
	itemP := p.treePassword.CurrentItem()
	itemP.SetText(1, v.Opis)
	itemP.SetText(102, v.Login)
	itemP.SetText(103, v.Haslo)
	if p.ksAction.IsChecked() {

		l, _ := tools.Decrypt(tools.Key, itemP.Text(102))
		h, _ := tools.Decrypt(tools.Key, itemP.Text(103))
		itemP.SetText(2, l)
		itemP.SetText(3, h)

	} else {

		itemP.SetText(2, "**********")
		itemP.SetText(3, "**********")
	}

	itemP.SetText(4, v.Adnotacje)
}

//##################### ADNOTACJE  W OKNIE DODAAWANIA I MODYFIKACJI

//type Adnotacje struct {
//	//*widgets.QMainWindow
//	widgets.QTextEdit
//	Max int
//}

//func (a *Adnotacje) SetMax(m int) {
//	a.Max = m
//}

//func (a *Adnotacje) Ogranicz() {
//	if len(a.ToPlainText()) > a.Max {
//		a.TextCursor().DeletePreviousChar()
//	}
//}

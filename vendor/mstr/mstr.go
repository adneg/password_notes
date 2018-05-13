package mstr

// wszystkie struktury danych w tableach
type User struct {
	ID_UZYTKOWNIKA int    `gorm:"primary_key;AUTO_INCREMENT:index"`
	EMAIL          string `gorm:"type:varchar(300);unique_index;not null"`
	HASLO          string `gorm:"type:varchar(300);unique_index;not null"`
	IMIE           string `gorm:"type:varchar(300);unique_index;not null"`
	NAZWISKO       string `gorm:"type:varchar(300);unique_index;not null"`
}

func (User) TableName() string {
	return "UZYTKOWNICY"
}

type Uprawnienia_Domyslne struct {
	ID_UPRAWNIENIA   int    `gorm:"primary_key;AUTO_INCREMENT:index"`
	ID_RODZICA       int    `gorm:"type:bigint REFERENCES UPRAWNIENIA_DOMYSLNE(ID_UPRAWNIENIA)"`
	FUNKCJA          string `gorm:"type:varchar(300);unique_index;not null"`
	WARTOSC_DOMYSLNA int    `gorm:"type:TINYINT(1);not null"`
	OPIS             string `gorm:"type:varchar(500);unique_index;not null"`
}

func (Uprawnienia_Domyslne) TableName() string {
	return "UPRAWNIENIA_DOMYSLNE"
}

type Uprawnienia_Nadane struct {
	// sqlite
	ID_UPRAWNIENIA_FK int `gorm:"type:bigint REFERENCES UPRAWNIENIA_DOMYSLNE(ID_UPRAWNIENIA) ON DELETE CASCADE ON UPDATE CASCADE"`
	ID_UZYTKOWNIKA_FK int `gorm:"type:bigint REFERENCES UZYTKOWNICY(ID_UZYTKOWNIKA) ON DELETE CASCADE ON UPDATE CASCADE"`

	//ID_UPRAWNIENIA_FK int
	//ID_UZYTKOWNIKA_FK int
	WARTOSC int `gorm:"type:TINYINT(1);not null"`
}

func (Uprawnienia_Nadane) TableName() string {
	return "UPRAWNIENIA_NADANE"
}

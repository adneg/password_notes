package mstr

type PasswordRecord struct {
	Id   int `gorm:"primary_key;AUTO_INCREMENT:index"`
	Opis string
	//Opis zmienic na tytuł
	Login     string
	Haslo     string
	Adnotacje string
	//Ip        string

}

func (PasswordRecord) TableName() string {
	return "NOTATNIK_HASEL"
}

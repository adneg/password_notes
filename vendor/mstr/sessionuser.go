package mstr

import (
	"time"
)

//obsluga sessji
var (
	AktywnaSesja Session
)

type Session struct {
	Sessiontopgoid string
}

type SingleSessionUser struct {
	Id       int
	Imie     string
	Nazwisko string
	Login    string
	Haslo    string
	Addr     string
	Czas     time.Time
}

func (s SingleSessionUser) CopyToNewRecord() (n SingleSessionUser) {
	n = s
	return
}

func NewSingleSessionUser(Id int, Imie string,
	Nazwisko string, Login string, Haslo string,
	Addr string, Czas time.Time) *SingleSessionUser {
	return &SingleSessionUser{Id: Id, Imie: Imie, Nazwisko: Nazwisko, Login: Login, Haslo: Haslo, Addr: Addr, Czas: Czas}
}

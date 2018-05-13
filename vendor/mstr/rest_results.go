package mstr

// wszystkie rodzaje rezultatów z restapi
import (
	"time"
)

type AccesInfo struct {
	Id              int
	IdR             int
	Funkcja         string
	Wartosc         int
	Opis            string
	WartoscDomyslna int
	Status          int
	// JESLI STATUS >-1 I WartoscDomyslna < Wartosc PRAWA ZOSTAŁY UZNANE
	// JESLI -1 TO WARTOSC DOMYSLNA NIE ZOSTAŁA NADPISANA
	// JESLI STATUS >-1 I WartoscDomyslna > Wartosc PRAWA ZOSTAŁY ZABRANE
	// PRZYDATNE PRZY DOBRANIU PRAWIDLOWYCH CZYNNOSCI PRZY MODUFIKACJI POŹNIEJSZEJ - ORAZ KOLOROWANIU
}

type Version struct {
	Numer float32
	Data  time.Time
}

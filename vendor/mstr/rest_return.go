package mstr

// wszystkie rodzaje returnow z restapi
type ReturnUpdate struct {
	//Status   bool
	Affected int64
	// jezeli bool false to wystapil blad
	//jezeli Affected 0 nic nie zapisano
}

type Returner struct {
	//Status   bool
	Id       int
	Affected int64

	// jezeli bool false to wystapil blad
	//jezeli Affected 0 nic nie zapisano
}

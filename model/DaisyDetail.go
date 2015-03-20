package model

type DaisyDetail struct {
	Id      int
	DaisyId int    `sql:"not null"`
	Part    int    `sql:"not null"`
	Status  string `sql:"not null"`
	Notes   string `sql:"not null"`
}

func (dd DaisyDetail) TableName() string {
	return "daisydetail"
}

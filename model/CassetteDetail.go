package model

type CassetteDetail struct {
	Id         int
	CassetteId int    `sql:"not null"`
	Part       int    `sql:"not null"`
	Status     string `sql:"not null"`
	Notes      string `sql:"not null"`
}

func (bb CassetteDetail) TableName() string {
	return "cassettedetail"
}

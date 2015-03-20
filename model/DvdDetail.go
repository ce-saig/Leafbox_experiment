package model

type DvdDetail struct {
	Id     int
	DvdId  int    `sql:"not null"`
	Part   int    `sql:"not null"`
	Status string `sql:"not null"`
	Notes  string `sql:"not null"`
}

func (d DvdDetail) TableName() string {
	return "dvddetail"
}

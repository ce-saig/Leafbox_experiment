package model

type CdDetail struct {
	Id      int
	CdId    int    `sql:"not null"`
	Part    int    `sql:"not null"`
	Status  string `sql:"not null"`
	Notes   string `sql:"not null"`
	TrackFr int    `sql:"not null"`
	TrackTo int    `sql:"not null"`
}

func (cd CdDetail) TableName() string {
	return "cddetail"
}

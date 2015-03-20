package model

type Member struct {
	Id           int
	Name         string
	Gender       string
	Address      string
	District     string
	PostCode     string
	PhoneNo      string
	MemberStatus string
	col9         string
	col10        string
	col11        string
}

func (m Member) TableName() string {
	return "member"
}

package model

type Leaf struct {
	Id       int64 `json:"id"`
	DomainId int64 `json:"domain_id"`
	MaxId    int64 `json:"max_id"`
	Step     int64 `json:"step"`
	Status   int   `json:"status"`
}

func (*Leaf) GetTableName() string {
	return "leaf"
}

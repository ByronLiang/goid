package model

type Leaf struct {
	Id       int64 `json:"id"`
	DomainId int64 `json:"domain_id,omitempty"`
	MaxId    int64 `json:"max_id,omitempty"`
	Step     int64 `json:"step,omitempty"`
	Status   int   `json:"status,omitempty"`
}

func (*Leaf) GetTableName() string {
	return "leaf"
}

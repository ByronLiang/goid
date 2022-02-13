package db

import (
	"github.com/ByronLiang/goid/pkg/model"
)

const (
	StatusOn  = 1
	StatusOff = 2
)

var LeafDao = &leafDao{}

type leafDao struct {
}

func (*leafDao) Create(leaf *model.Leaf) error {
	return DB.Create(&leaf).Error
}

func (*leafDao) GetByDomainId(domainId int64) (*model.Leaf, error) {
	res := new(model.Leaf)
	err := DB.Where("domain_id = ?", domainId).First(&res).Error
	return res, err
}

func (*leafDao) GetAll() ([]*model.Leaf, error) {
	res := make([]*model.Leaf, 0)
	err := DB.Where("status = ?", StatusOn).Find(&res).Error
	return res, err
}

func (*leafDao) UpdateMaxId(originMaxId, domainId, currentMaxId int64) int64 {
	return DB.Exec("update leaf set max_id = ? where domain_id = ? and max_id = ?", currentMaxId, domainId, originMaxId).RowsAffected
}

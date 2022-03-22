package db

import (
	"github.com/ByronLiang/goid/pkg/model"
	"github.com/jinzhu/gorm"
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

func FilterDomainId(domainIds ...int64) FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("domain_id in (?)", domainIds)
	}
}

func FilterStatus(status int) FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", status)
	}
}

func (*leafDao) GetLeaf(filters ...FilterFunc) (res []*model.Leaf, err error) {
	var mDb *gorm.DB
	mDb = GetMDB()
	for _, filter := range filters {
		mDb = filter(mDb)
	}
	err = mDb.Find(&res).Error
	return
}

func (*leafDao) GetStatusOnLeaf(status int, domainIds ...int64) (res []*model.Leaf, err error) {
	mDb := GetMDB()
	if len(domainIds) > 0 {
		mDb = mDb.Where("domain_id in (?)", domainIds)
	}
	if status > 0 {
		mDb = mDb.Where("status = ?", status)
	}
	err = mDb.Find(&res).Error
	return
}

func (*leafDao) UpdateMaxId(originMaxId, domainId, currentMaxId int64) int64 {
	return DB.Exec("update leaf set max_id = ? where domain_id = ? and max_id = ?", currentMaxId, domainId, originMaxId).RowsAffected
}

func (*leafDao) UpdateStatus(leaf model.Leaf) (int64, error) {
	mDb := DB.Exec("update leaf set status = ? where id = ?", leaf.Status, leaf.Id)
	return mDb.RowsAffected, mDb.Error
}

package model

import (
	"yoyo-mall/pkg/errno"
)

type CategoryModel struct {
	ID       uint32
	ParentID uint32 // 父类目id，0则为根类目
	Name     string
	Image    string // 类目图片
}

func (c *CategoryModel) TableName() string {
	return "category"
}

func (c *CategoryModel) Create() error {
	return DB.Self.Create(c).Error
}

func (c *CategoryModel) Save() error {
	return DB.Self.Save(c).Error
}

func GetCategoryByID(id uint32) (*CategoryModel, error) {
	m := &CategoryModel{}
	d := DB.Self.First(m, "id = ?", id)
	if d.RecordNotFound() {
		return m, errno.ErrRecordNotFound
	}
	return m, nil
}

func GetCategoryList() ([]*CategoryModel, error) {
	list := make([]*CategoryModel, 0)
	err := DB.Self.Find(&list).Error
	return list, err
}

func GetCid2(cid uint32) ([]*CategoryModel, error) {
	list := make([]*CategoryModel, 0)
	err := DB.Self.Where("parent_id = ?", cid).Find(&list).Error
	return list, err
}

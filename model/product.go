package model

import (
	"errors"
	"time"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/util"

	"gorm.io/gorm"
)

type ProductModel struct {
	ID          uint32
	Cid         uint32    // 一级类目
	Cid2        uint32    // 二级类目
	Title       string    // 商品标题
	BookName    string    // 书名
	Author      string    // 作者
	Publisher   string    // 出版社
	Price       float32   // 原价
	CurPrice    float32   // 优惠价
	Stock       int       // 库存
	Detail      string    // 详情，暂时空着占位，没想好怎么做
	Images      string    // 图片，分号分割
	Status      int8      // 状态，0正常，1下架
	PublishTime time.Time // 出版时间
	CreateTime  time.Time
	IsDeleted   bool
	DeleteTime  *time.Time
}

func (p *ProductModel) TableName() string {
	return "product"
}

func (p *ProductModel) Create() error {
	p.CreateTime = util.GetCurrentTime()
	return DB.Self.Create(p).Error
}

func (p *ProductModel) Save() error {
	return DB.Self.Save(p).Error
}

func GetProductByID(id uint32) (*ProductModel, error) {
	model := &ProductModel{}
	d := DB.Self.Where("is_deleted = 0").Where("id = ?", id).First(model)
	if errors.Is(d.Error, gorm.ErrRecordNotFound) {
		return nil, errno.ErrRecordNotFound
	}
	return model, d.Error
}

func DeleteProduct(id uint32) (err error) {
	deleteTime := util.GetCurrentTime()
	err = DB.Self.Model(&ProductModel{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"is_deleted": 1, "delete_time": deleteTime}).
		Error
	return
}

type ProductItemModel struct {
	ID          uint32
	Title       string
	Author      string
	Publisher   string
	BookName    string
	Cid         uint32
	Cid2        uint32
	Price       float32
	CurPrice    float32 `gorm:"column:cur_price"` // column后不能由空格
	Images      string
	SaleNum     int
	CommentNum  int
	CommentRate float32
	Score       float32
	PublishTime time.Time `gorm:"column:publish_time"`
}

func ProductList(sql string) ([]*ProductItemModel, error) {
	products := make([]*ProductItemModel, 0)
	err := DB.Self.Raw(sql).Scan(&products).Error
	return products, err
}

func ProductSearch(sql string) ([]*ProductItemModel, error) {
	list := make([]*ProductItemModel, 0)
	err := DB.Self.Raw(sql).Scan(&list).Error
	return list, err
}

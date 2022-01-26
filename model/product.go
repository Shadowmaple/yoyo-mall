package model

import "time"

type ProductModel struct {
	ID          uint32
	Cid         uint32     // 一级类目
	Cid2        uint32     // 二级类目
	Title       string     // 商品标题
	BookName    string     // 书名
	Author      string     // 作者
	Publisher   string     // 出版社
	Price       float32    // 原价
	CurPrice    float32    // 优惠价
	Stock       int        // 库存
	Detail      string     // 详情，暂时空着占位，没想好怎么做
	Images      string     // 图片，分号分割
	Status      int8       // 状态，0正常，1下架
	PublishTime *time.Time // 出版时间
	CreateTime  *time.Time
	UpdateTime  *time.Time
	IsDeleted   bool
	DeleteTime  *time.Time
}

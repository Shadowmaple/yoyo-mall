package order

import (
	"yoyo-mall/model"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/service/logistics"
	"yoyo-mall/util"
)

func New(userID uint32, req *NewOrderItem) (orderID uint32, err error) {
	now := util.GetCurrentTime()
	order := &model.OrderModel{
		UserID:      userID,
		Status:      0,
		Payment:     req.Payment,
		Freight:     req.Freight,
		TotalFee:    req.TotalFee,
		Coupon:      req.Coupon,
		ReceiveName: req.ReceiveName,
		ReceiveTel:  req.ReceiveTel,
		ReceiveAddr: req.ReceiveAddr,
		Refund:      "",
		OrderCode:   "",
		CreateTime:  now,
	}

	productRecords := make([]*model.OrderProductModel, 0)

	// 开启事务，创建订单记录和订单-商品记录

	tx := model.DB.Self.Begin()

	if err = tx.Create(order).Error; err != nil {
		tx.Rollback()
		return
	}

	orderID = order.ID

	for _, product := range req.Products {
		productRecords = append(productRecords, &model.OrderProductModel{
			OrderID:    orderID,
			ProductID:  product.ID,
			Num:        product.Num,
			Price:      product.Price,
			CurPrice:   product.CurPrice,
			TotalFee:   product.TotalFee,
			Image:      product.Image,
			CreateTime: now,
		})
	}

	if err = tx.Create(&productRecords).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}

// 变更订单状态
// expectedStatus 只能为1/2/3/5/6，即完成付款、已发货、已签收、取消订单、退货
func UpdateStatus(id uint32, expectedStatus int8) error {
	record, err := model.GetOrderByID(id)
	if err != nil {
		return err
	}

	record.Status = expectedStatus
	now := util.GetCurrentTime()
	var logisticState int8

	switch expectedStatus {
	case 1:
		// 变为已付款/待发货，物流信息
		record.PayTime = now
		logisticState = 0
	case 2:
		// 变为已发货/待收货
		record.DeliverTime = now
		logisticState = 1
	case 3:
		// 变为已签收/待评价
		record.ConfirmTime = now
		logisticState = 2
	case 5:
		// 未支付下取消订单
		logisticState = 3
	case 6:
		// 变为退货
		logisticState = 4
	default:
		return errno.ErrOrderExpectedStatus
	}

	if err := record.Save(); err != nil {
		return err
	}

	// 插入新物流记录
	if err := logistics.NewLogistics(record.ID, logisticState); err != nil {
		return err
	}

	return nil
}

package order

import (
	"yoyo-mall/model"
	"yoyo-mall/util"
)

func List(userID uint32, limit, page, kind int) (list []*OrderItem, err error) {
	list = make([]*OrderItem, 0)

	// 订单状态
	// kind：0全部订单，1待付款，2待发货，3待收货，4待评价，5退款售后
	// status：0->待付款，1->待发货，2->待收货，3->待评价，4->交易完成，5->交易取消，6->退货中，7->交易关闭
	var status int8 = -1
	switch kind {
	case 1:
		status = 0
	case 2:
		status = 1
	case 3:
		status = 2
	case 4:
		status = 3
	case 5:
		status = 6
	}

	records, err := model.OrderList(userID, limit, limit*page, status)
	if err != nil {
		return
	}

	var item *OrderItem
	for _, record := range records {
		item, err = processRawOrder(record)
		if err != nil {
			return
		}

		list = append(list, item)
	}

	return
}

func GetInfo(id uint32) (item *OrderItem, err error) {
	record, err := model.GetOrderByID(id)
	if err != nil {
		return
	}

	item, err = processRawOrder(record)
	if err != nil {
		return
	}

	return
}

func processRawOrder(record *model.OrderModel) (item *OrderItem, err error) {
	products := make([]*ProductItem, 0)
	curRecords, err := model.GetProductByOrderID(record.ID)
	if err != nil {
		return
	}

	for _, record := range curRecords {
		products = append(products, &ProductItem{
			ID:       record.ProductID,
			Title:    record.Title,
			Author:   record.Author,
			Num:      record.Num,
			TotalFee: record.TotalFee,
			Price:    record.Price,
			CurPrice: record.CurPrice,
			Image:    record.Image,
		})
	}

	title := getGeneralTitle(products)
	image := getGeneralImage(products)

	item = &OrderItem{
		ID:          record.ID,
		Status:      record.Status,
		TotalFee:    record.TotalFee,
		Payment:     record.Payment,
		Coupon:      record.Coupon,
		Freight:     record.Freight,
		ReceiveName: record.ReceiveName,
		ReceiveTel:  record.ReceiveTel,
		ReceiveAddr: record.ReceiveAddr,
		OrderCode:   record.OrderCode,
		CreateTime:  util.GetStandardTime(record.CreateTime),
		PayTime:     util.GetStandardTime(record.PayTime),
		DeliverTime: util.GetStandardTime(record.DeliverTime),
		ConfirmTime: util.GetStandardTime(record.ConfirmTime),
		ProductNum:  len(products),
		Title:       title,
		Image:       image,
		Products:    products,
	}

	return
}

func getGeneralTitle(products []*ProductItem) string {
	if len(products) == 0 {
		return ""
	}
	if len(products) == 1 {
		return products[0].Title
	}

	// todo：应该根据bookName来
	s := products[0].Title + "等"

	return s
}

func getGeneralImage(products []*ProductItem) string {
	if len(products) == 0 {
		return ""
	}
	return products[0].Image
}

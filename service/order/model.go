package order

type OrderItem struct {
	ID          uint32         `json:"id"`
	Status      int8           `json:"status"`
	TotalFee    float32        `json:"total_fee"`
	Payment     float32        `json:"payment"`
	Coupon      float32        `json:"coupon"`
	Freight     float32        `json:"freight"`
	ReceiveName string         `json:"receive_name"`
	ReceiveTel  string         `json:"receive_tel"`
	ReceiveAddr string         `json:"receive_addr"`
	OrderCode   string         `json:"order_code"`
	CreateTime  string         `json:"create_time"`
	PayTime     string         `json:"pay_time"`
	DeliverTime string         `json:"deliver_time"`
	ConfirmTime string         `json:"confirm_time"`
	ProductNum  int            `json:"product_num"`
	Title       string         `json:"title"` // 展示的标题，由多件商品标题组合而成
	Image       string         `json:"image"` // 某个商品封面图片，多个商品取一个
	Products    []*ProductItem `json:"products"`
}

type ProductItem struct {
	ID       uint32  `json:"id"`
	Title    string  `json:"title"`
	Author   string  `json:"author"`
	Num      int     `json:"num"`
	TotalFee float32 `json:"total_fee"`
	Price    float32 `json:"price"`
	CurPrice float32 `json:"cur_price"`
	Image    string  `json:"image"`
}

type NewOrderItem struct {
	TotalFee    float32        `json:"total_fee"`
	Payment     float32        `json:"payment"`
	Coupon      float32        `json:"coupon"`
	Freight     float32        `json:"freight"`
	ReceiveName string         `json:"receive_name"`
	ReceiveTel  string         `json:"receive_tel"`
	ReceiveAddr string         `json:"receive_addr"`
	ProductNum  int            `json:"product_num"`
	Products    []*ProductItem `json:"products"`
}

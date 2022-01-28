package address

import (
	"yoyo-mall/model"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/util"
)

type AddressItem struct {
	ID         uint32 `json:"id"`
	Name       string `json:"name"`
	Tel        string `json:"tel"`
	Province   string `json:"province"` // 省
	City       string `json:"city"`     // 市
	District   string `json:"district"` // 区县
	Detail     string `json:"detail"`   // 街道详情
	IsDefault  bool   `json:"is_default"`
	CreateTime string `json:"create_time"`
}

func List(userID uint32) (list []*AddressItem, err error) {
	list = make([]*AddressItem, 0)

	records, err := model.AddressList(userID)
	if err != nil {
		if err == errno.ErrRecordNotFound {
			return list, nil
		}
		return
	}

	for _, item := range records {
		list = append(list, &AddressItem{
			ID:         item.ID,
			Name:       item.Name,
			Tel:        item.Tel,
			Province:   item.Province,
			City:       item.City,
			District:   item.District,
			Detail:     item.Detail,
			IsDefault:  item.IsDefault,
			CreateTime: util.GetStandardTime(item.CreateTime),
		})
	}

	return
}

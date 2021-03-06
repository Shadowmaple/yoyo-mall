package address

import (
	"yoyo-mall/model"
)

type AddressInfo struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Tel       string `json:"tel"`
	Province  string `json:"province"`
	City      string `json:"city"`
	District  string `json:"district"`
	Detail    string `json:"detail"`
	IsDefault bool   `json:"is_default"`
}

func Update(userID uint32, addr *AddressInfo) error {
	record, err := model.GetAddressByID(addr.ID)
	if err != nil {
		return err
	}

	// 是否需要更新默认地址
	changeDetfault := false
	if addr.IsDefault && !record.IsDefault {
		changeDetfault = true
	}

	record.Name = addr.Name
	record.Tel = addr.Tel
	record.Province = addr.Province
	record.City = addr.City
	record.District = addr.District
	record.Detail = addr.Detail
	record.IsDefault = addr.IsDefault
	if err := record.Save(); err != nil {
		return err
	}

	// 设置为了默认地址，则需要将其它地址都变为非默认地址
	if changeDetfault {
		if err := model.UpdateNotDefaultAddress(userID, addr.ID); err != nil {
			return err
		}
	}

	return nil
}

func Add(userID uint32, addr *AddressInfo) error {
	record := &model.AddressModel{
		UserID:    userID,
		Name:      addr.Name,
		Tel:       addr.Tel,
		Province:  addr.Province,
		City:      addr.City,
		District:  addr.District,
		Detail:    addr.Detail,
		IsDefault: addr.IsDefault,
	}
	if err := record.Create(); err != nil {
		return err
	}

	// 设置了默认地址
	if record.IsDefault {
		if err := model.UpdateNotDefaultAddress(userID, record.ID); err != nil {
			return err
		}
	}

	return nil
}

func Delete(userID, address uint32) error {
	return model.DeleteAddress(userID, address)
}

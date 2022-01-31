package category

import "yoyo-mall/model"

type Cid2Item struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type CidItem struct {
	ID   uint32      `json:"id"`
	Name string      `json:"name"`
	List []*Cid2Item `json:"list"`
}

func GetList() ([]*CidItem, error) {
	list := make([]*CidItem, 0)

	cids, err := model.GetCategoryList()
	if err != nil {
		return nil, err
	}

	for _, cid := range cids {
		if cid.ParentID != 0 {
			continue
		}
		models, err := model.GetCid2(cid.ID)
		if err != nil {
			return nil, err
		}
		cid2List := make([]*Cid2Item, 0, len(models))
		for _, model := range models {
			cid2List = append(cid2List, &Cid2Item{
				ID:    model.ID,
				Name:  model.Name,
				Image: model.Image,
			})
		}
		list = append(list, &CidItem{
			ID:   cid.ID,
			Name: cid.Name,
			List: cid2List,
		})
	}

	return list, nil
}

package ip

func RequestIPEntityToRequestIPModel(entity RequestIPEntity) RequestIPModel {
	return RequestIPModel{
		ID:        entity.ID,
		IP:        entity.IP,
		CreatedAt: entity.CreatedAt,
	}
}

func RequestIPModelToRequestIPEntity(model RequestIPModel) RequestIPEntity {
	return RequestIPEntity{
		ID:        model.ID,
		IP:        model.IP,
		CreatedAt: model.CreatedAt,
	}
}

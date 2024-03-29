package ip

func RequestIPEntityToRequestIPModel(entity RequestIPEntity) RequestIPModel {
	return RequestIPModel{
		ID:        entity.ID,
		IP:        entity.IP,
		CreatedOn: entity.CreatedOn,
	}
}

func RequestIPModelToRequestIPResponseDTO(model RequestIPModel) RequestIPResponseDTO {
	return RequestIPResponseDTO{
		ID:        model.ID,
		IP:        model.IP,
		CreatedOn: model.CreatedOn,
	}
}

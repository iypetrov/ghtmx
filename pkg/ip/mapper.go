package ip

func RequestIPEntityToRequestIPModel(entity RequestIPEntity) RequestIPModel {
	return RequestIPModel(entity)
}

func RequestIPModelToRequestIPEntity(model RequestIPModel) RequestIPEntity {
	return RequestIPEntity(model)
}

func RequestIPModelToRequestIPResponseDTO(model RequestIPModel) RequestIPResponseDTO {
	return RequestIPResponseDTO{
		ID:        model.ID.String(),
		IP:        model.IP,
		CreatedAt: model.CreatedAt,
	}
}

func StatsIPEntitiesToStatsIPModels(entities []StatsIPEntity) []StatsIPModel {
	models := make([]StatsIPModel, len(entities))
	for i, entity := range entities {
		models[i] = StatsIPModel(entity)
	}
	return models
}

func StatsIPModelsToStatsIPEntities(models []StatsIPModel) []StatsIPEntity {
	entities := make([]StatsIPEntity, len(models))
	for i, model := range models {
		entities[i] = StatsIPEntity(model)
	}
	return entities
}

func StatsIPModelsToStatsIPResponseDTOs(models []StatsIPModel) []StatsIPResponseDTO {
	dtos := make([]StatsIPResponseDTO, len(models))
	for i, model := range models {
		dtos[i] = StatsIPResponseDTO(model)
	}
	return dtos
}

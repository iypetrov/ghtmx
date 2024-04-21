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
		models[i] = StatsIPModel{
			IP:    entity.IP,
			Count: entity.Count,
		}
	}
	return models
}

func StatsIPModelsToStatsIPEntities(models []StatsIPModel) []StatsIPEntity {
	entities := make([]StatsIPEntity, len(models))
	for i, model := range models {
		entities[i] = StatsIPEntity{
			IP:    model.IP,
			Count: model.Count,
		}
	}
	return entities
}

func StatsIPModelsToStatsIPResponseDTOs(models []StatsIPModel) []StatsIPResponseDTO {
	dtos := make([]StatsIPResponseDTO, len(models))
	for i, model := range models {
		dtos[i] = StatsIPResponseDTO{
			IP:    model.IP,
			Count: model.Count,
		}
	}
	return dtos
}

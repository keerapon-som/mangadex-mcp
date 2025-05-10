package api

import (
	"mangadex_mcp/entities"
	"mangadex_mcp/packages"
	"mangadex_mcp/repository"
)

type MangaAPI struct {
	mangaApiPackage *packages.MangaDexMangaAPIManga
	mapRepo         *repository.MapNameWithIDRepo
}

func NewMangaAPI(mangaApiPackage *packages.MangaDexMangaAPIManga, mapRepo *repository.MapNameWithIDRepo) *MangaAPI {
	service := &MangaAPI{
		mangaApiPackage: mangaApiPackage,
		mapRepo:         mapRepo,
	}

	return service
}

func (m *MangaAPI) GetManga(paramsArguments entities.GetMangaArguments) (packages.MangaResponse, error) {

	for i, value := range paramsArguments.IncludedTags { // convert tag name to tag id
		paramsArguments.IncludedTags[i] = m.mapRepo.AllTagMapId()[value]
	}

	for i, value := range paramsArguments.ExcludedTags { // convert tag name to tag id
		paramsArguments.ExcludedTags[i] = m.mapRepo.AllTagMapId()[value]
	}

	return m.mangaApiPackage.GetManga(paramsArguments)
}

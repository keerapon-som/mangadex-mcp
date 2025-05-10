package repository

import (
	"mangadex_mcp/packages"
	"sync"
)

type MapRepo struct {
	mapTag    *MapTagRepo
	mapAuthor *MapAuthorRepo
}

type MapTagRepo struct {
	mu          sync.Mutex
	allTagMapId map[string]string
	listAllTags []string
}

type MapAuthorRepo struct {
	mu             sync.Mutex
	allAuthorMapId map[string]string
	listAllAuthors []string
}

type MapNameWithIDRepo struct {
	mangaPackage  *packages.MangaDexMangaAPIManga
	authorPackage *packages.MangaDexAuthorAPI

	mapRepo *MapRepo
}

func NewMapNameWithIDRepo(mangaPackage *packages.MangaDexMangaAPIManga, authorPackage *packages.MangaDexAuthorAPI) *MapNameWithIDRepo {

	repo := &MapNameWithIDRepo{
		mangaPackage:  mangaPackage,
		authorPackage: authorPackage,
		mapRepo: &MapRepo{
			mapTag: &MapTagRepo{
				mu:          sync.Mutex{},
				allTagMapId: make(map[string]string),
				listAllTags: make([]string, 0),
			},
			mapAuthor: &MapAuthorRepo{
				mu:             sync.Mutex{},
				allAuthorMapId: make(map[string]string),
				listAllAuthors: make([]string, 0),
			},
		},
	}

	repo.updateTags()
	// repo.updateAuthors()

	return repo
}

func (m *MapNameWithIDRepo) updateTags() error {
	tags, err := m.mangaPackage.GetTags()
	if err != nil {
		return err
	}

	m.mapRepo.mapTag.mu.Lock()
	defer m.mapRepo.mapTag.mu.Unlock()

	m.mapRepo.mapTag.allTagMapId = make(map[string]string)
	for _, tag := range tags.Data {
		m.mapRepo.mapTag.allTagMapId[tag.Attributes.Name.En] = tag.ID
		m.mapRepo.mapTag.listAllTags = append(m.mapRepo.mapTag.listAllTags, tag.Attributes.Name.En)
	}
	return nil
}

func (m *MapNameWithIDRepo) AllTagMapId() map[string]string {
	return m.mapRepo.mapTag.allTagMapId
}

func (m *MapNameWithIDRepo) ListAllTags() []string {
	return m.mapRepo.mapTag.listAllTags
}

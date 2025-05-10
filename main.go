package main

import (
	"mangadex_mcp/api"
	"mangadex_mcp/mcp"
	"mangadex_mcp/packages"
	"mangadex_mcp/repository"
)

func main() {
	MangaDexMangaAPI := packages.NewMangaDexMangaAPI("https://api.mangadex.dev/manga")
	MangaDexAuthorAPI := packages.NewMangaDexAuthorAPI("https://api.mangadex.dev/author")
	mapRepo := repository.NewMapNameWithIDRepo(MangaDexMangaAPI, MangaDexAuthorAPI)

	mcpRoute := mcp.NewMCPRoute(
		api.NewMangaAPI(MangaDexMangaAPI, mapRepo),
		mapRepo,
	)

	mcpRoute.Run()
}

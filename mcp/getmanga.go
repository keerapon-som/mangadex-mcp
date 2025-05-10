package mcp

import (
	"context"
	"encoding/json"
	"mangadex_mcp/entities"

	"github.com/mark3labs/mcp-go/mcp"
)

func (h *Handler) mangaTool() mcp.Tool {

	return mcp.NewTool("get_manga",
		mcp.WithDescription("Get manga from MangaDex"),
		mcp.WithNumber("limit",
			mcp.Description("Limit of data"),
		),
		mcp.WithNumber("offset",
			mcp.Description("Offset of data"),
		),
		mcp.WithString("title",
			mcp.Description("Title of manga"),
		),
		mcp.WithString("authorOrArtist",
			mcp.Description("uuid of author or artist"),
		),
		mcp.WithArray("authors",
			mcp.Description("array of authors"),
		),
		mcp.WithArray("artists",
			mcp.Description("array of artists"),
		),
		mcp.WithString("year",
			mcp.Description("Year of release or none"),
		),
		mcp.WithArray("includedTags",
			mcp.Description("array of included tags"),
			mcp.Enum(h.mapRepo.ListAllTags()...),
		),
		mcp.WithString("includedTagsMode",
			mcp.Description("Mode of included tags"),
			mcp.Enum("AND", "OR"),
		),
		mcp.WithArray("excludedTags",
			mcp.Description("array of excluded tags"),
			mcp.Enum(h.mapRepo.ListAllTags()...),
		),
		mcp.WithString("excludedTagsMode",
			mcp.Description("Mode of excluded tags"),
			mcp.Enum("AND", "OR"),
		),
		mcp.WithArray("status",
			mcp.Description("array of status"),
			mcp.Enum("ongoing", "completed", "hiatus", "cancelled"),
		),
		mcp.WithArray("originalLanguage",
			mcp.Description("array of original language"),
		),
		mcp.WithArray("excludedOriginalLanguage",
			mcp.Description("array of excluded original language"),
		),
		mcp.WithArray("availableTranslatedLanguage",
			mcp.Description("array of available translated language"),
		),
		mcp.WithArray("publicationDemographic",
			mcp.Description("array of publication demographic"),
			mcp.Enum("shounen", "shoujo", "seinen", "josei", "none"),
		),
		mcp.WithArray("ids",
			mcp.Description("array of ids"),
		),
		mcp.WithArray("contentRating",
			mcp.Description("array of content rating"),
			mcp.Enum("safe", "suggestive", "erotica", "pornographic"),
		),
		mcp.WithString("createdAtSince",
			mcp.Description("Created at since"),
		),
		mcp.WithString("updatedAtSince",
			mcp.Description("Updated at since"),
		),
		mcp.WithArray("includes",
			mcp.Description("array of includes"),
			mcp.Enum("manga", "author", "artist", "cover_art", "tag", "creator"),
		),
		mcp.WithString("hasAvailableChapters",
			mcp.Description("has available chapters"),
			mcp.Enum("true", "false"),
		),
		mcp.WithString("group",
			mcp.Description("group uuid"),
		),
	)

}

func (h *Handler) mangaHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	bytes, err := json.Marshal(request.Params.Arguments)
	if err != nil {
		return nil, err
	}

	var mangaArguments entities.GetMangaArguments
	if err := json.Unmarshal(bytes, &mangaArguments); err != nil {
		return nil, err
	}

	mangaResponse, err := h.MangaAPI.GetManga(mangaArguments)
	if err != nil {
		return nil, err
	}

	jsonMangaResponse, err := json.Marshal(mangaResponse)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(jsonMangaResponse)), nil
}

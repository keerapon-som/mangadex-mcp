package entities

type GetMangaArguments struct {
	Limit                       *int     `url:"limit,omitempty"`
	Offset                      *int     `url:"offset,omitempty"`
	Title                       string   `url:"title,omitempty"`
	AuthorOrArtist              string   `url:"authorOrArtist,omitempty"`
	Authors                     []string `url:"authors,omitempty,brackets"`
	Artists                     []string `url:"artists,omitempty,brackets"`
	Year                        string   `url:"year,omitempty"`
	IncludedTags                []string `url:"includedTags,omitempty,brackets"`
	IncludedTagsMode            string   `url:"includedTagsMode,omitempty"`
	ExcludedTags                []string `url:"excludedTags,omitempty,brackets"`
	ExcludedTagsMode            string   `url:"excludedTagsMode,omitempty"`
	Status                      []string `url:"status,omitempty,brackets"`
	OriginalLanguage            []string `url:"originalLanguage,omitempty,brackets"`
	ExcludedOriginalLanguage    []string `url:"excludedOriginalLanguage,omitempty,brackets"`
	AvailableTranslatedLanguage []string `url:"availableTranslatedLanguage,omitempty,brackets"`
	PublicationDemographic      []string `url:"publicationDemographic,omitempty,brackets"`
	Ids                         []string `url:"ids,omitempty,brackets"`
	ContentRating               []string `url:"contentRating,omitempty,brackets"`
	CreatedAtSince              string   `url:"createdAtSince,omitempty"`
	UpdatedAtSince              string   `url:"updatedAtSince,omitempty"`
	Includes                    []string `url:"includes,omitempty,brackets"`
	HasAvailableChapters        string   `url:"hasAvailableChapters,omitempty"`
	Group                       string   `url:"group,omitempty"`
}

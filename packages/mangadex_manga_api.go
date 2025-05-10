package packages

import (
	"encoding/json"
	"fmt"
	"mangadex_mcp/entities"
	"net/http"

	"github.com/google/go-querystring/query"
)

type MangaDexMangaAPIManga struct {
	baseurl string
}

func NewMangaDexMangaAPI(baseurl string) *MangaDexMangaAPIManga {
	return &MangaDexMangaAPIManga{
		baseurl: baseurl,
	}
}

// GetManga retrieves a list of manga based on the provided query parameters
func (m *MangaDexMangaAPIManga) GetManga(paramslist entities.GetMangaArguments) (MangaResponse, error) {

	params, err := query.Values(paramslist)
	if err != nil {
		return MangaResponse{}, fmt.Errorf("failed to encode query parameters: %w", err)
	}

	reqURL := fmt.Sprintf("%s?%s", m.baseurl, params.Encode())
	fmt.Println("รีเควส ยูอาแอว ", reqURL)
	resp, err := http.Get(reqURL)
	if err != nil {
		return MangaResponse{}, fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return MangaResponse{}, fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	var mangaResp MangaResponse
	if err := json.NewDecoder(resp.Body).Decode(&mangaResp); err != nil {
		return MangaResponse{}, fmt.Errorf("failed to decode response body: %w", err)
	}

	return mangaResp, nil
}

func (m *MangaDexMangaAPIManga) GetTags() (MangaTagsResponse, error) {
	reqURL := fmt.Sprintf("%s/tag", m.baseurl)
	fmt.Println("รีเควส ยูอาแอว tags ", reqURL)
	resp, err := http.Get(reqURL)
	if err != nil {
		return MangaTagsResponse{}, fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return MangaTagsResponse{}, fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	var mangaTagsResp MangaTagsResponse
	if err := json.NewDecoder(resp.Body).Decode(&mangaTagsResp); err != nil {
		return MangaTagsResponse{}, fmt.Errorf("failed to decode response body: %w", err)
	}

	return mangaTagsResp, nil
}

type MangaResponse struct {
	Result   string `json:"result"`
	Response string `json:"response"`
	Data     []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Title                  map[string]string   `json:"title"`
			AltTitles              []map[string]string `json:"altTitles"`
			Description            map[string]string   `json:"description"`
			IsLocked               bool                `json:"isLocked"`
			Links                  map[string]string   `json:"links"`
			OriginalLanguage       string              `json:"originalLanguage"`
			LastVolume             string              `json:"lastVolume"`
			LastChapter            string              `json:"lastChapter"`
			PublicationDemographic string              `json:"publicationDemographic"`
			Status                 string              `json:"status"`
			Year                   int                 `json:"year"`
			ContentRating          string              `json:"contentRating"`

			Tags []struct {
				ID         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					Name        map[string]string `json:"name"`
					Description map[string]string `json:"description"`
					Group       string            `json:"group"`
					Version     int               `json:"version"`
				} `json:"attributes"`
				Relationships []struct {
					ID         string `json:"id"`
					Type       string `json:"type"`
					Related    string `json:"related"`
					Attributes struct {
					} `json:"attributes"`
				} `json:"relationships"`
			} `json:"tags"`
			State                          string   `json:"state"`
			Version                        int      `json:"version"`
			CreatedAt                      string   `json:"createdAt"`
			UpdatedAt                      string   `json:"updatedAt"`
			ChapterNumbersResetOnNewVolume bool     `json:"chapterNumbersResetOnNewVolume"`
			AvailableTranslatedLanguages   []string `json:"availableTranslatedLanguages"`
			LatestUploadedChapter          string   `json:"latestUploadedChapter"`
		} `json:"attributes"`
		Relationships []struct {
			ID         string `json:"id"`
			Type       string `json:"type"`
			Related    string `json:"related"`
			Attributes struct {
			} `json:"attributes"`
		} `json:"relationships"`
	} `json:"data"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type MangaTagsResponse struct {
	Result   string `json:"result"`
	Response string `json:"response"`
	Data     []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name struct {
				En string `json:"en"`
			} `json:"name"`
			Description struct {
			} `json:"description"`
			Group   string `json:"group"`
			Version int    `json:"version"`
		} `json:"attributes"`
		Relationships []interface{} `json:"relationships"`
	} `json:"data"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

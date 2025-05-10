package packages

import (
	"encoding/json"
	"fmt"
	"mangadex_mcp/entities"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type MangaDexAuthorAPI struct {
	baseurl string
}

func NewMangaDexAuthorAPI(baseurl string) *MangaDexAuthorAPI {
	return &MangaDexAuthorAPI{
		baseurl: baseurl,
	}
}

// GetAuthor retrieves a list of author based on the provided query parameters
func (m *MangaDexAuthorAPI) GetAuthor(paramslist entities.GetAuthorArguments) (AuthorResponse, error) {

	params, err := query.Values(paramslist)
	if err != nil {
		return AuthorResponse{}, fmt.Errorf("failed to encode query parameters: %w", err)
	}

	reqURL := fmt.Sprintf("%s?%s", m.baseurl, params.Encode())
	fmt.Println("รีเควส ยูอาแอว ", reqURL)
	resp, err := http.Get(reqURL)
	if err != nil {
		return AuthorResponse{}, fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return AuthorResponse{}, fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	var authorResp AuthorResponse
	if err := json.NewDecoder(resp.Body).Decode(&authorResp); err != nil {
		return AuthorResponse{}, fmt.Errorf("failed to decode response body: %w", err)
	}

	return authorResp, nil
}

type AuthorResponse struct {
	Result   string `json:"result"`
	Response string `json:"response"`
	Data     []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name      string `json:"name"`
			ImageURL  string `json:"imageUrl"`
			Biography struct {
			} `json:"biography"`
			Twitter   string    `json:"twitter"`
			Pixiv     string    `json:"pixiv"`
			MelonBook string    `json:"melonBook"`
			FanBox    string    `json:"fanBox"`
			Booth     string    `json:"booth"`
			Namicomi  string    `json:"namicomi"`
			NicoVideo string    `json:"nicoVideo"`
			Skeb      string    `json:"skeb"`
			Fantia    string    `json:"fantia"`
			Tumblr    string    `json:"tumblr"`
			Youtube   string    `json:"youtube"`
			Weibo     string    `json:"weibo"`
			Naver     string    `json:"naver"`
			Website   string    `json:"website"`
			CreatedAt time.Time `json:"createdAt"`
			UpdatedAt time.Time `json:"updatedAt"`
			Version   int       `json:"version"`
		} `json:"attributes"`
		Relationships []struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"relationships"`
	} `json:"data"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

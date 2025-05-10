package entities

type GetAuthorArguments struct {
	Limit    *int     `url:"limit,omitempty"`
	Offset   *int     `url:"offset,omitempty"`
	Ids      []string `url:"ids,omitempty,brackets"`
	Name     string   `url:"name,omitempty"`
	Includes []string `url:"includes,omitempty,brackets"`
}

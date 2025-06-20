package hub

type repositoryResponse struct {
	PullCount   int    `json:"pull_count"`
	StarCount   int    `json:"star_count"`
	LastUpdated string `json:"last_updated"`
}

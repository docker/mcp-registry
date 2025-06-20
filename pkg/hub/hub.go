package hub

import (
	"context"
	"encoding/json"
	"net/http"
)

func GetRepositoryInfo(ctx context.Context, repo string) (*repositoryResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://hub.docker.com/v2/repositories/"+repo+"/", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var repoResp repositoryResponse
	if err := json.NewDecoder(response.Body).Decode(&repoResp); err != nil {
		return nil, err
	}

	return &repoResp, nil
}

package services

import (
	"net/http"
	"sync"
	"work/src/api/config"
	"work/src/api/domain/github"
	"work/src/api/domain/repositories"
	"work/src/api/log/option_b"
	"work/src/api/providers/github_provider"
	"work/src/api/utils/errors"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(clientID string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(clientID string, requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(clientID string, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	option_b.Info("about to send request to external api",
		option_b.Field("client_id", clientID),
		option_b.Field("status", "pending"),
		option_b.Field("authenticated", clientID != ""))

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		option_b.Error("response obtained from external api",
			option_b.Field("client_id", clientID),
			option_b.Field("status", "error"),
			option_b.Field("authenticated", clientID != ""))
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	option_b.Info("response obtained from external api",
		option_b.Field("client_id", clientID),
		option_b.Field("status", "access"),
		option_b.Field("authenticated", clientID != ""))

	result := repositories.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}

func (s *reposService) CreateRepos(clientID string, requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResult(&wg, input, output)

	for _, current := range requests {
		wg.Add(1)
		go s.createRepoConcurrent(clientID, current, input)
	}

	wg.Wait()
	close(input)

	result := <-output

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}
	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}
	return result, nil
}

func (s *reposService) handleRepoResult(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	output <- results
}

func (s *reposService) createRepoConcurrent(clientID string, input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	result, err := s.CreateRepo(clientID, input)
	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	output <- repositories.CreateRepositoriesResult{Response: result}
}

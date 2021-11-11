package retrieve

import (
	"fmt"
	"github.com/google/go-github/v39/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"io/ioutil"
)

type Kind int64

const (
	Patches Kind = iota
	Commits
	Comments
)

type result struct {
	result string
	err    error
}

func Retrieve(org, project string, pull int, token string, kind Kind) ([]byte, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	retrievalResult := result{
		result: "",
		err:    nil,
	}

	//_, _ := regexp.Compile(`.*(:coin)\s*:\s*.*`)

	switch kind {
	case Patches:
		patches, _, err := getPatches(org, project, pull, client, ctx)
		retrievalResult.result = patches
		retrievalResult.err = err
		break
	case Commits:
		commits, err := getCommits(org, project, pull, client, ctx)
		retrievalResult.result = commits
		retrievalResult.err = err
		break
		//case Comments:
		//	comments, _, err := getComments(org, project, pull, client, ctx)
		//	retrievalResult.result = comments
		//	retrievalResult.err = err
		//	break;
	}

	if retrievalResult.err != nil {
		return []byte(""), fmt.Errorf("error getting patches %v", retrievalResult.err)
	}

	return []byte(retrievalResult.result), nil
}

func getPatches(org string, project string, pull int, client *github.Client, ctx context.Context) (string, *github.Response, error) {
	patches, response, err := client.PullRequests.GetRaw(ctx, org, project, pull, github.RawOptions{
		Type: github.Patch,
	})
	return patches, response, err
}

func getCommits(org string, project string, pull int, client *github.Client, ctx context.Context) (string, error) {
	pr, _, err := client.PullRequests.Get(ctx, org, project, pull)
	if err != nil {
		return "", fmt.Errorf("error retrieving commits url: %v", err)
	}
	prUrl := pr.GetCommitsURL()
	commitsResponse, err := client.Client().Get(prUrl)

	if err != nil {
		return "", fmt.Errorf("error reading commit url: %v", err)
	}

	b, err := ioutil.ReadAll(commitsResponse.Body)
	defer commitsResponse.Body.Close()
	if err != nil {
		return "", fmt.Errorf("error reading commit url response body %v", err)
	}

	return string(b), err
}

//todo this needs to be implemented some time, and the current version in `get_review_weights` needs to be replaced
//func getReviews(org string, project string, pull int, client *github.Client, ctx context.Context) (string, *github.Response, error) {
//	patches, response, err := client.PullRequests.GetRaw(ctx, org, project, pull, github.RawOptions{
//		Type: github.Patch,
//	})
//	return patches, response, err
//}

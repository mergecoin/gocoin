package functions

import (
	"fmt"
	"github.com/google/go-github/v39/github"
	"github.com/ventureharbour/gocoin/commitscanner"
	"github.com/ventureharbour/gocoin/diffscanner/diffstream"
	"github.com/ventureharbour/gocoin/mint_scorer/lines"
	"github.com/ventureharbour/gocoin/mint_scorer/preambles"
	"github.com/ventureharbour/gocoin/retrieve"
	"encoding/json"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type CommitWeight struct {
	Name string
	Weight float64
}

type CommitWeights struct {
	Weights []CommitWeight
}

func determineCommitWeight(element commitscanner.CommitShard, token, org, repo string) float64 {
	//TODO do something mor complex with these commits
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	s, _, _ := client.Repositories.GetCommitRaw(ctx, org, repo, element.Sha, github.RawOptions{
		Type: github.Patch,
	});

	stream := diffstream.NewDiffStream([]byte(s))

	stream.InitializeData()

	return stream.GenerateScore(&lines.UnimplementedLineScorerExample{}, &preambles.UnimplementedPreambleScorerExample{})
}

func CalculateCommitWeights(org, project, token string, pull int) (string, error) {
	commits := commitscanner.Commits{}
	jsonVal, err := retrieve.Retrieve(org, project, pull, token, retrieve.Commits)

	if err != nil {
		return "", fmt.Errorf("error retrieving commits %v", err)
	}

	jsonString := string(jsonVal)
	err = commits.FromJson(jsonString)
	if err != nil {
		return "", fmt.Errorf("cannot unmarshal commits json %v", err)
	}

	cMap := make(map[string]float64)

	for _, element := range commits.Pool {
		value := determineCommitWeight(element, token, org, project)
		cMap[element.Author.Login] += value
	}

	sum := 0.0
	for _, element := range cMap {
		sum += element
	}

	cMap2 := make(map[string]float64)

	for key, element := range cMap {
		cMap2[key] = element/sum * 100
	}

	v, err := json.Marshal(cMap2)

	if err != nil {
		return "", fmt.Errorf("unable to marshal weight json %v", err)
	}

	return string(v), nil
}

package functions

import (
	"encoding/json"
	"fmt"
	"math"
)

type Determinator struct {
	Awards map[string]int64 `json:"awards"`
}

type Split struct {
	Review     uint8 `json:"review,string"`
	Contribute uint8 `json:"contribute,string"`
}

type IgnoreFiles struct {
	Names []string `json:"ignored"`
}

type DeterminationConfig struct {
	Split Split `json:"split"`
	Ignored IgnoreFiles `json:"ignored"`
}

func Sample(args ...string) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}


// Get determination!!
func Determine(org, project, token string, pull int, age uint, configuration []byte) (determination Determinator, err error) {
	determination = Determinator{
		Awards: make(map[string]int64),
	}

	config := DeterminationConfig{
		Split: Split{
			Review:     25,
			Contribute: 75,
		},
		Ignored: IgnoreFiles{
			Names: []string{"package-lock.json"},
		},
	}

	err = json.Unmarshal(configuration, &config)

	if err != nil {
		return determination, fmt.Errorf("unable to unmarshal config options %v", err)
	}

	reviewWeight, err := CalculateReviewAndCommentWeight(org, project, token, pull)
	if err != nil {
		return determination, fmt.Errorf("unable to retrieve review weights %v", err)
	}

	commitWeight, err := CalculateCommitWeights(org, project, token, pull)
	if err != nil {
		return determination, fmt.Errorf("unable to retrieve commit weights %v", err)
	}

	value, err := DeterminePullRequestWorth(org, project, token, pull, age)
	if err != nil {
		return determination, fmt.Errorf("unable to retrieve pr worth %v", err)
	}

	amountToContributors := value * (float64(config.Split.Contribute) / 100.0)
	amountToReviewers := value * (float64(config.Split.Review) / 100.0)

	for contributor, weight := range commitWeight {
		award, ok := determination.Awards[contributor]
		if ok {
			newValue := float64(award) + amountToContributors*(weight/100.0)
			determination.Awards[contributor] = int64(math.Round(newValue))
		} else {
			determination.Awards[contributor] = int64(math.Round(amountToContributors * (weight / 100.0)))
		}
	}

	for reviewer, weight := range reviewWeight {
		award, ok := determination.Awards[reviewer]
		if ok {
			newValue := float64(award) + amountToReviewers*(weight/100.0)
			determination.Awards[reviewer] = int64(math.Round(newValue)) + 1
		} else {
			determination.Awards[reviewer] = int64(math.Round(amountToReviewers*(weight/100.0))) + 1
		}
	}

	return determination, err
}

package functions

import (
	"fmt"
	"math"

	"github.com/ventureharbour/gocoin/diffscanner/diffstream"
	"github.com/ventureharbour/gocoin/mint_scorer/lines"
	"github.com/ventureharbour/gocoin/mint_scorer/preambles"
	"github.com/ventureharbour/gocoin/retrieve"
)

// Determines that dropped off value of a PR based on the weight of the changeset
// this is sort of like the number of changes,
// but with certain lines weighted up or down based on some criteria
func dropoff(changes float64) float64 {
	desired := 250.0
	if changes >= desired {
		val := math.Round(37.56 * math.Exp(-0.0005*changes))
		if val == 0 {
			return 1
		} else {
			return val
		}
	}
	val := math.Round(4.02 * math.Exp(0.01*changes))
	if val == 0 {
		return 1
	} else {
		return val
	}
}

// returns a
func ageOfPr(age uint, currentValue float64) float64 {
	val := currentValue * (1.08 * math.Exp(-0.08*float64(age)))
	if val < 1 {
		return 1
	}
	return val
}

// Determines an amount of mergecoin for a given PR
func DeterminePullRequestWorth(org, project, token string, pull int, age uint) (float64, error) {
	retrieved, err := retrieve.Retrieve(
		org,
		project,
		pull,
		token, retrieve.Patches)

	if err != nil {
		return 0.0, fmt.Errorf("retrieve commits %v", err)
	}

	stream := diffstream.NewDiffStream(retrieved)

	stream.InitializeData()

	changeWeights := stream.GenerateScore(&lines.BasicLineScorer{}, &preambles.UnimplementedPreambleScorerExample{})

	totalValue := dropoff((changeWeights))

	return ageOfPr(age, totalValue), nil
}

package functions

import (
	"fmt"
	"github.com/ventureharbour/gocoin/diffscanner/diffstream"
	"github.com/ventureharbour/gocoin/retrieve"
	"github.com/ventureharbour/gocoin/mint_scorer/lines"
	"github.com/ventureharbour/gocoin/mint_scorer/preambles"
)

// Determines an amount of mergecoin for a given PR
func DeterminePullRequestWorth(org, project, token string, pull int) (float64, error) {
	retrieved, err := retrieve.Retrieve(
		org,
		project,
		pull,
		token, retrieve.Patches);

	if err != nil {
		return 0.0, fmt.Errorf("retrieve commits %v", err)
	}

	stream := diffstream.NewDiffStream(retrieved)

	stream.InitializeData()

	//for _, x := range stream.Info.Data {
	//	fmt.Printf("--------- %s ---------", x.Name)
	//	fmt.Print("\n")
	//	fmt.Print("edits ", x.Edits, "\n")
	//	fmt.Print("new ", x.Additions, "\n")
	//	fmt.Print("deletions ", x.Deletes, "\n")
	//	fmt.Print("frags ", x.Fragments, "\n")
	//	fmt.Print("extension ", x.Extension, "\n")
	//	fmt.Print("----FRAGMENTS----\n")
	//	for _, y := range x.Fragments {
	//		fmt.Printf("%s", y.Lines, "\n")
	//	}
	//	fmt.Print("\n\n")
	//}
	//
	//fmt.Println("\n-----PREAMBLE-----\n %s", stream.Info.Preamble)
	//fmt.Println("\n----TOTAL SCORE FOR THIS PULL REQUEST----\n", stream.GenerateScore(&lines.UnimplementedLineScorerExample{}, &preambles.UnimplementedPreambleScorerExample{}))

	return stream.GenerateScore(&lines.UnimplementedLineScorerExample{}, &preambles.UnimplementedPreambleScorerExample{}), nil
}

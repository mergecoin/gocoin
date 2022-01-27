package main

import (
	"fmt"

	"github.com/ventureharbour/gocoin/functions"
)

func main() {
	//{"org": "ventureharbour", "repo": "truenorth", "pull": "12", "token": "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"}
	//weight, err := functions.CalculateCommitWeights("ventureharbour", "truenorth", "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", 590)
	// weight, err := functions.DeterminePullRequestWorth("ventureharbour", "truenorth", "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", 545, 1)
	// weight, err := functions.DeterminePullRequestWorth("ventureharbour", "ventureharbour.com", "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", 43, 1)

	weight, err := functions.Determine("ventureharbour", "ventureharbour.com", "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", 43, 1, []byte(`{"split": {"contribute": "60","review": "40"}, "ignored": {"ignored": ["assets/html"]}}`))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("weight %v", weight)
	//stream := diffstream.NewDiffStream(retrieve.Retrieve(
	//	"ventureharbour",
	//	"mergecoin",
	//	14,
	//	"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"))
	//
	//stream.InitializeData()
	//
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

	//fmt.Println("\n-----PREAMBLE-----\n %s", stream.Info.Preamble)
	//fmt.Println("\n----TOTAL SCORE FOR THIS PULL REQUEST----\n", stream.GenerateScore(&lines.UnimplementedLineScorerExample{}, &preambles.UnimplementedPreambleScorerExample{}))
}

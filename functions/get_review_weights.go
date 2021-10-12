package functions

import (
	"fmt"
	"github.com/google/go-github/v39/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"strings"
	"time"
)

type Score struct {
	Value float64
}

func (w *Score) Increase(amount float64) {
	w.Value += amount
}

type Weighting struct {
	owner string
	score float64
}

func (w *Weighting) UpdateName(name string) {
	w.owner = name
}

func (w *Weighting) IncrementScore() {
	w.score += 1
}

func (w *Weighting) IncreaseScore(amount float64) {
	w.score += amount
}

func getCommentWeight(body string, reactions *github.Reactions) float64 {
	weight := 1
	positiveReactionCount := reactions.GetHeart() + reactions.GetHooray() + reactions.GetPlusOne() + reactions.GetRocket()
	negativeReactionCount := reactions.GetMinusOne() + reactions.GetConfused()

	weight += positiveReactionCount - negativeReactionCount

	suggestions := strings.Count(body, "```suggestion")
	weight += suggestions

	return float64(weight)
}

func CalculateReviewAndCommentWeight(org, project, token string, pull int) (map[string]float64, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	commentArgs := github.PullRequestListCommentsOptions{
		Sort:        "created",
		Direction:   "desc",
		Since:       time.Time{},
		ListOptions: github.ListOptions{},
	}

	comments, _, err := client.PullRequests.ListComments(ctx, org, project, pull, &commentArgs);

	if err != nil {
		fmt.Printf("error retreiving comments: %v", err)
		return nil, err
	}

	weight := make(map[int64]*Weighting)

	for _, comment := range comments {
		id := comment.GetID()
		commenter := comment.GetUser().GetLogin()
		replyTo := comment.GetInReplyTo()

		commentWeight := getCommentWeight(comment.GetBody(), comment.GetReactions())

		original, ok := weight[id]
		if !ok {
			weight[id] = &Weighting{owner: commenter, score: commentWeight}
		} else {
			original.IncreaseScore(commentWeight)
			if original.owner == "" {
				original.UpdateName(commenter)
			}
		}

		if replyTo != 0 {
			parent, ok := weight[replyTo]
			if ok {
				parent.IncrementScore()
			} else {
				weight[replyTo] = &Weighting{
					owner: "",
					score: 1,
				}
			}
		}
	}

	combined := make(map[string]*Score)
	for _, comment := range weight {
		c, ok := combined[comment.owner]
		if ok {
			c.Increase(comment.score)
		} else {
			combined[comment.owner] = &Score{
				Value: comment.score,
			}
		}
	}

	reviews, _, err := client.PullRequests.ListReviews(ctx, org, project, pull, nil);

	if err != nil {
		fmt.Printf("error retreiving reviews: %v", err)
		return nil, err
	}

	for _, review := range reviews {
		if review.GetState() == "APPROVED" {
			comment, ok := combined[review.GetUser().GetLogin()]
			if ok {
				comment.Increase(comment.Value * 2.0)
			} else {
				combined[review.GetUser().GetLogin()] = &Score{
					Value: 5.0,
				}
			}
		}
	}

	sum := 0.0
	for _, s := range combined {
		sum += s.Value
	}

	result := make(map[string]float64)
	for key, val := range combined {
		result[key] = (val.Value / sum) * 100.0
	}

	return result, nil
}

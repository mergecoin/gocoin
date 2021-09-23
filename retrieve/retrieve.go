package retrieve

import (
	"fmt"
	"github.com/google/go-github/v39/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func Retrieve(org, project string, pull int, token string) []byte {
	//client := &http.Client{}
	//
	//req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/repos/%v/%v/pulls/%v", org, project, pull), nil)
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//req.Header.Set("Accept", `application/vnd.github.v3.patch`)
	//req.Header.Set("Authorization", fmt.Sprintf(`token %v`, token))
	//
	//resp, err := client.Do(req)
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("https://api.github.com/repos/%v/%v/pulls/%v", org, project, pull)
	//fmt.Print(string(body))
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//return body



	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	patches, _, err := client.PullRequests.GetRaw(ctx, org, project, pull, github.RawOptions{
		Type: github.Patch,
	})

	if err != nil {
		fmt.Println("err fail on err check for patches")
		fmt.Errorf("error getting patches %v", err)
	}

	//b, err := io.ReadAll(res.Body)
	//// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Println(string(b))

	return []byte(patches)
}

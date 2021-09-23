package retrieve

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Retrieve(org, project string, pull int, token string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/repos/%v/%v/pulls/%v", org, project, pull), nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", `application/vnd.github.v3.patch`)
	req.Header.Set("Authorization", fmt.Sprintf(`token %v`, token))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}

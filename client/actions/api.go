package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func makeRequest(data map[string]interface{}, endpoint string) apiResponse {
	client := &http.Client{}

	rawData, _ := json.Marshal(data)
	url := fmt.Sprintf("http://localhost:1623%s", endpoint)

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(rawData))
	if err != nil {
		return apiResponse{}
	}
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func (a apiResponse) prettyPrint() {
	dst := &bytes.Buffer{}
	_ = json.Indent(dst, a, "", "  ")
	fmt.Print(dst.String())
}

type apiResponse []byte

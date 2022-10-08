package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	token        string
	organization string
)
var result map[string]interface{}

type Body struct {
	InviteId float64 `json:"invitee_id"`
}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error when loading .env file")
	}

	token = os.Getenv("GITHUB_TOKEN")
	organization = os.Getenv("GITHUB_ORGANIZATION")
}

func createClient(method, url string, payload io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Printf("Got error %s", err.Error())
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token "+token)

	return req
}

func getUserId(username string) float64 {
	client := &http.Client{}
	url := "https://api.github.com/users/"

	request := createClient("GET", url+username, nil)
	response, err := client.Do(request)

	// json, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.NewDecoder(response.Body).Decode(&result)
	id := result["id"].(float64)
	return id
}

func SendInvite(username string) any {
	client := &http.Client{}

	var userId = getUserId(username)
	url := "https://api.github.com/orgs/" + organization + "/invitations"

	body := Body{
		InviteId: userId,
	}

	reqbody, _ := json.Marshal(body)
	request := createClient("POST", url, bytes.NewBuffer(reqbody))
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		//Failed to read response.
		panic(err)
	}
	if response.StatusCode == 201 {
		fmt.Println("[Success]:", username)
	} else {
		fmt.Println("[Failed]:", username)
	}

	return result
}

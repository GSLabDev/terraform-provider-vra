package vra

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Token ... Authentication Token
type tokenJsonstruct struct {
	Expires time.Time `json:"expires"`
	ID      string    `json:"id"`
	Tenant  string    `json:"tenant"`
}

// GetToken ... Get Token from existing store or create a new one.
func GetToken(host, userName, password, tenant string) (string, error) {

	token, err := getTokenFromHost(host, userName, password, tenant)
	if err != nil {
		fmt.Printf("[Error] Cannot get Token : %s", err.Error())
		return "", err
	}
	return token, nil
}

func getTokenFromHost(host, userName, password, tenant string) (string, error) {
	var jsonStr = []byte(`{"username":"` + userName + `","password":"` + password + `","tenant":"` + tenant + `"}`)
	log.Println(bytes.NewBuffer(jsonStr))
	req, err := http.NewRequest("POST", "https:/"+"/"+host+"/identity/api/tokens/", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("[ERROR] Error while requesting Token ", err)
		return "", err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR] Error while requesting Token ", err)
		return "", err
	}
	if resp.StatusCode == 400 {
		log.Println("[ERROR] Error in connection. Check Username, Password and tenant name")
		return "", fmt.Errorf("[ERROR] Error in connection. Check Username, Password and tenant name")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("[Error]")
	}

	var tokenGostruct tokenJsonstruct

	if err = json.Unmarshal(body, &tokenGostruct); err != nil {
		log.Printf("[ERROR] Error while unmarshal %s", err)
		return "", fmt.Errorf("[ERROR] Error while unmarshal %s", err)
	}

	token := tokenGostruct.ID
	return token, nil
}

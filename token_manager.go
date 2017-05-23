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
type Token struct {
	Expires time.Time `json:"expires"`
	ID      string    `json:"id"`
	Tenant  string    `json:"tenant"`
}

// GetToken ... Get Token from existing store or create a new one.
func GetToken(host, userName, password, tenant string) (Token, error) {

	var token Token
	fileData, err := ioutil.ReadFile(tenant + "_" + userName + ".tra")
	if err != nil {
		log.Println("[Warn] Cannot read Token File" + err.Error())
		token, err = getTokenFromHost(host, userName, password, tenant)
		if err != nil {
			fmt.Printf("[Error] Cannot get Token : %s", err.Error())
			return Token{}, err
		}
	} else {
		json.Unmarshal(fileData, &token)
		log.Println(token.ID)
		currentTime := time.Now().Local()
		if token.Expires.After(currentTime) {
			log.Println("Token is valid")
		} else {
			token, err = getTokenFromHost(host, userName, password, tenant)
			if err != nil {
				fmt.Printf("[Error] Cannot get Token : %s", err.Error())
				return Token{}, err
			}
		}
	}

	return token, nil
}

func getTokenFromHost(host, userName, password, tenant string) (Token, error) {
	var jsonStr = []byte(`{"username":"` + userName + `","password":"` + password + `","tenant":"` + tenant + `"}`)

	req, err := http.NewRequest("POST", "https://"+host+"/identity/api/tokens/", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("[ERROR] Error while requesting Token ", err)
		return Token{}, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR] Error while requesting Token ", err)
		return Token{}, err
	}
	defer resp.Body.Close()
	var record Token
	fileData, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(fileData, &record); err != nil {
		return Token{}, err
	}
	log.Printf("[INFO] Token details: Token ID: %s\n Token Expiry %s", record.ID, record.Expires.Format("2006-01-02-15:04:05"))

	err = ioutil.WriteFile(tenant+"_"+userName+".tra", fileData, 0644)
	return record, nil
}

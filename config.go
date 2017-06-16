package vra

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Config ... configuration for VRA
type Config struct {
	Host     string
	Username string
	Password string
	Tenant   string
}

// GetResponse ... get Response according to
func (c *Config) GetResponse(request *http.Request) ([]byte, error) {

	token, err := GetToken(c.Host, c.Username, c.Password, c.Tenant)
	if err != nil {
		log.Println("[ERROR] Error in getting token")
		return nil, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	var tempURL *url.URL
	tempURL, err = url.Parse("https:/" + "/" + c.Host + "/" + request.URL.Path)
	if err != nil {
		log.Println("[Error] URL is not in correct format")
		return nil, err
	}
	request.URL = tempURL
	tokenString := "Bearer " + token
	log.Println(request.URL)
	log.Println(tokenString)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", tokenString)

	client := &http.Client{Transport: tr}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(" [ERROR] Do: ", err)
		return nil, err
	
	}
	return ioutil.ReadAll(resp.Body)
}

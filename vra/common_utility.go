package vra

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func getBlueprintList(config Config) ([]byte, error) {
	url := "catalog-service/api/consumer/entitledCatalogItems/"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return nil, err
	}
	response, err := config.GetResponse(request)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
		return nil, err
	}
	return response, nil
}

func checkrequestStatus(d *schema.ResourceData, config Config, requestId string, timeOut int) error {
	timeout := time.After(time.Duration(timeOut) * time.Second)
	for {
		select {
		case <-time.After(1 * time.Second):
			state, err := getRequestResponse(config, requestId)
			if err == nil {
				if state == "SUCCESSFUL" {
					log.Println("[DEBUG] Blueprint executed SUCCESSFULLY")
					d.SetId(requestId)
					return nil
				} else if state == "SUBMITTED" || state == "IN_PROGRESS" {
				} else if state == "FAILED" {
					log.Println("[ERROR] Failed")
					return fmt.Errorf("[Error] Failed execution")
				}
			} else {
				return err
			}
		case <-timeout:
			fmt.Printf("timeout\n")
			return fmt.Errorf("[ERROR] Timeout")
		}
	}
}

func getRequestResponse(config Config, requestId string) (string, error) {

	url := "catalog-service/api/consumer/requests/" + requestId
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return "", fmt.Errorf("[ERROR] Error in creating http Request %s", err)
	}
	response, err := config.GetResponse(request)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
		return "", fmt.Errorf("[ERROR] Error in getting response %s", err)
	}
	var requestGostruct requestJsonstruct
	if err = json.Unmarshal(response, &requestGostruct); err != nil {
		log.Printf("[ERROR] Error while unmarshal %s", err)
		return "", fmt.Errorf("[ERROR] Error while unmarshal %s", err)
	}
	return requestGostruct.State, nil
}

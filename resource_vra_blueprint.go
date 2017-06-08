package vra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

type Info struct {
	Content []struct {
		_type       string `json:"@type"`
		CatalogItem struct {
			Callbacks          interface{} `json:"callbacks"`
			CatalogItemTypeRef struct {
				ID    string `json:"id"`
				Label string `json:"label"`
			} `json:"catalogItemTypeRef"`
			DateCreated string `json:"dateCreated"`
			Description string `json:"description"`
			Forms       struct {
				CatalogRequestInfoHidden bool `json:"catalogRequestInfoHidden"`
				ItemDetails              struct {
					FormID string `json:"formId"`
					Type   string `json:"type"`
				} `json:"itemDetails"`
				RequestDetails struct {
					FormID string `json:"formId"`
					Type   string `json:"type"`
				} `json:"requestDetails"`
				RequestFormScale  string `json:"requestFormScale"`
				RequestSubmission struct {
					FormID string `json:"formId"`
					Type   string `json:"type"`
				} `json:"requestSubmission"`
			} `json:"forms"`
			IconID          string `json:"iconId"`
			ID              string `json:"id"`
			IsNoteworthy    bool   `json:"isNoteworthy"`
			LastUpdatedDate string `json:"lastUpdatedDate"`
			Name            string `json:"name"`
			Organization    struct {
				SubtenantLabel interface{} `json:"subtenantLabel"`
				SubtenantRef   interface{} `json:"subtenantRef"`
				TenantLabel    string      `json:"tenantLabel"`
				TenantRef      string      `json:"tenantRef"`
			} `json:"organization"`
			ProviderBinding struct {
				BindingID   string `json:"bindingId"`
				ProviderRef struct {
					ID    string `json:"id"`
					Label string `json:"label"`
				} `json:"providerRef"`
			} `json:"providerBinding"`
			Quota       int  `json:"quota"`
			Requestable bool `json:"requestable"`
			ServiceRef  struct {
				ID    string `json:"id"`
				Label string `json:"label"`
			} `json:"serviceRef"`
			Status     string `json:"status"`
			StatusName string `json:"statusName"`
			Version    int    `json:"version"`
		} `json:"catalogItem"`
		EntitledOrganizations []struct {
			SubtenantLabel string `json:"subtenantLabel"`
			SubtenantRef   string `json:"subtenantRef"`
			TenantLabel    string `json:"tenantLabel"`
			TenantRef      string `json:"tenantRef"`
		} `json:"entitledOrganizations"`
	} `json:"content"`
	Links    []interface{} `json:"links"`
	Metadata struct {
		Number        int `json:"number"`
		Offset        int `json:"offset"`
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
	} `json:"metadata"`
}

type Template struct {
	BusinessGroupID string      `json:"businessGroupId"`
	CatalogItemID   string      `json:"catalogItemId"`
	Data            interface{} `json:"data"`
	Description     interface{} `json:"description"`
	Reasons         interface{} `json:"reasons"`
	RequestedFor    string      `json:"requestedFor"`
	Type            string      `json:"type"`
}

type RequestOutput struct {
	_type                      string `json:"@type"`
	ApprovalStatus             string `json:"approvalStatus"`
	CatalogItemProviderBinding struct {
		BindingID   string `json:"bindingId"`
		ProviderRef struct {
			ID    string `json:"id"`
			Label string `json:"label"`
		} `json:"providerRef"`
	} `json:"catalogItemProviderBinding"`
	CatalogItemRef struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	} `json:"catalogItemRef"`
	Components      interface{} `json:"components"`
	DateApproved    interface{} `json:"dateApproved"`
	DateCompleted   interface{} `json:"dateCompleted"`
	DateCreated     string      `json:"dateCreated"`
	DateSubmitted   string      `json:"dateSubmitted"`
	Description     interface{} `json:"description"`
	ExecutionStatus string      `json:"executionStatus"`
	IconID          string      `json:"iconId"`
	ID              string      `json:"id"`
	LastUpdated     string      `json:"lastUpdated"`
	Organization    struct {
		SubtenantLabel interface{} `json:"subtenantLabel"`
		SubtenantRef   string      `json:"subtenantRef"`
		TenantLabel    interface{} `json:"tenantLabel"`
		TenantRef      string      `json:"tenantRef"`
	} `json:"organization"`
	Phase             string      `json:"phase"`
	PostApprovalID    interface{} `json:"postApprovalId"`
	PreApprovalID     interface{} `json:"preApprovalId"`
	Quote             struct{}    `json:"quote"`
	Reasons           interface{} `json:"reasons"`
	RequestCompletion interface{} `json:"requestCompletion"`
	RequestData       struct {
		Entries []struct {
			Key   string `json:"key"`
			Value struct {
				ClassID     string      `json:"classId"`
				ComponentID interface{} `json:"componentId"`
				ID          string      `json:"id"`
				Label       string      `json:"label"`
				Type        string      `json:"type"`
			} `json:"value"`
		} `json:"entries"`
	} `json:"requestData"`
	RequestNumber            interface{} `json:"requestNumber"`
	RequestedBy              string      `json:"requestedBy"`
	RequestedFor             string      `json:"requestedFor"`
	RequestedItemDescription string      `json:"requestedItemDescription"`
	RequestedItemName        string      `json:"requestedItemName"`
	RequestorEntitlementID   string      `json:"requestorEntitlementId"`
	RetriesRemaining         int         `json:"retriesRemaining"`
	State                    string      `json:"state"`
	StateName                interface{} `json:"stateName"`
	Version                  int         `json:"version"`
	WaitingStatus            string      `json:"waitingStatus"`
}

func ExecuteBlueprint() *schema.Resource {
	return &schema.Resource{
		Create: ExecuteBlueprintC,
		Read:   ExecuteBlueprintR,
		Delete: ExecuteBlueprintD,
		Schema: map[string]*schema.Schema{
			"blueprint_name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateBlueprintName,
			},
			"file_name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateFileName,
			},
		},
	}
}

func ExecuteBlueprintC(d *schema.ResourceData, meta interface{}) error {

	config := meta.(Config)
	blueprintName := d.Get("blueprint_name").(string)
	filename := d.Get("file_name").(string)
	url := "catalog-service/api/consumer/entitledCatalogItems/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return fmt.Errorf("[ERROR] Error in creating http Request %s", err)
	}
	resp, err := config.GetResponse(req)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
		return fmt.Errorf("[ERROR] Error in getting response %s", err)
	}
	var record Info
	if err = json.Unmarshal(resp, &record); err != nil {
		log.Printf("[ERROR] Error while unmarshal catlogitems json", err)
	}
	i := record.Metadata.TotalElements
	var id = ""
	for j := 0; j < i; j++ {
		if record.Content[j].CatalogItem.Name == blueprintName {
			id = record.Content[j].CatalogItem.ID
		}
	}
	log.Println(id)
	if id == "" {
		log.Printf("[Error] Blueprint is not present")
		return fmt.Errorf("[ERROR]Blueprint is not present %s", blueprintName)
	}
	url = "catalog-service/api/consumer/entitledCatalogItems/" + id + "/requests/template"
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return fmt.Errorf("[ERROR] Error in creating http Request %s", err)
	}
	resp, err = config.GetResponse(req)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
		return fmt.Errorf("[ERROR] Error in getting response %s", err)
	}
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		return fmt.Errorf("[ERROR] file error %s", e)
		//os.Exit(1)
	}

	//structure for user's json data
	var users_data interface{}

	var record1 Template
	if err = json.Unmarshal(resp, &record1); err != nil {
		log.Printf("[ERROR] Error while unmarshal template json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal template json  %s", err)
	}

	err = json.Unmarshal(file, &users_data)
	if err != nil {
		log.Printf("[ERROR] Error while unmarshal file's json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal file's json %s", err)
	}
	record1.Data = users_data
	buff, _ := json.Marshal(&record1)
	var jsonStr = []byte(buff)

	url = "catalog-service/api/consumer/entitledCatalogItems/" + id + "/requests"
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return fmt.Errorf("[ERROR] Error in creating http Request %s", err)
	}
	resp, err = config.GetResponse(req)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
		return fmt.Errorf("[ERROR] Error in getting response %s", err)
	}

	var record2 RequestOutput
	if err = json.Unmarshal(resp, &record2); err != nil {
		log.Printf("[ERROR] Error while unmarshal requests json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal requests json %s", err)
	}
	reqid := record2.ID

	for {
		if record2.State == "SUCCESSFUL" {
			log.Println("[INFO] Blueprint executed SUCCESSFULLY")
			d.SetId(reqid)
			break
		} else if record2.State == "SUBMITTED" {
			time.Sleep(time.Second * 15)
			url = "catalog-service/api/consumer/requests/" + reqid
			req, err = http.NewRequest("GET", url, nil)
			if err != nil {
				log.Printf("[ERROR] Error in creating http Request %s", err)
				return fmt.Errorf("[ERROR] Error in creating http Request %s", err)
			}
			resp, err = config.GetResponse(req)
			if err != nil {
				log.Printf("[ERROR] Error in getting response %s", err)
				return fmt.Errorf("[ERROR] Error in getting response %s", err)
			}
			if err = json.Unmarshal(resp, &record2); err != nil {
				log.Printf("[ERROR] Error while unmarshal %s", err)
				return fmt.Errorf("[ERROR] Error while unmarshal %s", err)
			}
			//log.Println("[INFO] in submitted state" + record2.State)
		} else if record2.State == "IN_PROGRESS" {
			time.Sleep(time.Second * 15)
			url = "catalog-service/api/consumer/requests/" + reqid
			req, err = http.NewRequest("GET", url, nil)
			if err != nil {
				log.Printf("[ERROR] Error in creating http Request %s", err)
				return fmt.Errorf("[ERROR] Error in creating http Request %s", err)
			}
			resp, err = config.GetResponse(req)
			if err != nil {
				log.Printf("[ERROR] Error in getting response %s", err)
				return fmt.Errorf("[ERROR] Error in getting response %s", err)
			}
			if err = json.Unmarshal(resp, &record2); err != nil {
				log.Println(err)
			}
			//log.Println("[INFO] in progress state" + record2.State)
		} else if record2.State == "FAILED" {
			log.Println("[ERROR] Failed")
			break
		}
	}
	return nil
}

func ExecuteBlueprintR(d *schema.ResourceData, metadata interface{}) error {
	return nil
}

func ExecuteBlueprintD(d *schema.ResourceData, metadata interface{}) error {
	return nil

}

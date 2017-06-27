package vra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

func ExecuteBlueprint() *schema.Resource {
	return &schema.Resource{
		Create: ExecuteBlueprintCreate,
		Read:   ExecuteBlueprintRead,
		Delete: ExecuteBlueprintDelete,
		Schema: map[string]*schema.Schema{
			"blueprint_name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateBlueprintName,
			},
			"input_file_name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateFileName,
			},
			"time_out": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func ExecuteBlueprintCreate(d *schema.ResourceData, meta interface{}) error {

	config := meta.(Config)
	blueprintName := d.Get("blueprint_name").(string)
	inputFileName := d.Get("input_file_name").(string)
	timeout := d.Get("time_out").(int)
	response, err := getBlueprintList(config)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
		return fmt.Errorf("[ERROR] Error in getting response %s", err)
	}
	var blueprintlistGostruct blueprintlistJsonstruct

	if err = json.Unmarshal(response, &blueprintlistGostruct); err != nil {
		log.Printf("[ERROR] Error while unmarshal requests json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal requests json %s", err)
	}

	i := blueprintlistGostruct.Metadata.TotalElements
	var blueprintId = ""
	for j := 0; j < i; j++ {
		if blueprintlistGostruct.Content[j].CatalogItem.Name == blueprintName {
			blueprintId = blueprintlistGostruct.Content[j].CatalogItem.ID
		}
	}
	if blueprintId == "" {
		log.Printf("[Error] Blueprint is not present")
		return fmt.Errorf("[ERROR]Blueprint is not present %s", blueprintName)
	}

	url := "catalog-service/api/consumer/entitledCatalogItems/" + blueprintId + "/requests/template"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
	}
	response, err = config.GetResponse(request)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
	}
	file, e := ioutil.ReadFile(inputFileName)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
	}

	//structure for user's json data
	var usersInputJson interface{}

	var templateGostruct templateJsonstruct
	if err := json.Unmarshal(response, &templateGostruct); err != nil {
		log.Printf("[ERROR] Error while unmarshal template json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal template json  %s", err)
	}

	err = json.Unmarshal(file, &usersInputJson)
	if err != nil {
		log.Printf("[ERROR] Error while unmarshal file's json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal file's json %s", err)
	}
	templateGostruct.Data = usersInputJson
	buff, _ := json.Marshal(&templateGostruct)
	var jsonStr = []byte(buff)

	url = "catalog-service/api/consumer/entitledCatalogItems/" + blueprintId + "/requests"
	request, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Printf("[ERROR] Error in creating http Request %s", err)
		return fmt.Errorf("[ERROR] Error in creating http Request %s", err)
	}
	response, err = config.GetResponse(request)
	if err != nil {
		log.Printf("[ERROR] Error in getting response %s", err)
		return fmt.Errorf("[ERROR] Error in getting response %s", err)
	}
	var requestGostruct requestJsonstruct
	if err = json.Unmarshal(response, &requestGostruct); err != nil {
		log.Printf("[ERROR] Error while unmarshal requests json %s", err)
		return fmt.Errorf("[ERROR] Error while unmarshal requests json %s", err)
	}
	requestId := requestGostruct.ID
	if timeout == 0 {
		return checkrequestStatus(d, config, requestId, 50)
	} else {
		return checkrequestStatus(d, config, requestId, timeout)
	}
}

func ExecuteBlueprintRead(d *schema.ResourceData, metadata interface{}) error {
	return nil
}

func ExecuteBlueprintDelete(d *schema.ResourceData, metadata interface{}) error {
	return nil

}

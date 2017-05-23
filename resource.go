package vra

import (
	"bytes"
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

type MyJsonName struct {
	Content []struct {
		CatalogItem struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"catalogItem"`
	} `json:"content"`
	Metadata struct {
		Number        int `json:"number"`
		Offset        int `json:"offset"`
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
	} `json:"metadata"`
}

func ExecuteBlueprint() *schema.Resource {
	return &schema.Resource{
		Create: ExecuteBlueprintC,
		Read:   ExecuteBlueprintR,
		Delete: ExecuteBlueprintD,
		Schema: map[string]*schema.Schema{
			"blueprint_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func ExecuteBlueprintC(d *schema.ResourceData, metadata interface{}) error {

	config := metadata.(Config)
	blueprintName := d.Get("blueprint_name").(string)

	url := "/catalog-service/api/consumer/entitledCatalogItems/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("[ERROR] Error in creating http Request")
		return err
	}
	resp, err := config.GetResponse(req)
	if err != nil {
		log.Fatal(err)
	}
	var record1 MyJsonName
	if err := json.Unmarshal(resp, &record1); err != nil {
		log.Println(err)
	}
	i := record1.Metadata.TotalElements
	var id = ""
	for j := 0; j < i; j++ {
		if record1.Content[j].CatalogItem.Name == blueprintName {
			id = record1.Content[j].CatalogItem.ID
		}
	}

	url1 := "/catalog-service/api/consumer/entitledCatalogItems/" + id + "/requests/template"
	req, err = http.NewRequest("GET", url1, nil)
	resp, err = config.GetResponse(req)
	if err != nil {
		log.Fatal(err)
	}

	url2 := "/catalog-service/api/consumer/entitledCatalogItems/" + id + "/requests"
	req, err = http.NewRequest("POST", url2, bytes.NewBuffer(resp))
	_, err = config.GetResponse(req)
	if err != nil {
		log.Println("")
		return fmt.Errorf("[ERROR]Request failed\n%s", err)
	}

	return nil
}

func ExecuteBlueprintR(d *schema.ResourceData, metadata interface{}) error {
	return nil
}

func ExecuteBlueprintD(d *schema.ResourceData, metadata interface{}) error {
	return nil

}
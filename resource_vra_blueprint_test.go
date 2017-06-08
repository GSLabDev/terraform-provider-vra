//make testacc TEST=./builtin/providers/vra TESTARGS="-run="
package vra

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"log"
	"net/http"
	"os"
	"testing"
)

type Resources struct {
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

func testBasicPreCheckVM(t *testing.T) {

	testAccPreCheck(t)

	if v := os.Getenv("VRA_URL"); v == "" {
		t.Fatal("VRA URL must be set for acceptance tests")
	}

	if v := os.Getenv("VRA_TENANT"); v == "" {
		t.Fatal("VRA TENANT must be set for acceptance tests")
	}

	if v := os.Getenv("VRA_USER"); v == "" {
		t.Fatal("VRA USER NAME must be set for acceptance tests")
	}

	if v := os.Getenv("VRA_PASSWORD"); v == "" {
		t.Fatal("VRA PASSWORD must be set for acceptance tests")
	}
}

func TestAccExecuteBlueprint_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBlueprintDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckBlueprintConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlueprintExists("vra_execute_blueprint.Exe_blueprint"),
				),
			},
		},
	})
}

func testAccCheckBlueprintDestroy(s *terraform.State) error {
	return nil
}

func testAccCheckBlueprintExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No request ID is set")
		}

		config := testAccProvider.Meta().(Config)
		url := "catalog-service/api/consumer/requests/" + rs.Primary.ID
		req, err := http.NewRequest("GET", url, nil)
		resp, err := config.GetResponse(req)
		if err != nil {
			log.Fatal(err)
		}
		var record Resources
		if err := json.Unmarshal(resp, &record); err != nil {
			log.Println(err)
		}
		if record.State != "SUCCESSFUL" {
			return fmt.Errorf("[ERROR] Blueprint is not executed")
		}
		return nil
	}
}

const testAccCheckBlueprintConfigBasic = `
resource "vra_execute_blueprint" "Exe_blueprint"{
     	blueprint_name = "Create simple virtual machine"
        file_name = "data.json"    
}`

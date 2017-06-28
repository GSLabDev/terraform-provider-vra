package vra

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"os"
	"testing"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"vra": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}
func testAccPreCheck(t *testing.T) {
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

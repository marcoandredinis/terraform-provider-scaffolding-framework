package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExampleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		PreCheck:                 func() {},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccExampleResourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("scaffolding_example.test", "metadata.name", "abc"),
					resource.TestCheckResourceAttr("scaffolding_example.test2", "metadata.name", "abc"),
				),
			},
		},
	})
}

func testAccExampleResourceConfig() string {
	return `
resource "scaffolding_example" "test" {
  metadata = {
	name = "abc"
  }
}
resource "scaffolding_example" "test2" {
	metadata = {
	  name = scaffolding_example.test.metadata.name
	}
  }
`
}

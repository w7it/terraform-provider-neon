package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDatabaseResourceDefault(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccDatabaseResourceConfigDefault("todo-app", "default"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr("neon_database.test", "id", idRegex()),
					resource.TestCheckResourceAttr("neon_database.test", "name", "todo-app"),
					resource.TestCheckResourceAttr("neon_database.test", "owner_name", "default"),
					resource.TestCheckResourceAttr("neon_database.test", "branch_id", "br-billowing-morning-664877"),
					resource.TestCheckResourceAttr("neon_database.test", "project_id", "rapid-shadow-433566"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "neon_database.test",
				ImportState:       true,
				ImportStateId:     "rapid-shadow-433566:br-billowing-morning-664877:todo-app",
				ImportStateVerify: true,
			},
			// Update with null values
			{
				Config: testAccDatabaseResourceConfigDefault("todo-app", "default"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr("neon_database.test", "id", idRegex()),
					resource.TestCheckResourceAttr("neon_database.test", "name", "todo-app"),
					resource.TestCheckResourceAttr("neon_database.test", "owner_name", "default"),
					resource.TestCheckResourceAttr("neon_database.test", "branch_id", "br-billowing-morning-664877"),
					resource.TestCheckResourceAttr("neon_database.test", "project_id", "rapid-shadow-433566"),
				),
			},
			// Update and Read testing
			{
				Config: testAccDatabaseResourceConfigDefault("nue-todo-app", "todo-app"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr("neon_database.test", "id", idRegex()),
					resource.TestCheckResourceAttr("neon_database.test", "name", "nue-todo-app"),
					resource.TestCheckResourceAttr("neon_database.test", "owner_name", "todo-app"),
					resource.TestCheckResourceAttr("neon_database.test", "branch_id", "br-billowing-morning-664877"),
					resource.TestCheckResourceAttr("neon_database.test", "project_id", "rapid-shadow-433566"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "neon_database.test",
				ImportState:       true,
				ImportStateId:     "rapid-shadow-433566:br-billowing-morning-664877:nue-todo-app",
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccDatabaseResourceConfigDefault(name string, owner string) string {
	return fmt.Sprintf(`
resource "neon_database" "test" {
  name = "%s"
  owner_name = "%s"
  branch_id = "br-billowing-morning-664877"
  project_id = "rapid-shadow-433566"
}
`, name, owner)
}

// scripts_state.go
package scripts

import (
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// updateTerraformState updates the Terraform state with the latest Script information from the Jamf Pro API.
func updateTerraformState(d *schema.ResourceData, resource *jamfpro.ResourceScript) diag.Diagnostics {
	var diags diag.Diagnostics

	// Update the Terraform state with the fetched data
	resourceData := map[string]interface{}{
		"id":              resource.ID,
		"name":            resource.Name,
		"category_name":   resource.CategoryName,
		"category_id":     resource.CategoryId,
		"info":            resource.Info,
		"notes":           resource.Notes,
		"os_requirements": resource.OSRequirements,
		"priority":        resource.Priority,
		"script_contents": resource.ScriptContents,
		"parameter4":      resource.Parameter4,
		"parameter5":      resource.Parameter5,
		"parameter6":      resource.Parameter6,
		"parameter7":      resource.Parameter7,
		"parameter8":      resource.Parameter8,
		"parameter9":      resource.Parameter9,
		"parameter10":     resource.Parameter10,
		"parameter11":     resource.Parameter11,
	}

	// Iterate over the map and set each key-value pair in the Terraform state
	for key, val := range resourceData {
		if err := d.Set(key, val); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags

}

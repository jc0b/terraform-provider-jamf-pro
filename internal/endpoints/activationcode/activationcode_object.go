// activationcode_object.go
package activationcode

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// constructJamfProActivationCode constructs a ResourceActivationCode object from the provided schema data and logs its XML representation.
func constructJamfProActivationCode(d *schema.ResourceData) (*jamfpro.ResourceActivationCode, error) {

	activationCode := &jamfpro.ResourceActivationCode{
		OrganizationName: d.Get("organization_name").(string),
		Code:             d.Get("code").(string),
	}

	// Serialize and pretty-print the ActivationCode object as XML for logging
	resourceXML, err := xml.MarshalIndent(activationCode, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Jamf Pro Activation Code to XML: %v", err)
	}

	log.Printf("[DEBUG] Constructed Jamf Pro Activation Code XML:\n%s\n", string(resourceXML))

	return activationCode, nil
}
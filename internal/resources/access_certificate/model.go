// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessCertificateResultEnvelope struct {
	Result AccessCertificateModel `json:"result,computed"`
}

type AccessCertificateModel struct {
	ID                  types.String    `tfsdk:"id" json:"id,computed"`
	AccountID           types.String    `tfsdk:"account_id" path:"account_id"`
	ZoneID              types.String    `tfsdk:"zone_id" path:"zone_id"`
	Certificate         types.String    `tfsdk:"certificate" json:"certificate"`
	Name                types.String    `tfsdk:"name" json:"name"`
	AssociatedHostnames *[]types.String `tfsdk:"associated_hostnames" json:"associated_hostnames"`
}
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r APIShieldOperationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"operation_id": schema.StringAttribute{
						Description: "UUID identifier",
						Required:    true,
					},
					"state": schema.StringAttribute{
						Description: "Mark state of operation in API Discovery\n  * `review` - Mark operation as for review\n  * `ignored` - Mark operation as ignored\n",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("review", "ignored"),
						},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state APIShieldOperationModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
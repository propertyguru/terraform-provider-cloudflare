// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_list

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &TeamsListDataSource{}

func (d *TeamsListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"list_id": schema.StringAttribute{
				Description: "API Resource UUID tag.",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"list_count": schema.Float64Attribute{
				Description: "The number of items in the list.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"description": schema.StringAttribute{
				Description: "The description of the list.",
				Computed:    true,
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "API Resource UUID tag.",
				Computed:    true,
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the list.",
				Computed:    true,
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of list.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("SERIAL", "URL", "DOMAIN", "EMAIL", "IP"),
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Required: true,
					},
					"type": schema.StringAttribute{
						Description: "The type of list.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("SERIAL", "URL", "DOMAIN", "EMAIL", "IP"),
						},
					},
				},
			},
		},
	}
}

func (d *TeamsListDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

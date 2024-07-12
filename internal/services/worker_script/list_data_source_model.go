// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_script

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkerScriptsResultListDataSourceEnvelope struct {
	Result *[]*WorkerScriptsItemsDataSourceModel `json:"result,computed"`
}

type WorkerScriptsDataSourceModel struct {
	AccountID types.String                          `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                           `tfsdk:"max_items"`
	Items     *[]*WorkerScriptsItemsDataSourceModel `tfsdk:"items"`
}

type WorkerScriptsItemsDataSourceModel struct {
	ID            types.String                                       `tfsdk:"id" json:"id,computed"`
	CreatedOn     timetypes.RFC3339                                  `tfsdk:"created_on" json:"created_on,computed"`
	Etag          types.String                                       `tfsdk:"etag" json:"etag,computed"`
	Logpush       types.Bool                                         `tfsdk:"logpush" json:"logpush"`
	ModifiedOn    timetypes.RFC3339                                  `tfsdk:"modified_on" json:"modified_on,computed"`
	PlacementMode types.String                                       `tfsdk:"placement_mode" json:"placement_mode"`
	TailConsumers *[]*WorkerScriptsItemsTailConsumersDataSourceModel `tfsdk:"tail_consumers" json:"tail_consumers"`
	UsageModel    types.String                                       `tfsdk:"usage_model" json:"usage_model"`
}

type WorkerScriptsItemsTailConsumersDataSourceModel struct {
	Service     types.String `tfsdk:"service" json:"service,computed"`
	Environment types.String `tfsdk:"environment" json:"environment"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace"`
}
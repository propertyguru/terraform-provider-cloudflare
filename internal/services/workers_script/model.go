// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"bytes"
	"mime/multipart"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apiform"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersScriptResultEnvelope struct {
	Result WorkersScriptModel `json:"result"`
}

type WorkersScriptModel struct {
	ID            types.String                                                  `tfsdk:"id" json:"-,computed"`
	ScriptName    types.String                                                  `tfsdk:"script_name" path:"script_name,required"`
	AccountID     types.String                                                  `tfsdk:"account_id" path:"account_id,required"`
	Message       types.String                                                  `tfsdk:"message" json:"message,optional"`
	AnyPartName   *[]types.String                                               `tfsdk:"any_part_name" json:"<any part name>,optional"`
	Metadata      *WorkersScriptMetadataModel                                   `tfsdk:"metadata" json:"metadata,optional"`
	CreatedOn     timetypes.RFC3339                                             `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Etag          types.String                                                  `tfsdk:"etag" json:"etag,computed"`
	Logpush       types.Bool                                                    `tfsdk:"logpush" json:"logpush,computed"`
	ModifiedOn    timetypes.RFC3339                                             `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	PlacementMode types.String                                                  `tfsdk:"placement_mode" json:"placement_mode,computed"`
	StartupTimeMs types.Int64                                                   `tfsdk:"startup_time_ms" json:"startup_time_ms,computed"`
	UsageModel    types.String                                                  `tfsdk:"usage_model" json:"usage_model,computed"`
	TailConsumers customfield.NestedObjectList[WorkersScriptTailConsumersModel] `tfsdk:"tail_consumers" json:"tail_consumers,computed"`
}

func (r WorkersScriptModel) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type WorkersScriptMetadataModel struct {
	Bindings           customfield.NestedObjectList[WorkersScriptMetadataBindingsModel]      `tfsdk:"bindings" json:"bindings,computed_optional"`
	BodyPart           types.String                                                          `tfsdk:"body_part" json:"body_part,computed_optional"`
	CompatibilityDate  types.String                                                          `tfsdk:"compatibility_date" json:"compatibility_date,computed_optional"`
	CompatibilityFlags customfield.List[types.String]                                        `tfsdk:"compatibility_flags" json:"compatibility_flags,computed_optional"`
	KeepBindings       customfield.List[types.String]                                        `tfsdk:"keep_bindings" json:"keep_bindings,computed_optional"`
	Logpush            types.Bool                                                            `tfsdk:"logpush" json:"logpush,computed_optional"`
	MainModule         types.String                                                          `tfsdk:"main_module" json:"main_module,computed_optional"`
	Migrations         customfield.NestedObject[WorkersScriptMetadataMigrationsModel]        `tfsdk:"migrations" json:"migrations,computed_optional"`
	Placement          customfield.NestedObject[WorkersScriptMetadataPlacementModel]         `tfsdk:"placement" json:"placement,computed_optional"`
	Tags               customfield.List[types.String]                                        `tfsdk:"tags" json:"tags,computed_optional"`
	TailConsumers      customfield.NestedObjectList[WorkersScriptMetadataTailConsumersModel] `tfsdk:"tail_consumers" json:"tail_consumers,computed_optional"`
	UsageModel         types.String                                                          `tfsdk:"usage_model" json:"usage_model,computed_optional"`
	VersionTags        customfield.Map[types.String]                                         `tfsdk:"version_tags" json:"version_tags,computed_optional"`
}

type WorkersScriptMetadataBindingsModel struct {
	Name types.String `tfsdk:"name" json:"name,computed_optional"`
	Type types.String `tfsdk:"type" json:"type,computed_optional"`
}

type WorkersScriptMetadataMigrationsModel struct {
	DeletedClasses     *[]types.String                                            `tfsdk:"deleted_classes" json:"deleted_classes,optional"`
	NewClasses         *[]types.String                                            `tfsdk:"new_classes" json:"new_classes,optional"`
	NewTag             types.String                                               `tfsdk:"new_tag" json:"new_tag,computed_optional"`
	OldTag             types.String                                               `tfsdk:"old_tag" json:"old_tag,computed_optional"`
	RenamedClasses     *[]*WorkersScriptMetadataMigrationsRenamedClassesModel     `tfsdk:"renamed_classes" json:"renamed_classes,optional"`
	TransferredClasses *[]*WorkersScriptMetadataMigrationsTransferredClassesModel `tfsdk:"transferred_classes" json:"transferred_classes,optional"`
	Steps              *[]*WorkersScriptMetadataMigrationsStepsModel              `tfsdk:"steps" json:"steps,optional"`
}

type WorkersScriptMetadataMigrationsRenamedClassesModel struct {
	From types.String `tfsdk:"from" json:"from,computed_optional"`
	To   types.String `tfsdk:"to" json:"to,computed_optional"`
}

type WorkersScriptMetadataMigrationsTransferredClassesModel struct {
	From       types.String `tfsdk:"from" json:"from,computed_optional"`
	FromScript types.String `tfsdk:"from_script" json:"from_script,computed_optional"`
	To         types.String `tfsdk:"to" json:"to,computed_optional"`
}

type WorkersScriptMetadataMigrationsStepsModel struct {
	DeletedClasses     customfield.List[types.String]                                                            `tfsdk:"deleted_classes" json:"deleted_classes,computed_optional"`
	NewClasses         customfield.List[types.String]                                                            `tfsdk:"new_classes" json:"new_classes,computed_optional"`
	RenamedClasses     customfield.NestedObjectList[WorkersScriptMetadataMigrationsStepsRenamedClassesModel]     `tfsdk:"renamed_classes" json:"renamed_classes,computed_optional"`
	TransferredClasses customfield.NestedObjectList[WorkersScriptMetadataMigrationsStepsTransferredClassesModel] `tfsdk:"transferred_classes" json:"transferred_classes,computed_optional"`
}

type WorkersScriptMetadataMigrationsStepsRenamedClassesModel struct {
	From types.String `tfsdk:"from" json:"from,computed_optional"`
	To   types.String `tfsdk:"to" json:"to,computed_optional"`
}

type WorkersScriptMetadataMigrationsStepsTransferredClassesModel struct {
	From       types.String `tfsdk:"from" json:"from,computed_optional"`
	FromScript types.String `tfsdk:"from_script" json:"from_script,computed_optional"`
	To         types.String `tfsdk:"to" json:"to,computed_optional"`
}

type WorkersScriptMetadataPlacementModel struct {
	Mode types.String `tfsdk:"mode" json:"mode,computed_optional"`
}

type WorkersScriptMetadataTailConsumersModel struct {
	Service     types.String `tfsdk:"service" json:"service,required"`
	Environment types.String `tfsdk:"environment" json:"environment,computed_optional"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace,computed_optional"`
}

type WorkersScriptTailConsumersModel struct {
	Service     types.String `tfsdk:"service" json:"service,computed"`
	Environment types.String `tfsdk:"environment" json:"environment,computed"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace,computed"`
}
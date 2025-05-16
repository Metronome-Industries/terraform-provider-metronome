// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1CustomerDataSourceModel struct {
	CustomerID types.String                                            `tfsdk:"customer_id" path:"customer_id,required"`
	Data       customfield.NestedObject[V1CustomerDataDataSourceModel] `tfsdk:"data" json:"data,computed"`
}

func (m *V1CustomerDataSourceModel) toReadParams(_ context.Context) (params metronome.V1CustomerGetParams, diags diag.Diagnostics) {
	params = metronome.V1CustomerGetParams{
		CustomerID: metronome.F(m.CustomerID.ValueString()),
	}

	return
}

type V1CustomerDataDataSourceModel struct {
	ID                    types.String                                                                 `tfsdk:"id" json:"id,computed"`
	CreatedAt             timetypes.RFC3339                                                            `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CustomFields          customfield.Map[types.String]                                                `tfsdk:"custom_fields" json:"custom_fields,computed"`
	CustomerConfig        customfield.NestedObject[V1CustomerDataCustomerConfigDataSourceModel]        `tfsdk:"customer_config" json:"customer_config,computed"`
	ExternalID            types.String                                                                 `tfsdk:"external_id" json:"external_id,computed"`
	IngestAliases         customfield.List[types.String]                                               `tfsdk:"ingest_aliases" json:"ingest_aliases,computed"`
	Name                  types.String                                                                 `tfsdk:"name" json:"name,computed"`
	ArchivedAt            timetypes.RFC3339                                                            `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CurrentBillableStatus customfield.NestedObject[V1CustomerDataCurrentBillableStatusDataSourceModel] `tfsdk:"current_billable_status" json:"current_billable_status,computed"`
}

type V1CustomerDataCustomerConfigDataSourceModel struct {
	SalesforceAccountID types.String `tfsdk:"salesforce_account_id" json:"salesforce_account_id,computed"`
}

type V1CustomerDataCurrentBillableStatusDataSourceModel struct {
	Value       types.String      `tfsdk:"value" json:"value,computed"`
	EffectiveAt timetypes.RFC3339 `tfsdk:"effective_at" json:"effective_at,computed" format:"date-time"`
}

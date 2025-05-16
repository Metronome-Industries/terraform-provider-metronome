// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V1CustomersDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[V1CustomersItemsDataSourceModel] `json:"data,computed"`
}

type V1CustomersDataSourceModel struct {
	IngestAlias          types.String                                                  `tfsdk:"ingest_alias" query:"ingest_alias,optional"`
	OnlyArchived         types.Bool                                                    `tfsdk:"only_archived" query:"only_archived,optional"`
	CustomerIDs          *[]types.String                                               `tfsdk:"customer_ids" query:"customer_ids,optional"`
	SalesforceAccountIDs *[]types.String                                               `tfsdk:"salesforce_account_ids" query:"salesforce_account_ids,optional"`
	MaxItems             types.Int64                                                   `tfsdk:"max_items"`
	Items                customfield.NestedObjectList[V1CustomersItemsDataSourceModel] `tfsdk:"items"`
}

func (m *V1CustomersDataSourceModel) toListParams(_ context.Context) (params metronome.V1CustomerListParams, diags diag.Diagnostics) {
	mCustomerIDs := []string{}
	for _, item := range *m.CustomerIDs {
		mCustomerIDs = append(mCustomerIDs, item.ValueString())
	}
	mSalesforceAccountIDs := []string{}
	for _, item := range *m.SalesforceAccountIDs {
		mSalesforceAccountIDs = append(mSalesforceAccountIDs, item.ValueString())
	}

	params = metronome.V1CustomerListParams{
		CustomerIDs:          metronome.F(mCustomerIDs),
		SalesforceAccountIDs: metronome.F(mSalesforceAccountIDs),
	}

	if !m.IngestAlias.IsNull() {
		params.IngestAlias = metronome.F(m.IngestAlias.ValueString())
	}
	if !m.OnlyArchived.IsNull() {
		params.OnlyArchived = metronome.F(m.OnlyArchived.ValueBool())
	}

	return
}

type V1CustomersItemsDataSourceModel struct {
	ID                    types.String                                                              `tfsdk:"id" json:"id,computed"`
	CreatedAt             timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CustomFields          customfield.Map[types.String]                                             `tfsdk:"custom_fields" json:"custom_fields,computed"`
	CustomerConfig        customfield.NestedObject[V1CustomersCustomerConfigDataSourceModel]        `tfsdk:"customer_config" json:"customer_config,computed"`
	ExternalID            types.String                                                              `tfsdk:"external_id" json:"external_id,computed"`
	IngestAliases         customfield.List[types.String]                                            `tfsdk:"ingest_aliases" json:"ingest_aliases,computed"`
	Name                  types.String                                                              `tfsdk:"name" json:"name,computed"`
	ArchivedAt            timetypes.RFC3339                                                         `tfsdk:"archived_at" json:"archived_at,computed" format:"date-time"`
	CurrentBillableStatus customfield.NestedObject[V1CustomersCurrentBillableStatusDataSourceModel] `tfsdk:"current_billable_status" json:"current_billable_status,computed"`
}

type V1CustomersCustomerConfigDataSourceModel struct {
	SalesforceAccountID types.String `tfsdk:"salesforce_account_id" json:"salesforce_account_id,computed"`
}

type V1CustomersCurrentBillableStatusDataSourceModel struct {
	Value       types.String      `tfsdk:"value" json:"value,computed"`
	EffectiveAt timetypes.RFC3339 `tfsdk:"effective_at" json:"effective_at,computed" format:"date-time"`
}

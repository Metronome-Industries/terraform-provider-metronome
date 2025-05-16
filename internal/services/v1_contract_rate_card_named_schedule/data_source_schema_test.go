// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_rate_card_named_schedule_test

import (
	"context"
	"testing"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_contract_rate_card_named_schedule"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/test_helpers"
)

func TestV1ContractRateCardNamedScheduleDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_contract_rate_card_named_schedule.V1ContractRateCardNamedScheduleDataSourceModel)(nil)
	schema := v1_contract_rate_card_named_schedule.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

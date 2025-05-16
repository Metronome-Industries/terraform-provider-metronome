// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_rate_card_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/metronome-terraform/internal/services/v1_contract_rate_card"
	"github.com/stainless-sdks/metronome-terraform/internal/test_helpers"
)

func TestV1ContractRateCardDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_contract_rate_card.V1ContractRateCardDataSourceModel)(nil)
	schema := v1_contract_rate_card.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

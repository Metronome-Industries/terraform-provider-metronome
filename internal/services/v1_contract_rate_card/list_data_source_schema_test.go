// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_rate_card_test

import (
	"context"
	"testing"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_contract_rate_card"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/test_helpers"
)

func TestV1ContractRateCardsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_contract_rate_card.V1ContractRateCardsDataSourceModel)(nil)
	schema := v1_contract_rate_card.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

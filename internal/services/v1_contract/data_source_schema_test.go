// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/metronome-terraform/internal/services/v1_contract"
	"github.com/stainless-sdks/metronome-terraform/internal/test_helpers"
)

func TestV1ContractDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_contract.V1ContractDataSourceModel)(nil)
	schema := v1_contract.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

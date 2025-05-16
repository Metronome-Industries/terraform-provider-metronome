// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_test

import (
	"context"
	"testing"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_contract"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/test_helpers"
)

func TestV1ContractModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_contract.V1ContractModel)(nil)
	schema := v1_contract.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

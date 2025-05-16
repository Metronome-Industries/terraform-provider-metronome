// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_named_schedule_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/metronome-terraform/internal/services/v1_contract_named_schedule"
	"github.com/stainless-sdks/metronome-terraform/internal/test_helpers"
)

func TestV1ContractNamedScheduleModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_contract_named_schedule.V1ContractNamedScheduleModel)(nil)
	schema := v1_contract_named_schedule.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

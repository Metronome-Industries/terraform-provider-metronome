// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_credit_grant_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/metronome-terraform/internal/services/v1_credit_grant"
	"github.com/stainless-sdks/metronome-terraform/internal/test_helpers"
)

func TestV1CreditGrantModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_credit_grant.V1CreditGrantModel)(nil)
	schema := v1_credit_grant.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

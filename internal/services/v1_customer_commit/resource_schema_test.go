// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_commit_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/metronome-terraform/internal/services/v1_customer_commit"
	"github.com/stainless-sdks/metronome-terraform/internal/test_helpers"
)

func TestV1CustomerCommitModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_customer_commit.V1CustomerCommitModel)(nil)
	schema := v1_customer_commit.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

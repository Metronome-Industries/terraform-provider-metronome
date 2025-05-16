// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_test

import (
	"context"
	"testing"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_customer"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/test_helpers"
)

func TestV1CustomerModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_customer.V1CustomerModel)(nil)
	schema := v1_customer.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

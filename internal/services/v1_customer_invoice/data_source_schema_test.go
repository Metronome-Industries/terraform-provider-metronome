// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_invoice_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/metronome-terraform/internal/services/v1_customer_invoice"
	"github.com/stainless-sdks/metronome-terraform/internal/test_helpers"
)

func TestV1CustomerInvoiceDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_customer_invoice.V1CustomerInvoiceDataSourceModel)(nil)
	schema := v1_customer_invoice.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

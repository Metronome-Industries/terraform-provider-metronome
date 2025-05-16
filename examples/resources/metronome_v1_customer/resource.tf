resource "metronome_v1_customer" "example_v1_customer" {
  name = "name"
  billing_config = {
    billing_provider_customer_id = "billing_provider_customer_id"
    billing_provider_type = "aws_marketplace"
    aws_is_subscription_product = true
    aws_product_code = "aws_product_code"
    aws_region = "af-south-1"
    stripe_collection_method = "charge_automatically"
  }
  custom_fields = {
    foo = "string"
  }
  customer_billing_provider_configurations = [{
    billing_provider = "aws_marketplace"
    configuration = {
      "0" = "bar"
      "1" = "bar"
      "2" = "bar"
    }
    delivery_method = "direct_to_billing_provider"
    delivery_method_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  }]
  external_id = "x"
  ingest_aliases = ["x"]
}

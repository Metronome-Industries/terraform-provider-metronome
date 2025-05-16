resource "metronome_v1_customer_billing_config" "example_v1_customer_billing_config" {
  customer_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  billing_provider_type = "aws_marketplace"
  billing_provider_customer_id = "billing_provider_customer_id"
  aws_product_code = "aws_product_code"
  aws_region = "af-south-1"
  stripe_collection_method = "charge_automatically"
}

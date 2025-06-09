resource "metronome_v1_customer_commit" "example_v1_customer_commit" {
  access_schedule = {
    schedule_items = [{
      amount = 0
      ending_before = "2019-12-27T18:11:19.117Z"
      starting_at = "2019-12-27T18:11:19.117Z"
    }]
    credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  }
  customer_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  priority = 0
  product_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  type = "PREPAID"
  applicable_contract_ids = ["string"]
  applicable_product_ids = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
  applicable_product_tags = ["string"]
  custom_fields = {
    foo = "string"
  }
  description = "description"
  invoice_contract_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  invoice_schedule = {
    credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    recurring_schedule = {
      amount_distribution = "DIVIDED"
      ending_before = "2019-12-27T18:11:19.117Z"
      frequency = "MONTHLY"
      starting_at = "2019-12-27T18:11:19.117Z"
      amount = 0
      quantity = 0
      unit_price = 0
    }
    schedule_items = [{
      timestamp = "2019-12-27T18:11:19.117Z"
      amount = 0
      quantity = 0
      unit_price = 0
    }]
  }
  name = "x"
  netsuite_sales_order_id = "netsuite_sales_order_id"
  rate_type = "COMMIT_RATE"
  salesforce_opportunity_id = "salesforce_opportunity_id"
  specifiers = [{
    presentation_group_values = {
      foo = "string"
    }
    pricing_group_values = {
      foo = "string"
    }
    product_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    product_tags = ["string"]
  }]
  uniqueness_key = "x"
}

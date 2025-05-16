resource "metronome_v1_credit_grant" "example_v1_credit_grant" {
  customer_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  expires_at = "2019-12-27T18:11:19.117Z"
  grant_amount = {
    amount = 0
    credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  }
  name = "name"
  paid_amount = {
    amount = 0
    credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  }
  priority = 0
  credit_grant_type = "credit_grant_type"
  custom_fields = {
    foo = "string"
  }
  effective_at = "2019-12-27T18:11:19.117Z"
  invoice_date = "2019-12-27T18:11:19.117Z"
  product_ids = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
  reason = "reason"
  rollover_settings = {
    expires_at = "2019-12-27T18:11:19.117Z"
    priority = 0
    rollover_amount = {
      type = "MAX_PERCENTAGE"
      value = 0
    }
  }
  uniqueness_key = "x"
}

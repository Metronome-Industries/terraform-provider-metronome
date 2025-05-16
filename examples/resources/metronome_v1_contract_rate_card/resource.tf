resource "metronome_v1_contract_rate_card" "example_v1_contract_rate_card" {
  name = "name"
  aliases = [{
    name = "name"
    ending_before = "2019-12-27T18:11:19.117Z"
    starting_at = "2019-12-27T18:11:19.117Z"
  }]
  credit_type_conversions = [{
    custom_credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
    fiat_per_custom_credit = 0
  }]
  custom_fields = {
    foo = "string"
  }
  description = "description"
  fiat_credit_type_id = "2714e483-4ff1-48e4-9e25-ac732e8f24f2"
}

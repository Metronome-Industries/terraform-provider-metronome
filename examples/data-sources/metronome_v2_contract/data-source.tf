data "metronome_v2_contract" "example_v2_contract" {
  contract_id = "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"
  customer_id = "13117714-3f05-48e5-a6e9-a66093f13b4d"
  as_of_date = "2019-12-27T18:11:19.117Z"
  include_balance = true
  include_ledgers = true
}

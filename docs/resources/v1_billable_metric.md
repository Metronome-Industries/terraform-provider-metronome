---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "metronome_v1_billable_metric Resource - metronome"
subcategory: ""
description: |-
  
---

# metronome_v1_billable_metric (Resource)



## Example Usage

```terraform
resource "metronome_v1_billable_metric" "example_v1_billable_metric" {
  name = "name"
  aggregation_key = "aggregation_key"
  aggregation_type = "COUNT"
  custom_fields = {
    foo = "string"
  }
  event_type_filter = {
    in_values = ["string"]
    not_in_values = ["string"]
  }
  group_keys = [["string"]]
  property_filters = [{
    name = "name"
    exists = true
    in_values = ["string"]
    not_in_values = ["string"]
  }]
  sql = "sql"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The display name of the billable metric.

### Optional

- `aggregation_key` (String) Specifies the type of aggregation performed on matching events. Required if `sql` is not provided.
- `aggregation_type` (String) Specifies the type of aggregation performed on matching events.
Available values: "COUNT", "LATEST", "MAX", "SUM", "UNIQUE".
- `billable_metric_id` (String)
- `custom_fields` (Map of String) Custom fields to attach to the billable metric.
- `event_type_filter` (Attributes) An optional filtering rule to match the 'event_type' property of an event. (see [below for nested schema](#nestedatt--event_type_filter))
- `group_keys` (List of List of String) Property names that are used to group usage costs on an invoice. Each entry represents a set of properties used to slice events into distinct buckets.
- `property_filters` (Attributes List) A list of filters to match events to this billable metric. Each filter defines a rule on an event property. All rules must pass for the event to match the billable metric. (see [below for nested schema](#nestedatt--property_filters))
- `sql` (String) The SQL query associated with the billable metric. This field is mutually exclusive with aggregation_type, event_type_filter, property_filters, aggregation_key, and group_keys. If provided, these other fields must be omitted.

### Read-Only

- `data` (Attributes) (see [below for nested schema](#nestedatt--data))

<a id="nestedatt--event_type_filter"></a>
### Nested Schema for `event_type_filter`

Optional:

- `in_values` (List of String) A list of event types that are explicitly included in the billable metric. If specified, only events of these types will match the billable metric. Must be non-empty if present.
- `not_in_values` (List of String) A list of event types that are explicitly excluded from the billable metric. If specified, events of these types will not match the billable metric. Must be non-empty if present.


<a id="nestedatt--property_filters"></a>
### Nested Schema for `property_filters`

Required:

- `name` (String) The name of the event property.

Optional:

- `exists` (Boolean) Determines whether the property must exist in the event. If true, only events with this property will pass the filter. If false, only events without this property will pass the filter. If null or omitted, the existence of the property is optional.
- `in_values` (List of String) Specifies the allowed values for the property to match an event. An event will pass the filter only if its property value is included in this list. If undefined, all property values will pass the filter. Must be non-empty if present.
- `not_in_values` (List of String) Specifies the values that prevent an event from matching the filter. An event will not pass the filter if its property value is included in this list. If null or empty, all property values will pass the filter. Must be non-empty if present.


<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `aggregation_key` (String) A key that specifies which property of the event is used to aggregate data. This key must be one of the property filter names and is not applicable when the aggregation type is 'count'.
- `aggregation_type` (String) Specifies the type of aggregation performed on matching events.
Available values: "COUNT", "LATEST", "MAX", "SUM", "UNIQUE".
- `archived_at` (String) RFC 3339 timestamp indicating when the billable metric was archived. If not provided, the billable metric is not archived.
- `custom_fields` (Map of String)
- `event_type_filter` (Attributes) An optional filtering rule to match the 'event_type' property of an event. (see [below for nested schema](#nestedatt--data--event_type_filter))
- `group_keys` (List of List of String) Property names that are used to group usage costs on an invoice. Each entry represents a set of properties used to slice events into distinct buckets.
- `id` (String) ID of the billable metric
- `name` (String) The display name of the billable metric.
- `property_filters` (Attributes List) A list of filters to match events to this billable metric. Each filter defines a rule on an event property. All rules must pass for the event to match the billable metric. (see [below for nested schema](#nestedatt--data--property_filters))
- `sql` (String) The SQL query associated with the billable metric

<a id="nestedatt--data--event_type_filter"></a>
### Nested Schema for `data.event_type_filter`

Read-Only:

- `in_values` (List of String) A list of event types that are explicitly included in the billable metric. If specified, only events of these types will match the billable metric. Must be non-empty if present.
- `not_in_values` (List of String) A list of event types that are explicitly excluded from the billable metric. If specified, events of these types will not match the billable metric. Must be non-empty if present.


<a id="nestedatt--data--property_filters"></a>
### Nested Schema for `data.property_filters`

Read-Only:

- `exists` (Boolean) Determines whether the property must exist in the event. If true, only events with this property will pass the filter. If false, only events without this property will pass the filter. If null or omitted, the existence of the property is optional.
- `in_values` (List of String) Specifies the allowed values for the property to match an event. An event will pass the filter only if its property value is included in this list. If undefined, all property values will pass the filter. Must be non-empty if present.
- `name` (String) The name of the event property.
- `not_in_values` (List of String) Specifies the values that prevent an event from matching the filter. An event will not pass the filter if its property value is included in this list. If null or empty, all property values will pass the filter. Must be non-empty if present.

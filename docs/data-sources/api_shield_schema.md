---
page_title: "cloudflare_api_shield_schema Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_api_shield_schema (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `created_at` (String)
- `find_one_by` (Attributes) (see [below for nested schema](#nestedatt--find_one_by))
- `kind` (String) Kind of schema
- `name` (String) Name of the schema
- `omit_source` (Boolean) Omit the source-files of schemas and only retrieve their meta-data.
- `schema_id` (String)
- `source` (String) Source of the schema
- `validation_enabled` (Boolean) Flag whether schema is enabled for validation.
- `zone_id` (String) Identifier

<a id="nestedatt--find_one_by"></a>
### Nested Schema for `find_one_by`

Required:

- `zone_id` (String) Identifier

Optional:

- `omit_source` (Boolean) Omit the source-files of schemas and only retrieve their meta-data.
- `page` (String) Page number of paginated results.
- `per_page` (String) Maximum number of results per page.
- `validation_enabled` (Boolean) Flag whether schema is enabled for validation.


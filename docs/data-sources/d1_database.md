---
page_title: "cloudflare_d1_database Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_d1_database (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) Account identifier tag.
- `created_at` (String) Specifies the timestamp the resource was created as an ISO8601 string.
- `database_id` (String)
- `file_size` (Number) The D1 database's size, in bytes.
- `find_one_by` (Attributes) (see [below for nested schema](#nestedatt--find_one_by))
- `name` (String)
- `num_tables` (Number)
- `uuid` (String)
- `version` (String)

<a id="nestedatt--find_one_by"></a>
### Nested Schema for `find_one_by`

Required:

- `account_id` (String) Account identifier tag.

Optional:

- `name` (String) a database name to search for.
- `page` (Number) Page number of paginated results.
- `per_page` (Number) Number of items per page.


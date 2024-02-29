# gabi-query-preparer
## Description

In order to run queries with [Gabi][gabi-url], you need to give it the queries with the following JSON format:

```json
{
  "query": "SELECT \"mt\".\"id\" FROM \"my_table\" AS \"mt\""
}
```

Even though the above query is overly complex with the double quote skipping on purpose, you can imagine how difficult
it is having to format a complex SQL query with joins, selects, sub-queries and the like to JSON. Even more if you have
to go back and forth making tweaks to the query, and modifying just minor parts of it.

This small utility program takes a `sql` file and outputs it in the standard output ready to be used by Gabi.

## Usage

1. Build it if you haven't already with `go build .`
2. Run it with `./gabi-query-preparer --query-file-path query.sql`

[gabi-url]: https://github.com/app-sre/gabi

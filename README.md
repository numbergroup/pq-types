# pq-types 

This Go package provides additional types for PostgreSQL:

* `Int32Array` for `int[]` (compatible with [`intarray`](http://www.postgresql.org/docs/current/static/intarray.html) module);
* `Int64Array` for `bigint[]`;
* `StringArray` for `varchar[]`;
* `JSONText` for `varchar`, `text`, `json` and `jsonb`;
* `PostGISPoint`, `PostGISBox2D` and `PostGISPolygon`.

Install it: `go get github.com/numbergroup/pq-types`

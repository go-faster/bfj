# bfj

```console
go test -count 1 -benchmem -bench . | bfj | jq
```

```json
{
  "name_full": "BenchmarkFoo-32",
  "name": "Foo",
  "cpus": 32,
  "count": 9704703,
  "duration_ns": 152.8,
  "mb_per_sec": 804.96,
  "alloc_bytes": 16,
  "alloc_count": 1
}
```

package pkg

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target ./ent --feature sql/lock,sql/upsert,sql/modifier ../schema

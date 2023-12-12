# business-things

A playground where I figure mapping between the business domain, openapi, ent etc.

```
go run -mod=mod entgo.io/ent/cmd/ent new Car
go generate ./ent

brew install ariga/tap/atlas
atlas schema inspect \
  -u "ent://ent/schema" \
  --dev-url "sqlite://file?mode=memory&_fk=1" \
  -w

atlas migrate diff migration_name \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "sqlite://file?mode=memory&_fk=1"

atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url "sqlite://file.db?_fk=1"
```

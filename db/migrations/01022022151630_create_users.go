package migrations

import (
	r "github.com/go-rel/rel"
)

func MigrateCreateUsers(schema *r.Schema) {
	schema.CreateTableIfNotExists("users", func(t *r.Table) {
		t.ID("id", r.Primary(true))
		t.String("first_name", r.Required(true))
		t.String("last_name", r.Required(true))
		t.Int("age", r.Required(true))
		t.DateTime("created_at", r.Required(true))
		t.DateTime("updated_at", r.Required(true))
	})

	schema.CreateIndex("users", "UI_users_id", []string{"id"})
}

func RollbackCreateUsers(schema *r.Schema) {
	schema.DropIndex("users", "user_space.UI_users_id")
	schema.DropTableIfExists("user_space.users")
}

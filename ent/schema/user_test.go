package schema_test

import (
	"context"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"

	"issue/ent"
	"issue/ent/enttest"
)

func TestUserDefaultValue(t *testing.T) {
	c := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	u := c.User.Create().SetInput(
		ent.CreateUserInput{
			Age:  28,
			Name: "Yoni",
		}).SaveX(context.Background())
	require.Equal(t, "Yoni", u.Name)
	require.Equal(t, 28, u.Age)
}

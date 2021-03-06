package schema_test

import (
	"context"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"

	"issue/ent/enttest"
)

func TestUserDefaultValue(t *testing.T) {
	c := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	u := c.User.Create().SetAge(38).SaveX(context.Background())
	require.Equal(t, time.Now().Round(time.Minute), u.CreatedAt.Round(time.Minute))
}

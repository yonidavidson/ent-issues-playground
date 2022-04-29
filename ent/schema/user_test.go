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
	"issue/ent/user"
)

func TestUserDefaultValue(t *testing.T) {
	ctx := context.Background()
	c := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	c.User.Create().SetAge(38).ExecX(ctx)
	c.User.Create().SetAge(20).ExecX(ctx)
	u := c.User.Query().Order(ent.Asc(user.FieldAge)).FirstX(ctx)
	require.Equal(t, 20, u.Age)
	u = c.User.Query().Order(ent.Desc(user.FieldAge)).FirstX(ctx)
	require.Equal(t, 38, u.Age)
}

func TestUserSum(t *testing.T) {
	ctx := context.Background()
	c := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	c.User.Create().SetAge(20).ExecX(ctx)
	c.User.Create().SetAge(20).ExecX(ctx)
	c.User.Create().SetAge(38).ExecX(ctx)
	type r struct {
		Age   int `json:"age"`
		Sum   int `json:"sum"`
		Count int `json:"count"`
	}
	var v []r
	expected := []r{{Age: 20, Sum: 40, Count: 2}, {Age: 38, Sum: 38, Count: 1}}
	c.User.Query().
		GroupBy(user.FieldAge).
		Aggregate(ent.Count(), ent.Sum(user.FieldAge)).
		ScanX(ctx, &v)
	require.EqualValues(t, expected, v)
}

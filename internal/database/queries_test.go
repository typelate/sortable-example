package database_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/typelate/sortable-example/internal/database"
)

func setup(t *testing.T) (context.Context, *database.Queries) {
	t.Helper()
	ctx := t.Context()
	pool, err := database.Setup(ctx)
	if err != nil {
		t.Skipf("database not available: %v", err)
	}
	t.Cleanup(pool.Close)
	tx, err := pool.Begin(ctx)
	require.NoError(t, err)
	t.Cleanup(func() { _ = tx.Rollback(context.Background()) })
	return ctx, database.New(tx)
}

func TestQueries(t *testing.T) {
	ctx, q := setup(t)

	lists, err := q.Lists(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, lists)

	list, err := q.ListByID(ctx, lists[0].ID)
	require.NoError(t, err)
	assert.Equal(t, lists[0], list)

	tasks, err := q.TasksByListID(ctx, list.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tasks)
	for i := 1; i < len(tasks); i++ {
		assert.GreaterOrEqual(t, tasks[i-1].Priority, tasks[i].Priority)
	}

	task := tasks[len(tasks)-1]
	newPriority := tasks[0].Priority + 1
	err = q.SetTaskPriority(ctx, database.SetTaskPriorityParams{
		Priority: newPriority,
		ID:       task.ID,
	})
	require.NoError(t, err)

	updated, err := q.TasksByListID(ctx, list.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updated)
	assert.Equal(t, task.ID, updated[0].ID)
	assert.Equal(t, newPriority, updated[0].Priority)
}
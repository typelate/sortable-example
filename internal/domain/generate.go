package domain

//go:generate rm -rf internal/fake
//go:generate mkdir -p internal/fake
//go:generate counterfeiter -generate

//counterfeiter:generate -o internal/fake/read_only_querier.go     --fake-name ReadOnlyQuerier     ../database.ReadOnlyQuerier
//counterfeiter:generate -o internal/fake/task_priority_updater.go --fake-name TaskPriorityUpdater ../database.TaskPriorityUpdater
//counterfeiter:generate -o internal/fake/transaction_manager.go   --fake-name TransactionManager  . TransactionManager

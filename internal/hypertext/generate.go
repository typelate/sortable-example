package hypertext

//go:generate muxt generate --use-receiver-type-package=github.com/typelate/sortable-example/internal/domain --use-receiver-type=Service --output-routes-func=Routes
//go:generate rm -rf internal/fake
//go:generate mkdir -p internal/fake
//go:generate counterfeiter -generate

//counterfeiter:generate -o internal/fake/routes_receiver.go --fake-name RoutesReceiver . RoutesReceiver

package hypertext

//go:generate muxt generate --receiver-type-package=github.com/typelate/sortable-example/internal/domain --receiver-type=Service --routes-func Routes
//go:generate rm -rf internal/fake
//go:generate mkdir -p internal/fake
//go:generate counterfeiter -generate

//counterfeiter:generate -o internal/fake/routes_receiver.go --fake-name RoutesReceiver . RoutesReceiver

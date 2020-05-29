package wait_test

import (
	"context"

	"github.com/yext/testcontainers-go"
	"github.com/yext/testcontainers-go/wait"
)

//
// https://github.com/yext/testcontainers-go/issues/183
func ExampleHTTPStrategy() {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "gogs/gogs:0.11.91",
		ExposedPorts: []string{"3000/tcp"},
		WaitingFor:   wait.ForHTTP("/").WithPort("3000/tcp"),
	}

	gogs, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}

	defer gogs.Terminate(ctx)

	// Here you have a running container
}

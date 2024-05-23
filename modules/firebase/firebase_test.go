package firebase_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/firebase"
)

func TestFirebase(t *testing.T) {
	ctx := context.Background()

	thing := fmt.Sprintf("%s/test", os.Getenv("PWD"))
	fmt.Println(thing)

	container, err := firebase.RunContainer(
		ctx,
		testcontainers.WithImage(firebase.IMAGE_NAME),
		firebase.WithRoot(fmt.Sprintf("%s/firebase", os.Getenv("PWD"))),
		firebase.WithCache(),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container after the test is complete
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	// perform assertions
}

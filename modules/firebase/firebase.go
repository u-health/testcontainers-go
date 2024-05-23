package firebase

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/mount"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"strings"
	"time"
)

// FirebaseContainer represents the Firebase container type used in the module
type FirebaseContainer struct {
	testcontainers.Container
}

const (
	UI_PORT        = "4000/tcp"
	HUB_PORT       = "4400/tcp"
	LOGGING_PORT   = "4600/tcp"
	FUNCTIONS_PORT = "5001/tcp"
	FIRESTORE_PORT = "8080/tcp"
	PUBSUB_PORT    = "8085/tcp"
	DATABASE_PORT  = "9000/tcp"
	AUTH_PORT      = "9099/tcp"
	STORAGE_PORT   = "9199/tcp"
	HOSTING_PORT   = "6000/tcp"
)

const IMAGE_NAME = "ghcr.io/u-health/docker-firebase-emulator:13.6.0"

// WithRoot changes the default firebase root path on local machine
func WithRoot(rootPath string) testcontainers.CustomizeRequestOption {
	return func(req *testcontainers.GenericContainerRequest) error {
		if !strings.HasSuffix(rootPath, "/firebase") {
			return fmt.Errorf("root path must end with '/firebase': %s", rootPath)
		}
		req.Files = append(req.Files, testcontainers.ContainerFile{
			HostFilePath:      rootPath,
			ContainerFilePath: "/srv/firebase",
		})

		return nil
	}
}

// WithData names the data directory in firebase mount
func WithData(dataPath string) testcontainers.CustomizeRequestOption {
	return func(req *testcontainers.GenericContainerRequest) error {
		req.Env["DATA_DIRECTORY"] = dataPath
		return nil
	}
}

func cache(volumeName string, volumeOptions *mount.VolumeOptions) testcontainers.CustomizeRequestOption {
	return func(req *testcontainers.GenericContainerRequest) error {
		m := testcontainers.ContainerMount{
			Source: testcontainers.DockerVolumeMountSource{
				Name:          volumeName,
				VolumeOptions: volumeOptions,
			},
			Target: "/root/.cache/firebase",
		}
		req.Mounts = append(req.Mounts, m)
		return nil
	}
}

// WithCache enables firebase binary cache based on session (meaningful only when multiple tests are used)
func WithCache() testcontainers.CustomizeRequestOption {
	volumeName := fmt.Sprintf("firestore-cache-%s", testcontainers.SessionID())
	volumeOptions := &mount.VolumeOptions{
		Labels: testcontainers.GenericLabels(),
	}

	return cache(volumeName, volumeOptions)
}

// RunContainer creates an instance of the Firebase container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*FirebaseContainer, error) {
	req := testcontainers.ContainerRequest{
		Image: IMAGE_NAME,
		ExposedPorts: []string{
			UI_PORT,
			HUB_PORT,
			LOGGING_PORT,
			FUNCTIONS_PORT,
			FIRESTORE_PORT,
			PUBSUB_PORT,
			DATABASE_PORT,
			AUTH_PORT,
			STORAGE_PORT,
			HOSTING_PORT,
		},

		Env: map[string]string{},
		WaitingFor: wait.ForAll(
			wait.ForHTTP("/").WithPort("4000").WithStartupTimeout(3 * time.Minute),
		),
	}

	genericContainerReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}

	for _, opt := range opts {
		if err := opt.Customize(&genericContainerReq); err != nil {
			return nil, fmt.Errorf("customize: %w", err)
		}
	}

	container, err := testcontainers.GenericContainer(ctx, genericContainerReq)
	if err != nil {
		return nil, err
	}

	return &FirebaseContainer{Container: container}, nil
}

func (c *FirebaseContainer) FirestoreConnectionString(ctx context.Context) (string, error) {
	host, err := c.Host(ctx)
	if err != nil {
		return "", err
	}
	port, err := c.MappedPort(ctx, FIRESTORE_PORT)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", host, port.Port()), nil
}

func (c *FirebaseContainer) AuthConnectionString(ctx context.Context) (string, error) {
	host, err := c.Host(ctx)
	if err != nil {
		return "", err
	}
	port, err := c.MappedPort(ctx, AUTH_PORT)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", host, port.Port()), nil
}

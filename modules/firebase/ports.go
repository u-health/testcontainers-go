package firebase

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
)

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

func (c *FirebaseContainer) connectionString(ctx context.Context, portName nat.Port) (string, error) {
	host, err := c.Host(ctx)
	if err != nil {
		return "", err
	}
	port, err := c.MappedPort(ctx, portName)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", host, port.Port()), nil
}

func (c *FirebaseContainer) UIConnectionString(ctx context.Context) (string, error) {
	return c.connectionString(ctx, UI_PORT)
}

func (c *FirebaseContainer) HubConnectionString(ctx context.Context) (string, error) {
	return c.connectionString(ctx, HUB_PORT)
}

func (c *FirebaseContainer) LoggingConnectionString(ctx context.Context) (string, error) {
	return c.connectionString(ctx, LOGGING_PORT)
}

func (c *FirebaseContainer) FunctionsConnectionString(ctx context.Context) (string, error) {
	return c.connectionString(ctx, FUNCTIONS_PORT)
}

func (c *FirebaseContainer) FirestoreConnectionString(ctx context.Context) (string, error) {
	return c.connectionString(ctx, FIRESTORE_PORT)
}

func (c *FirebaseContainer) PubSubConnectionString(ctx context.Context) (string, error) {
	return c.connectionString(ctx, PUBSUB_PORT)
}

func (c *FirebaseContainer) DatabaseConnectionString(ctx context.Context) (string, error) {
	return c.connectionString(ctx, DATABASE_PORT)
}

func (c *FirebaseContainer) AuthConnectionString(ctx context.Context) (string, error) {
	return c.connectionString(ctx, AUTH_PORT)
}

func (c *FirebaseContainer) StorageConnectionString(ctx context.Context) (string, error) {
	return c.connectionString(ctx, STORAGE_PORT)
}

func (c *FirebaseContainer) HostingConnectionString(ctx context.Context) (string, error) {
	return c.connectionString(ctx, HOSTING_PORT)
}

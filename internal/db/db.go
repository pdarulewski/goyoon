// Package db manages the sqlite db connection for keeping track of user progress.
package db

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	// sqlite driver.
	_ "modernc.org/sqlite"

	"github.com/adrg/xdg"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

const (
	// GoyoonDir is the directory name for storing application data.
	GoyoonDir = "goyoon"

	// dirPerm is goyoon permission.
	dirPerm = 0o750
)

// Schema stores sql migrations.
//
//go:embed sql/schema.sql
var Schema embed.FS

// ErrDB is a general db error.
var ErrDB = errors.New("db error")

// DB stores the db connection.
type DB struct {
	db *sql.DB
}

// New creates a new database connection.
func New(ctx context.Context, dbPath string) (*DB, error) {
	_, err := os.Stat(dbPath)
	if err != nil {
		return nil, fmt.Errorf("%w: problem with db file: %w", ErrDB, err)
	}

	database, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("%w: connection error: %w", ErrDB, err)
	}

	if err := database.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("%w: ping: %w", ErrDB, err)
	}

	return &DB{
		db: database,
	}, nil
}

// Migrate runs an up migration on a given database.
func (d *DB) Migrate() error {
	databaseInstance, err := sqlite.WithInstance(d.db, &sqlite.Config{})
	if err != nil {
		return fmt.Errorf("")
	}

	sourceInstance, err := iofs.New(Schema, ".")
	if err != nil {
		return fmt.Errorf("")
	}

	migrator, err := migrate.NewWithInstance("iofs", sourceInstance, "sqlite", databaseInstance)
	if err != nil {
		return fmt.Errorf("")
	}

	if err := migrator.Up(); err != nil {
		return fmt.Errorf("")
	}

	return nil
}

type PathProvider interface {
	Path() string
}

// XDGStatePathProvider provides a path to the state directory.
type XDGStatePathProvider struct{}

// Path returns a path to the sqlite file in the XDG state directory.
func (*XDGStatePathProvider) Path() string {
	return filepath.Join(xdg.StateHome, GoyoonDir, ".sqlite")
}

// Initialize sets up the sqlite database if it doesn't exist.
func Initialize(provider PathProvider) error {
	dbPath := filepath.Clean(provider.Path())

	_, err := os.Stat(dbPath)
	if err == nil {
		return nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("%w: file error: %w", ErrDB, err)
	}

	if err := os.MkdirAll(filepath.Dir(dbPath), dirPerm); err != nil {
		return fmt.Errorf("%w: creating directory: %w", ErrDB, err)
	}

	file, err := os.Create(dbPath)
	if err != nil {
		return fmt.Errorf("%w: creating db file: %w", ErrDB, err)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("%w:	closing db file: %w", ErrDB, err)
	}

	return nil
}

package db_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pdarulewski/goyoon/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		dbPath := filepath.Clean(filepath.Join(t.TempDir(), ".sqlite"))
		file, err := os.Create(dbPath)
		require.NoError(t, err)

		err = file.Close()
		require.NoError(t, err)

		db, err := db.New(t.Context(), dbPath)
		require.NoError(t, err)

		assert.NotNil(t, db)
	})

	t.Run("file does not exist", func(t *testing.T) {
		t.Parallel()

		db, err := db.New(t.Context(), "abc")
		require.ErrorIs(t, err, os.ErrNotExist)

		assert.Nil(t, db)
	})
}

type MockPathProvider struct {
	path string
}

func (m *MockPathProvider) Path() string {
	return m.path
}

func TestInitialize(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		path := filepath.Join(t.TempDir(), "mydbfile")
		mock := &MockPathProvider{path: path}

		err := db.Initialize(mock)
		require.NoError(t, err)

		assert.FileExists(t, path)
	})

	t.Run("file already exists", func(t *testing.T) {
		t.Parallel()

		path := filepath.Join(t.TempDir(), "mydbfile")

		file, err := os.Create(filepath.Clean(path))
		require.NoError(t, err)

		err = file.Close()
		require.NoError(t, err)

		mock := &MockPathProvider{path: path}

		err = db.Initialize(mock)
		require.NoError(t, err)

		assert.FileExists(t, path)
	})
}

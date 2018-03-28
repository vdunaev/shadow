package internal

import (
	"time"

	"github.com/kihamo/shadow/components/database"
)

type MigrationItem struct {
	database.Migration

	source    string
	migration database.Migration
}

func NewMigrationItem(migration database.Migration, source string) *MigrationItem {
	return &MigrationItem{
		migration: migration,
		source:    source,
	}
}

func (m *MigrationItem) Source() string {
	return m.source
}

func (m *MigrationItem) Id() string {
	return m.migration.Id()
}

func (m *MigrationItem) Up() []string {
	return m.migration.Up()
}

func (m *MigrationItem) Down() []string {
	return m.migration.Down()
}

func (m *MigrationItem) ModAt() time.Time {
	return m.migration.ModAt()
}

func (m *MigrationItem) AppliedAt() *time.Time {
	return m.migration.AppliedAt()
}
package repository

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)
import _ "github.com/jinzhu/gorm/dialects/sqlite"

func NewGORMDatabaseConnection(dbType string, url string) (*gorm.DB, error) {
	db, err := gorm.Open(dbType, url)
	return db, err
}

func NewGORMSqlLiteConnection(path string) (*gorm.DB, error) {
	return NewGORMDatabaseConnection("sqlite3", path)
}

func randomUUID() string {
	s := uuid.NewV4()
	return s.String()
}

func Clean(db *gorm.DB) {
	db.DropTableIfExists(&CashInAgent{}, &ClientApplicationInformation{})
}

func DefaultGenesis(db *gorm.DB) *gorm.DB {
	genesisDB := db.Debug()
	key := randomUUID()
	hasher := sha1.New()
	hasher.Write([]byte(key))
	genesisDB.
		AutoMigrate(
			&CashInAgent{},
			&PaymentModel{},
			&PaymentStatusHistory{},
			&ClientApplicationInformation{},
		).
		FirstOrCreate(&CashInAgent{
			ID:       "SandboxCounter",
			Endpoint: "localhost:4000",
			Type:     AgentTypeSandbox,
		}, &CashInAgent{ID: "SandboxCounter", Type: AgentTypeSandbox}).
		FirstOrCreate(&ClientApplicationInformation{
			ID:     key,
			Secret: base64.URLEncoding.EncodeToString(hasher.Sum(nil)),
			Name:   "VeloGo",
		}, &ClientApplicationInformation{Name: "VeloGo"})

	return db
}

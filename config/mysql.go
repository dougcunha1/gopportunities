package config

import (
	"fmt"
	"github.com/dougcunha1/gopportunities/schemas"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

/*
	If you want to use SQLite, uncomment this function and comment the MySQL one
*/

//func InitializeSQlite() (*gorm.DB, error) {
//	logger := GetLogger("sqlite")
//	dbPath := "./db/main.db"
//	// Check if the database file exist
//	_, err := os.Stat(dbPath)
//	if os.IsNotExist(err) {
//		logger.Info("database file not found, creating...")
//		// Create database file and directory
//		err = os.MkdirAll("./db", os.ModePerm)
//		if err != nil {
//			return nil, err
//		}
//		file, err := os.Create(dbPath)
//		if err != nil {
//			return nil, err
//		}
//		file.Close()
//	}
//	// Create DB and connect
//	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
//	if err != nil {
//		logger.Errorf("sqlite opening error: %v", err)
//		return nil, err
//	}
//	// Migrate the schema
//	err = db.AutoMigrate(&schemas.Opening{})
//	if err != nil {
//		logger.Errorf("sqlite automigration error: %v", err)
//		return nil, err
//	}
//	// Return the DB
//	return db, nil
//}

// DBConfig holds the database configuration parameters
type DBConfig struct {
	UserName   string
	DBPassword string
	HostName   string
	DBPort     string
	DBName     string
	Charset    string
	ParseTime  string
	Loc        string
}

// NewDBConfig creates a new DBConfig instance with values from environment variables
func NewDBConfig() *DBConfig {
	return &DBConfig{
		UserName:   os.Getenv("USER_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		HostName:   os.Getenv("HOST_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		Charset:    os.Getenv("CHARSET"),
		ParseTime:  os.Getenv("PARSE_TIME"),
		Loc:        os.Getenv("LOC"),
	}
}

// DSN constructs the data source name (DSN) for the database connection
func (c *DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		c.UserName, c.DBPassword, c.HostName, c.DBPort, c.DBName, c.Charset, c.ParseTime, c.Loc)
}

func InitializeMySQL() (*gorm.DB, error) {
	// Initialize logger for mysql
	logger := GetLogger("mysql")
	// Load .env key/value pairs
	err := godotenv.Load()
	if err != nil {
		logger.Errorf("error loading .env file: %v", err)
		return nil, err
	}
	// Create a new DBConfig instance
	config := NewDBConfig()
	// Construct DSN
	dsn := config.DSN()
	// Connect to MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("mysql opening error: %v", err)
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("mysql automigration error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}

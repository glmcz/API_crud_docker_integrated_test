package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"simpleCloudService/internal/model"
)

type PostgresConfig struct {
	Uri      string `yaml:"uri"`
	Database string `yaml:"database"`
	//Host     string `yaml:"host"`
	//User     string `yaml:"user"`
	//Password string `yaml:"password"`
	//DBName   string `yaml:"dbname"`
	//Port     string `yaml:"port"`
	//SSLMode  string `yaml:"sslmode"`
	//TimeZone string `yaml:"timezone"`
}

func (p *PostgresConfig) ToString() string {
	re := regexp.MustCompile(`dbname=\S*`)
	// Replace the matched pattern with "dbname=<new_value>"
	updatedUri := re.ReplaceAllString(p.Uri, fmt.Sprintf("dbname=%s", p.Database))
	return updatedUri

	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", p.Host, p.User, p.Password, p.DBName, p.Port, p.SSLMode, p.TimeZone)
}

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(config *PostgresConfig) (*PostgresRepository, error) {
	// TODO create a connection pool?
	db, err := gorm.Open(postgres.Open(config.ToString()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database")
	}

	return &PostgresRepository{db: db}, nil
}

func (db *PostgresRepository) EmptyAutoMigrate() error {
	err := db.db.AutoMigrate(&model.User{})
	if err != nil {
		return fmt.Errorf("failed to auto migrate users table")
	}

	return nil
}

func (db *PostgresRepository) CreateUser(user *model.User) error {
	if err := db.db.Create(user).Error; err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (db *PostgresRepository) GetUser(userID uuid.UUID) (*model.User, error) {
	var user model.User
	if err := db.db.First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return &user, nil
}

func (db *PostgresRepository) UpdateUser(user *model.User) error {
	if err := db.db.Save(user).Error; err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (db *PostgresRepository) DeleteUser(userID string) error {
	if err := db.db.Delete(userID).Error; err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

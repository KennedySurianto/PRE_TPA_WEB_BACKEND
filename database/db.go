package database

import (
    "pre_tpa_web/model"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// Exported function
func ConnectDatabase() *gorm.DB {
    dbURL := "postgres://postgres:postgres@localhost:5432/pretpa"
    // <username>:<password>@localhost:5432/<db_name>

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    err = db.AutoMigrate(
        &model.User{},
    )

    if err != nil {
        panic(err)
    }
    
    return db
}
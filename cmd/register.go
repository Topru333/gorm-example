/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"` 
	Name      string
	Mail      string
	CreatedAt time.Time 
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")

		name := args[0]
		mail := args[1]

		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		login := os.Getenv("POSTGRE_LOGIN")
		password := os.Getenv("POSTGRE_PASSWORD")
		database := os.Getenv("POSTGRE_DB")
		host := os.Getenv("HOST_DB")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", host, login, password, database)
		conn := postgres.Open(dsn)
		db, err := gorm.Open(conn, &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to database")
		}
		user := User{Name: name, Mail: mail}
		result := db.Create(&user)

		if result.Error != nil {
			log.Fatalf("failed to insert user into the database: %v", result.Error)
		}
		fmt.Println("User registered successfully")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}

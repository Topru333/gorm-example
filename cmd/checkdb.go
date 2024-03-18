/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var table string

var checkdbCmd = &cobra.Command{
	Use:   "checkdb",
	Short: "check existence of table",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("checkdb called")
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

		exists := db.Migrator().HasTable(table)
		if exists {
			fmt.Printf("Table %s exists.\n", table)
		} else {
			fmt.Printf("Table %s does not exist.\n", table)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkdbCmd)
	checkdbCmd.Flags().StringVarP(&table, "table", "t", "", "table to check")
}

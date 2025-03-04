package main

import (
	"fmt"

	"be/internal/delivery/cron"
	"be/internal/repository"
	"be/internal/usecase"

	"github.com/spf13/viper"
)

func main() {
	// Initialize Viper
	viper.SetConfigName("config")    // name of config file (without extension)
	viper.SetConfigType("json")      // or viper.SetConfigType("YAML")
	viper.AddConfigPath("./configs") // optionally look for config in the configs directory
	viper.AddConfigPath(".")         // optionally look for config in the working directory
	viper.AutomaticEnv()             // read in environment variables that match

	c := cron.New()

	// Initialize dependencies
	jiraRepo := repository.NewJiraRepository()      // Database/API connection
	jiraUsecase := usecase.NewJiraUsecase(jiraRepo) // Business logic

	// Register cron jobs
	cron.RegisterJobs(c, jiraUsecase)

	fmt.Println("Starting Cron Jobs...")
	c.Start()

	// Keep the process running
	select {}
}

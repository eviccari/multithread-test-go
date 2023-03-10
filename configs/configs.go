package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/eviccari/multithread-test-go/internal/utils"
	"github.com/spf13/viper"
)

const (
	ConfigPath1 = "./cmd/env"
	ConfigPath2 = "./env"
	ConfigPath3 = "."
	ConfigType  = "json"
	ConfigName  = ".env"
)

var Environment string
var GithubUsersQuantity int
var MultiThreadSize int
var GithubAPIMaxPageSize int
var DBUser string
var DBPassword string
var DBHostName string
var DBPort string
var DBEngine string
var DBName string

func init() {
	viper.AddConfigPath(ConfigPath1)
	viper.AddConfigPath(ConfigPath2)
	viper.AddConfigPath(ConfigPath3)
	viper.SetConfigType(ConfigType)
	viper.SetConfigName(ConfigName)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error on load job configurations: %s", err.Error())
	}

	GithubUsersQuantity, _ = strconv.Atoi(os.Getenv("GITHUB_USERS_QUANTITY"))
	if GithubUsersQuantity == 0 {
		GithubUsersQuantity = viper.GetInt("GithubUsersQuantity")
	}

	MultiThreadSize, _ = strconv.Atoi(os.Getenv("MULTITHREAD_SIZE"))
	if MultiThreadSize == 0 {
		MultiThreadSize = viper.GetInt("MultiThreadSize")
	}

	GithubAPIMaxPageSize, _ = strconv.Atoi(os.Getenv("GITHUB_API_MAX_PAGE_SIZE"))
	if GithubAPIMaxPageSize == 0 {
		GithubAPIMaxPageSize = viper.GetInt("GithubAPIMaxPageSize")
	}

	Environment = os.Getenv("ENVIRONMENT")
	if utils.IsEmptyString(Environment) {
		Environment = viper.GetString("Environment")
	}

	DBUser = os.Getenv("DATABASE_USER_NAME")
	if utils.IsEmptyString(DBUser) {
		DBUser = viper.GetString("DBUser")
	}

	DBPassword = os.Getenv("DATABASE_USER_PASS")
	if utils.IsEmptyString(DBPassword) {
		DBPassword = viper.GetString("DBPassword")
	}

	DBHostName = os.Getenv("DATABASE_HOST_NAME")
	if utils.IsEmptyString(DBHostName) {
		DBHostName = viper.GetString("DBHostName")
	}

	DBPort = os.Getenv("DATABASE_PORT")
	if utils.IsEmptyString(DBPort) {
		DBPort = viper.GetString("DBPort")
	}

	DBName = os.Getenv("DATABASE_NAME")
	if utils.IsEmptyString(DBName) {
		DBName = viper.GetString("DBName")
	}

	DBEngine = os.Getenv("DATABASE_ENGINE")
	if utils.IsEmptyString(DBEngine) {
		DBEngine = viper.GetString("DBEngine")
	}

	printConfigs()
}

func printConfigs() {
	log.Printf("GithubUsersQuantity.: %d", GithubUsersQuantity)
	log.Printf("MultiThreadSize.....: %d", MultiThreadSize)
	log.Printf("GithubAPIMaxPageSize: %d", GithubAPIMaxPageSize)
	log.Printf("DBUser..............: %s", DBUser)

	if Environment == "DEV" {
		log.Printf("DBPassword..........: %s", DBPassword) //TODO - DO NOT PUT THIS IN PRODUCTION
	}

	log.Printf("DBHostName..........: %s", DBHostName)
	log.Printf("DBPort..............: %s", DBPort)
	log.Printf("DBName..............: %s", DBName)
	log.Printf("DBEngine............: %s", DBEngine)
}

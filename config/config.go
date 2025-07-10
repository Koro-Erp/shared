package config

import (
	"log"
	"os"

	"github.com/Koro-Erp/shared/models"
	"github.com/Koro-Erp/shared/util"
	"github.com/joho/godotenv"
)


func LoadConfig() (models.DbConfig,models.ServiceUrlConfig,models.KeyConfig){
	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found or couldn't load it — assuming Docker environment.")
    }

    dbConfig :=  models.DbConfig{
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        DBSSLMode:  os.Getenv("DB_SSLMODE"),
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
	}

	serviceUrlsConfig :=  models.ServiceUrlConfig{
		AuthServiceURL:  os.Getenv("AUTH_SERVICE_URL"),
        HrmServiceURL:  os.Getenv("HRM_SERVICE_URL"),
		LogsServiceURL:  os.Getenv("LOG_SERVICE_URL"),
		PmServiceUrl:  os.Getenv("PM_SERVICE_URL"),
		ScmServiceURL:  os.Getenv("SCM_SERVICE_URL"),
		AccountingServiceURL:  os.Getenv("ACCOUNTING_SERVICE_URL"),
		ManufacturingServiceURL:  os.Getenv("MANUFACTURING_SERVICE_URL"),
		NotificationsServiceURL:  os.Getenv("NOTIFICATION_SERVICE_URL"),
		SalesServiceURL:  os.Getenv("SALES_SERVICE_URL"),
		DocumentsServiceURL:  os.Getenv("DOCUMENTS_SERVICE_URL"),
		CrmServiceURL:  os.Getenv("CRM_SERVICE_URL"),

	}

	serviceUrlsConfig.SaveAppLogsUrl = serviceUrlsConfig.LogsServiceURL + "/app-logs"
	serviceUrlsConfig.SaveGatewayLogsUrl = serviceUrlsConfig.LogsServiceURL + "/gateway-logs"
	serviceUrlsConfig.CheckUserExistsUrl = serviceUrlsConfig.AuthServiceURL + "/users/%s/exists"


	KeyConfig := models.KeyConfig{
		EncryptionKey: os.Getenv("AUTO_INCREMENT_STRATEGY"),
		PublicKey: util.LoadPublicKey("internal/config/keys/public.pem"),
		JwtKey : os.Getenv("JWT_SECRET_KEY"),
	}

	return dbConfig,serviceUrlsConfig,KeyConfig
}
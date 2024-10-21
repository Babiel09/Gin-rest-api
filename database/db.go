package database

import (
	"log"
	"time"

	"github.com/Babiel09/Gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	log.Println("Iniciando a conexão com o banco de dados, isso pode demorar de 20 s ou mais.")
	time.Sleep(20 * time.Second)

	dsn := "host=go_db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Amigo{})
	log.Println("Conexão bem sucedida.")
}

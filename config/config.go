package config

import (
	a_rep "evermos_pbi/internal/features/address/repository"
	c_rep "evermos_pbi/internal/features/categories/repository"
	d_rep "evermos_pbi/internal/features/detailtransaction/repository"
	l_rep "evermos_pbi/internal/features/logproduct/repository"
	p_rep "evermos_pbi/internal/features/products/repository"
	s_rep "evermos_pbi/internal/features/stores/repository"
	t_rep "evermos_pbi/internal/features/transaction/repository"
	u_rep "evermos_pbi/internal/features/users/repository"

	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type setting struct {
	User        string
	Host        string
	Password    string
	Port        string
	DBName      string
	JWTSecret   string
	CldKey      string
	MidTransKey string
	Schema      string
}

func ImportSetting() setting {
	var result setting

	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		err := godotenv.Load(".env")
		if err != nil {
			return setting{}
		}
	} else {
		log.Println("file not exist")
	}
	result.User = os.Getenv("DB_USER")
	result.Host = os.Getenv("DB_HOST")
	result.Port = os.Getenv("DB_PORT")
	result.DBName = os.Getenv("DB_NAME")
	result.Password = os.Getenv("DB_PASSWORD")
	result.JWTSecret = os.Getenv("JWT_SECRET")
	result.CldKey = os.Getenv("CLOUDINARY_KEY")
	result.MidTransKey = os.Getenv("MIDTRANS_KEY")
	result.Schema = os.Getenv("SCHEMA")
	return result
}

func ConnectDB() (*gorm.DB, error) {
	s := ImportSetting()
	var connStr = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", s.Host, s.User, s.Password, s.Port, s.DBName)
	schem := ImportSetting().Schema
	schem = schem + "."
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: schem,
		},
	})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&u_rep.User{}, &s_rep.Store{}, &a_rep.Address{}, &c_rep.Category{}, &p_rep.Product{}, &l_rep.LogProduct{}, &t_rep.Transaction{}, &d_rep.DetailTransaction{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

package config

import (
	"fmt"
	"os"
	"reflect"
)

var (
	C Config
)

type Config struct {
	Project_path   string
	Auth_Data_Path string
	DB_Path        string
	Tmp_Path       string
}

func init() {
	fmt.Println("Initializing config")

	fmt.Println("Getting current working directory")
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	C = Config{
		Project_path:   path,
		Auth_Data_Path: path + "/auth_data/",
		DB_Path:        path + "/main.db",
		Tmp_Path:       path + "/temp/",
	}
	if os.Chdir(C.Auth_Data_Path) != nil {
		fmt.Println("Creating auth_data directory")
		os.Mkdir(C.Auth_Data_Path, 0755)
	}
	fmt.Println("Config initialized -")
	t := reflect.TypeOf(C)
	v := reflect.ValueOf(C)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Name, v.Field(i).Interface())
	}
	fmt.Println("--------------------")
}

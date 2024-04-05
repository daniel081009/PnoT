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
	Project_path string
	DB_Path      string
	Tmp_Path     string
	RootPassword string
}

func init() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	C = Config{
		Project_path: path,
		DB_Path:      path + "/main.db",
		Tmp_Path:     path + "/temp/",
		RootPassword: "",
	}

	fmt.Println("Config -------------")
	t := reflect.TypeOf(C)
	v := reflect.ValueOf(C)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Name, v.Field(i).Interface())
	}
	fmt.Println("--------------------")
}

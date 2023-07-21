package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvVaraibles(){
	if err:=godotenv.Load();err!=nil{
		 fmt.Printf("error while loading env variables")
	}
}
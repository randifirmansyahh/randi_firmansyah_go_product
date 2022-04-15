package helper

import (
	"encoding/json"
	"errors"
	"log"
	"randi_firmansyah/app/models/productModel"
	"randi_firmansyah/app/models/userModel"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CheckEnv(err error) {
	if err != nil {
		log.Fatal("Failed to load environment")
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func UnMarshall(from string, to interface{}) {
	if err := json.Unmarshal([]byte(from), &to); err != nil {
		log.Println(err.Error())
		return
	}
}

func SearchProduct(models []productModel.Product, id int) (interface{}, error) {
	var oneData productModel.Product
	for i := 0; i < len(models); i++ {
		if models[i].Id == id {
			oneData = models[i]
			return oneData, nil
		}
	}
	return nil, errors.New("data not found")
}

func SearchUser(models []userModel.User, id int) (interface{}, bool) {
	var oneData userModel.User
	for i := 0; i < len(models); i++ {
		if models[i].Id == id {
			oneData = models[i]
			return oneData, true
		}
	}
	return nil, false
}

func ExtractTokens(token string) string {
	/// Bearer asdhaskdkasdhsagdhasdgjasgdhasdghadgjadsaj
	strarr := strings.Split(token, " ")
	if len(strarr) == 2 {
		return strarr[1]
	}
	return ""
}

func ExpiredTime(menit int) time.Time {
	return time.Now().Add(time.Duration(menit) * time.Minute) // expired date
}

func CheckError(err error) {
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func CheckFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func StringToint(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println("error =>", err)
	}
	return i
}

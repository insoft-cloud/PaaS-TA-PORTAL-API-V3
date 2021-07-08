package config

import (
	"bufio"
	"encoding/json"
	"github.com/go-playground/validator"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var Config map[string]string
var Validate validator.Validate

func SetConfig() {
	Config = ReadConfig("config.yml")
}

func GetDomainConfig() string {
	return Config["cloudfoundry_domain"]
}

func ReadConfig(filename string) map[string]string {
	Config = map[string]string{}
	if len(filename) == 0 {
		return Config
	}
	file, err := os.Open(filename)

	if err != nil {
		return nil
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if equal := strings.Index(line, ":"); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				Config[key] = value
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil
		}
	}
	return Config
}

//객체 유효성 체크 {
func ValidateConfig() {
	Validate = *validator.New()
}

func ValidationsCheck(s interface{}) (*Error, error) {
	err := Validate.Struct(s)
	if err != nil {
		errM := ""
		for _, err := range err.(validator.ValidationErrors) {
			errM += "'" + err.StructNamespace() + "' Error:Field validation for '" + err.Field() + "'ailed on the ‘required’\n"
		}
		rErrs := &Errors{Code: 500, Detail: errM, Title: "Portal Validation Error"}
		var rErr Error
		rErr.Errors = append(rErr.Errors, *rErrs)
		return &rErr, nil
	}
	return nil, err
}

func Validation(r *http.Request, value interface{}) (interface{}, bool) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &value)
	result, err := ValidationsCheck(value)
	if result != nil {
		return result, false
	} else if err != nil {
		return err, false
	}
	return value, true
}

// }

func ErrorMessage(msg string, code int, w http.ResponseWriter) interface{} {
	w.WriteHeader(code)
	rErrs := &Errors{Code: code, Detail: msg, Title: "Portal API Error"}
	return rErrs
}

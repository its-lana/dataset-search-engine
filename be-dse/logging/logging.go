package logging

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
	"strconv"

	"be-dse/datastruct"
)


var data datastruct.Datasets

var dataFields = GetDataFields(data)

func GetDateTimeNowInString() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

func ValidateAndReturnSortQuery(sortBy string) (string, error) {
	splits := strings.Split(sortBy, ".")
	if len(splits) != 2 {
		return "", errors.New("malformed sortBy query parameter, should be field.orderdirection")
	}

	field, order := splits[0], splits[1]

	if order != "desc" && order != "asc" {
		return "", errors.New("malformed orderdirection in sortBy query parameter, should be asc or desc")
	}

	if !StringInSlice(dataFields, field) {
		return "", errors.New("unknown field in sortBy query parameter")
	}

	return fmt.Sprintf("%s %s", field, strings.ToUpper(order)), nil
}


func ValidateAndReturnFilterMap(filter string) (string, error) {
    splits := strings.Split(filter, ".")
	if len(splits) != 2 {
		return "", errors.New("malformed filter query parameter, should be field.value")
	}

	field, value := splits[0], splits[1]

	if !StringInSlice(dataFields, field) {
		return "", errors.New("unknown field in filter query parameter")
	}

	return fmt.Sprintf("%s = '%s'", field, value), nil
}

func ValidateAndReturnTahun(filter string) (int64, int64, error) {
    splits := strings.Split(filter, ".")
	if len(splits) != 2 {
		return 0, 0, errors.New("malformed filter query parameter, should be tahunAwal.tahunAkhir")
	}

	tahunAwal_, tahunAkhir_ := splits[0], splits[1]

	tahunAwal, err := strconv.Atoi(tahunAwal_)
	tahunAkhir, err := strconv.Atoi(tahunAkhir_)

	if err != nil {
		return 0,0, errors.New("unknown field in filter query parameter")
	}

	return int64(tahunAwal), int64(tahunAkhir), nil
}

const LOG_PREFIX_LOG string = "LOG"
const ERROR_PREFIX_LOG string = "ERR"
const DEBUG_PREFIX_LOG string = "DEBUG"

const DEFAULT_LOG_FILE_PATH string = "aph-service.log"

var logFileName string = DEFAULT_LOG_FILE_PATH


func GetDataFields(datasets datastruct.Datasets ) []string {
    var field []string
    v := reflect.ValueOf(datasets)
    for i := 0; i < v.Type().NumField(); i++ {
        field = append(field, v.Type().Field(i).Tag.Get("json"))
    }
    return field
}
func StringInSlice(strSlice []string, s string) bool {
    for _, v := range strSlice {
        if v == s {
            return true
        }
    }
    return false
}


func SetLogFileName(filename string) {
	logFileName = filename
}

func Debug(msg string) {
	internalLog(msg, DEBUG_PREFIX_LOG)
}

func Log(msg string) {
	internalLog(msg, LOG_PREFIX_LOG)
}

func Error(msg string) {
	internalLog(msg, LOG_PREFIX_LOG)
}

func internalLog(msg string, prefix string) {
	// prepare the message
	output_msg := fmt.Sprintf("[%s] [%s] %s", prefix, GetDateTimeNowInString(), msg)

	// print to screen and append to log file
	fmt.Println(output_msg)
	appendToLogFile(output_msg, prefix)
}

func appendToLogFile(output_msg string, prefix string) {
	// append log to file
	f, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintln(f, output_msg)
}
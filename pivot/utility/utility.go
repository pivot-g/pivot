package utility

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pivot-g/pivot/pivot/log"
)

func ConvMapInterface(i map[interface{}]interface{}) map[string]interface{} {
	out := map[string]interface{}{}
	for key, value := range i {
		out[fmt.Sprintf("%v", key)] = value
	}

	return out
}

func GetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal("Creating default config")
	}
	return val
}

func FileExist(dir string) bool {
	o := true
	_, err := ioutil.ReadDir(dir)
	if err != nil {
		o = false
	}
	return o
}
func Mkdir(dir string, prem fs.FileMode) {
	os.Mkdir(dir, 0750)
}

func WriteFile(file string, data []byte, perm fs.FileMode) bool {
	status := true
	if perm == 0 {
		perm = 0750
	}
	err := ioutil.WriteFile(file, data, perm)
	if err != nil {
		status = false
		log.Debug("Unable to write file", file)
	}
	return status

}

func In(list []string, str string) bool {
	status := false
	for _, l := range list {
		if str == l {
			status = true
			break
		}
	}
	return status
}

func ListFiles(dir *string) []string {
	files, err := ioutil.ReadDir(*dir)
	var file_list []string

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		log.Debug(file.Name(), file.IsDir())
		if !file.IsDir() {
			file_list = append(file_list, file.Name())

		}
	}
	return file_list
}

func ListYamlandJsonFile(dir *string) []string {
	file_list := ListFiles(dir)
	extension := []string{"json", "yml", "yaml"}
	file_out := []string{}
	for _, file := range file_list {
		split := strings.Split(file, ".")
		if In(extension, strings.ToLower(split[len(split)-1])) {
			file_out = append(file_out, *dir+"/"+file)
		}
	}
	return file_out
}
func ReadYamlandJsonFile(dir *string) *map[string][]byte {
	out := map[string][]byte{}
	for _, file := range ListYamlandJsonFile(dir) {
		yfile, _ := ioutil.ReadFile(file)
		out[file] = yfile
	}
	return &out

}

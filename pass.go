package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"github.com/joho/godotenv"
)

var password_store_dir string

func initialize() {

	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }
	password_store_dir = os.ExpandEnv(os.Getenv("PASSWORD_STORE_DIR"))
	if _, err := os.Stat(password_store_dir); err != nil {
		fmt.Println("passwordStore didn't exist, initializing")
		os.MkdirAll(password_store_dir, 0777)
	}
}

func initPassStoreDir(path string) (string, error) {
	//if path is empty, set it to the current directory
	var _path string
	var err error
	if path == "" {
		_path,err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}else{
		if _,err = os.Stat(path); err != nil {
			log.Fatal(err)
		}
		_path = path
	}
	os.Setenv("PASSWORD_STORE_DIR", _path)
	fmt.Printf("passwordStore initialized to %s", os.Getenv("PASSWORD_STORE_DIR"))
	return _path, err
}

func listPassStore() {
	err := filepath.Walk(password_store_dir, func(path string, info os.FileInfo, err error) error {
		// Indent based on the depth of the file/directory
		indent := strings.Repeat("  ", strings.Count(path, "/"))

		// Print the file/directory name
		fmt.Println(indent + info.Name())

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func helper(option string) {
	switch option {
	case "init":
		fmt.Println("Usage: pass init <path/to/directory>\n" +
			"    Initializes the password store at the specified directory\n" +
			"    If no directory is specified, the current directory is used\n")
	case "list":
		fmt.Println("Usage: pass list\n" +
			"    Lists all the passwords in the password store\n")
	case "insert":
		fmt.Println("Usage: pass insert <path/to/secret>\n" +
			"    Inserts a password at the specified path\n")
		default:
		fmt.Println("Usage: pass <command> \n" +
			"Available commands:\n\n" +
			"    pass init <path/to/directory>\n" +
			"    pass list\n" +
			"    pass insert <path/to/secret>\n" +
			"    pass help <command>\n" +
			" Environment Variables:\n" +
			"    PASSWORD_STORE_DIR = directory where the password store is defined\n")
		}

}
func insertPass(path string) {
	nPath := password_store_dir + "/" + path
	os.MkdirAll(nPath, 0777)
	fmt.Println("inserting password at path: " + nPath)
}
func main() {
	initialize()
	var command string
	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	switch command {
	case "init":
		initPassStoreDir(os.Args[2])
	case "list":
		listPassStore()
	case "help":
		helper(os.Args[2])
	case "insert":
		insertPass(os.Args[2])
	default:
		helper("")
	}
}

//TODO: create a config file

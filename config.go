package main

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
)

// State is the current state of the client
var State Session

//GetConfig retrieves configuration file from ~./config/discord-cli, if it doesn't exist it calls CreateConfig()
func GetConfig() {
	//Get User
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	//Get File
	file, err := os.Open(usr.HomeDir + "/.config/discord-cli/config.json")
	if err != nil {
		log.Println(err)
		CreateConfig()
		log.Fatalln("Created new config file, please edit contents!")
	}

	//Decode File
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&State)
	if err != nil {
		log.Println("Failed to decode configuration file")
		log.Fatalf("Error: %s", err)
	}
}

//CreateConfig creates folder inside $HOME and makes a new empty configuration file
func CreateConfig() {
	//Get User
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	var EmptyStruct Session

	//Create Folder
	err = os.MkdirAll(usr.HomeDir+"/.config/discord-cli/", os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	//Create File
	file, err := os.Create(usr.HomeDir + "/.config/discord-cli/config.json")
	if err != nil {
		log.Fatalln(err)
	}

	//Marshall EmptyStruct
	raw, err := json.Marshal(EmptyStruct)
	if err != nil {
		log.Fatalln(err)
	}

	//PrintToFile
	_, err = file.Write(raw)
	if err != nil {
		log.Fatalln(err)
	}

	file.Close()
}

//CheckState checks the current state for essential missing information, errors will fail the program
func CheckState() {
	//Get User
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if State.Username == "" {
		log.Fatalln("No Username Specified, please edit " + usr.HomeDir + "/.config/discord-cli/config.json")
	}

	if State.Password == "" {
		log.Fatalln("No Password Specified, please edit " + usr.HomeDir + "/.config/discord-cli/config.json")
	}

}

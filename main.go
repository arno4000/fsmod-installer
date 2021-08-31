package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mholt/archiver/v3"
)

func main() {
	downloadFolder := os.Getenv("USERPROFILE") + "\\Downloads\\fsmods"
	if _, err := os.Stat(downloadFolder); os.IsNotExist(err) {
		err = os.Mkdir(downloadFolder, 0777)
		if err != nil {
			log.Fatalln("Failed to create folder at", downloadFolder, err)
		}
	}
	mods, err := ioutil.ReadDir(downloadFolder)
	if err != nil {
		log.Fatalln("Failed to get mods", err)
	}
	if len(mods) == 0 {
		fmt.Println("Bitte Mods aus Zip i", downloadFolder, "tue und nomau usfüehre")
	}
	for _, mod := range mods {
		modFile := downloadFolder + "\\" + mod.Name()
		installMethod := checkInstallMethod()
		if installMethod == "ms" {
			modFolder := os.Getenv("USERPROFILE") + "\\AppData\\Local\\Packages\\Microsoft.FlightSimulator_8wekyb3d8bbwe\\LocalCache\\Packages\\Community"
			err := archiver.Unarchive(modFile, modFolder)
			if err != nil {
				log.Fatalln(err)
			}

		} else if installMethod == "steam" {
			modFolder := os.Getenv("USERPROFILE") + "\\AppData\\Local\\Packages\\Microsoft.FlightDashboard_8wekyb3d8bbwe\\LocalCache\\Packages\\Community"
			err := archiver.Unarchive(modFile, modFolder)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func checkInstallMethod() string {
	var installMethod string
	if _, err := os.Stat(os.Getenv("USERPROFILE") + "\\AppData\\Local\\Packages\\Microsoft.FlightSimulator_8wekyb3d8bbwe"); !os.IsNotExist(err) {
		installMethod = "ms"
	} else {
		installMethod = "steam"
	}
	return installMethod
}

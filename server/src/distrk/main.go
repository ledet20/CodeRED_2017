package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	// "time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/mqtt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"distrk/db"
)

type PatientJSON struct {
	ID       uint
	Name     string
	Consults []ConsultJSON
}

type ConsultJSON struct {
	ID        uint
	Desc      string
	Location  string
	Priority  int
	Completed bool
}

func main() {
	db.Migrate()

	// Connects the the public PAHO MQTT broker.
	mqttAdaptor := mqtt.NewAdaptor("tcp://iot.eclipse.org:1883", "distrk_server")

	// Work thread.
	work := func() {
		mqttAdaptor.On("distrk/new", func(msg mqtt.Message) {
			id, err := strconv.Atoi(string(msg.Payload()))
			if err == nil {
				patientJSON, err := getPatientJSON(id)
				if err == nil {
					parsed, err := json.Marshal(patientJSON)
					if err == nil {
						fmt.Println("Sent")
						mqttAdaptor.Publish("distrk/patient", parsed)
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
					mqttAdaptor.Publish("distrk/patient/not-found", []byte{})
				}
			} else {
				fmt.Println("Invalid payload \"", id, "\"")
			}
		})

		// gobot.Every(1*time.Second, func() {
		// 	patientJSON, err := getPatientJSON(1)
		// 	if err == nil {
		// 		parsed, err := json.Marshal(patientJSON)
		// 		if err == nil {
		// 			fmt.Println("Publishing...")
		// 			mqttAdaptor.Publish("distrk/patient", parsed)
		// 		}
		// 	}
		// })
	}

	// Starts the listener.
	robot := gobot.NewRobot(
		"mqttRobot",
		[]gobot.Connection{mqttAdaptor},
		work,
	)

	robot.Start()
}

func getPatientJSON(id int) (PatientJSON, error) {
	query, err := gorm.Open("sqlite3", "/tmp/distrk.db")

	if err == nil {
		var patient db.Patient
		query.Preload("Consults").Preload("Consults.Department").First(&patient, id)

		if patient.Model.ID != 0 {
			consultsJSON := make([]ConsultJSON, 0)
			for i := 0; i < len(patient.Consults); i++ {
				consultsJSON = append(consultsJSON, ConsultJSON{
					patient.Consults[i].Model.ID,
					patient.Consults[i].Department.Name,
					patient.Consults[i].Department.Location,
					patient.Consults[i].Department.Priority,
					patient.Consults[i].Completed,
				})
			}
			return PatientJSON{
				patient.Model.ID,
				patient.Name,
				consultsJSON,
			}, nil
		}

		return PatientJSON{}, errors.New("Patient not found")
	}

	return PatientJSON{}, err
}

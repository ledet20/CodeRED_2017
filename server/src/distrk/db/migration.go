package db

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Migrate() {
	db, err := gorm.Open("sqlite3", "/tmp/distrk.db")

	if err == nil {
		db.DropTableIfExists(
			"admission_types_departments",
			"patients_consults",
			&Consult{},
			&Department{},
			&Patient{},
			&AdmissionType{},
		)

		db.AutoMigrate(
			&AdmissionType{},
			&Patient{},
			&Department{},
			&Consult{},
		)

		loadData(db)
	}

	defer db.Close()
}

func loadData(db *gorm.DB) {
	db.Exec("INSERT INTO admission_types VALUES (?, ?, ?, ?, ?)", 1, time.Now(), time.Now(), nil, "Consultation")
	db.Exec("INSERT INTO admission_types VALUES (?, ?, ?, ?, ?)", 2, time.Now(), time.Now(), nil, "Surgery")
	db.Exec("INSERT INTO admission_types VALUES (?, ?, ?, ?, ?)", 3, time.Now(), time.Now(), nil, "Transfer")
	db.Exec("INSERT INTO admission_types VALUES (?, ?, ?, ?, ?)", 4, time.Now(), time.Now(), nil, "Delivery")

	db.Exec("INSERT INTO departments VALUES (?, ?, ?, ?, ?, ?, ?)", 1, time.Now(), time.Now(), nil, "Medication", "1st Floor", 0)
	db.Exec("INSERT INTO departments VALUES (?, ?, ?, ?, ?, ?, ?)", 2, time.Now(), time.Now(), nil, "Recheck", "3rd Floor", 1)
	db.Exec("INSERT INTO departments VALUES (?, ?, ?, ?, ?, ?, ?)", 3, time.Now(), time.Now(), nil, "Accomodation", "2nd Floor", 2)
	db.Exec("INSERT INTO departments VALUES (?, ?, ?, ?, ?, ?, ?)", 4, time.Now(), time.Now(), nil, "Transportation", "1st Floor", 10)

	db.Exec("INSERT INTO patients VALUES(?, ?, ?, ?, ?, ?)", 1, time.Now(), time.Now(), nil, 1, "John Doe")
	db.Exec("INSERT INTO patients VALUES(?, ?, ?, ?, ?, ?)", 2, time.Now(), time.Now(), nil, 2, "Jane Doe")

	db.Exec("INSERT INTO consults VALUES(?, ?, ?, ?, ?, ?)", 1, time.Now(), time.Now(), nil, 1, false)
	db.Exec("INSERT INTO consults VALUES(?, ?, ?, ?, ?, ?)", 2, time.Now(), time.Now(), nil, 2, false)
	db.Exec("INSERT INTO consults VALUES(?, ?, ?, ?, ?, ?)", 3, time.Now(), time.Now(), nil, 3, true)
	db.Exec("INSERT INTO consults VALUES(?, ?, ?, ?, ?, ?)", 4, time.Now(), time.Now(), nil, 4, false)

	db.Exec("INSERT INTO patients_consults VALUES(?, ?)", 1, 1)
	db.Exec("INSERT INTO patients_consults VALUES(?, ?)", 1, 2)
	db.Exec("INSERT INTO patients_consults VALUES(?, ?)", 2, 3)
	db.Exec("INSERT INTO patients_consults VALUES(?, ?)", 2, 4)
}

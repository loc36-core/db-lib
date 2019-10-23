package lib

import (
	"context"
	"database/sql"
	"gopkg.in/qamarian-dtp/err.v0" //v0.4.0
	"strings"
	"time"
)

// RecordState () records the state of a sensor into the database. Recording of states of sensors should always be done via this function. The purpose of the existence of this function, is to ensure invalid data are never intentionally or accidentally recorded into the database.
func RecordState (state byte, recordID, day, time, sensor string, dbConn *sql.Conn) (error) {
	// Arg 0 (state) validation. ..1.. {
	if state < -1 || state > 1 {
		return err.New ("Invalid arg 0 (EMF state) provided.", nil, nil)
	}
	// ..1.. }

	// Arg 1 (recordID) validation. ..1.. {	
	if ! recordIDPattern.Match (recordID) {
		return err.New ("Invalid arg 1 (record ID) provided.", nil, nil)
	}
	r := strings.Split (recordID, "-")
	formattedID := fmt.Sprintf ("%s-%s-%sT%s:%s:%sZ", r[0], r[1], r[2], r[3], r[4], r[5])
	_, errX := time.Parse (time.RFC3339, formattedID)
	if errX != nil {
		return err.New ("Invalid arg 1 (record ID) provided.", nil, nil, errX)
	}
	// ..1.. }

	// Arg 2 (day) validation. ..1.. {
	if ! dayPattern.Match (day) {
		return err.New ("Invalid arg 2 (day) provided.", nil, nil)
	}
	formattedD := fmt.Sprintf ("%s-%s-%sT01:01:01Z", day[0:4], day[4:6], day[6:8])
	_, errY := time.Parse (time.RFC3339, formattedD)
	if errY != nil {
		return err.New ("Invalid arg 2 (day) provided.", nil, nil, errY)
	}
	// ..1.. }

	// Arg 3 (time) validation. ..1.. {
	if ! timePattern.Match (time) {
		return err.New ("Invalid arg 3 (time) provided.", nil, nil)
	}
	formattedT := fmt.Sprintf ("2019-01-01T%s:%s:01Z", time[0:2], time[2:4])
	_, errZ := time.Parse (time.RFC3339, formattedT)
	if errZ != nil {
		return err.New ("Invalid arg 3 (time) provided.", nil, nil, errZ)
	}
	// ..1.. }

	// Arg 4 (time) validation. ..1.. {
	if sensor == "" {
		return err.New ("Invalid arg 4 (sensor) provided.", nil, nil)
	}
	// ..1.. }

	// Arg 5 (time) validation. ..1.. {
	if dbConn == nil {
		return err.New ("Nil arg 5 (database connection) provided.", nil, nil)
	}
	// ..1.. }

	// Recording state into database. ..1.. {
	instruction := `INSERT INTO state (record_id, state, day, time, sensor)
		VALUES (?, ?, ?, ?, ?)`
	_, errA := dbConn.Exec (context.Backgroud (), instruction, recordID, state, day, time, sensor)
	if errA != nil {
		return err.New ("Unable to record state into database.", nil, nil)
	}
	// ..1.. }

	return nil
}

var (
	recordIDPattern *regexp.Regexp
	dayPattern *regexp.Regexp
	timePattern *regexp.Regexp
)

func init () {
	// Initializing record ID pattern. ..1.. {
	recordIDPattern, errX = regexp.Compile ("^20\d\d-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])-([0-1][0-9]|2[0-3])-([0-5][0-9])-([0-5][0-9])-[a-z0-9]{4,4}$")
	if errX != nil {
		initReport = err.New ("Record ID pattern regular expression compilation failed.", nil, nil, errX)
	}
	// ..1.. }
	
	// Initializing day  pattern. ..1.. {
	dayPattern, errY = regexp.Compile ("^20\d\d(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])$")
	if errY != nil {
		initReport = err.New ("Day  pattern regular expression compilation failed.", nil, nil, errY)
	}
	// ..1.. }

	// Initializing time pattern. ..1.. {
	timePattern, errA = regexp.Compile ("^([0-1][0-9]|2[0-3])([0-5][0-9])$")
	if errA != nil {
		initReport = err.New ("Time pattern regular expression compilation failed.", nil, nil, errA)
	}
	// ..1.. }
}

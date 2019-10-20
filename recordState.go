package lib

import (
	"database/sql"
	"gopkg.in/qamarian-dtp/err.v0" //v0.4.0
	"strings"
	"time"
)

func RecordState (state byte, recordID, day, time, sensor string, dbConn *sql.Conn) (error) {
	// ..1.. {
	if state < -1 || state > 1 {
		return err.New ("Invalid EMF state provided.", nil, nil)
	}
	// ..1.. }

	// ..1.. {	
	if ! recordIDPattern.Match (recordID) {
		return err.New ("Invalid record ID provided.", nil, nil)
	}
	r := strings.Split (recordID, "-")
	formattedID := fmt.Sprintf ("%s-%s-%sT%s:%s:%sZ", r[0], r[1], r[2], r[3], r[4], r[5])
	_, errX := time.Parse (time.RFC3339, formattedID)
	if errX != nil {
		return err.New ("Invalid record ID provided.", nil, nil, errX)
	}
	// ..1.. }

	// ..1.. {
	if ! dayPattern.Match (day) {
		return err.New ("Invalid day provided.", nil, nil)
	}
	formattedD := fmt.Sprintf ("%s-%s-%sT01:01:01Z", day[0:4], day[4:6], day[6:8])
	_, errY := time.Parse (time.RFC3339, formattedD)
	if errY != nil {
		return err.New ("Invalid day provided.", nil, nil, errY)
	}
	// ..1.. }

	// ..1.. {
	if ! timePattern.Match (time) {
		return err.New ("Invalid time provided.", nil, nil)
	}
	formattedT := fmt.Sprintf ("2019-01-01T%s:%s:01Z", time[0:2], time[2:4])
	_, errZ := time.Parse (time.RFC3339, formattedT)
	if errZ != nil {
		return err.New ("Invalid time provided.", nil, nil, errZ)
	}
	// ..1.. }


}

var (
	recordIDPattern *regexp.Regexp
	dayPattern *regexp.Regexp
	timePattern *regexp.Regexp
)

func init () {
	recordIDPattern, errX = regexp.Compile ("20\d\d-(0[1-9]|10|11|12)-(0[1-9]|[1-2][0-9]|30|31)-(0[1-9]|1[0-9]|2[0-4])-(0[1-9]|[1-5][0-9]|60)-(0[1-9]|[1-5][0-9]|60)-\[a-z0-9]{4,4}")
	if errX != nil {
		initReport = err.New ("Record ID pattern regular expression compilation failed.", nil, nil, errX)
	}
	
	dayPattern, errY = regexp.Compile ("20\d\d(0[1-9]|10|11|12)(0[1-9]|[1-2][0-9]|30|31)")
	if errY != nil {
		initReport = err.New ("Day  pattern regular expression compilation failed.", nil, nil, errY)
	}

	timePattern, errA = regexp.Compile ("(0[1-9]|1[0-9]|2[0-4])-(0[1-9]|[1-5][0-9]|60)")
	if errA != nil {
		initReport = err.New ("Time pattern regular expression compilation failed.", nil, nil, errA)
	}
}

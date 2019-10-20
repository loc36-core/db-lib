package lib

import (
	"database/sql"
	"gopkg.in/qamarian-dtp/err.v0" //v0.4.0
	"strings"
	"time"
)

func RecordState (state byte, recordID, day, time, sensor string, dbConn *sql.Conn) () {
	if state < -1 || state > 1 {
		return err.New ("Invalid EMF state provided.", nil, nil)
	}
	
	if ! recordIDPattern.Match (recordID) {
		return err.New ("Invalid record ID provided.", nil, nil)
	}
	t := strings.Split (time, "-")
	formattedT := fmt.Sprintf ("%s-%s-%sT%s:%s:%sZ01:00", t[0], t[1], t[2], t[3], t[4], t[5])
	_, errX := time.Parse (time.RFC3339, formattedT)
	if errX != nil {
		return err.New ("Invalid record ID provided.", nil, nil)
	}
	
	
}

var (
	recordIDPattern *regexp.Regexp
	dayPattern *regexp.Regexp
	timePattern *regexp.Regexp
)

func init () {
	recordIDPattern, errX := regexp.Compile ("20\d\d-(0[1-9]|10|11|12)-(0[1-9]|[1-2][0-9]|30|31)-(0[1-9]|1[0-9]|2[0-4])-(0[1-9]|[1-5][0-9]|60)-(0[1-9]|[1-5][0-9]|60)-\[a-z0-9]{4,4}")
	if errX != nil {
		initReport = err.New ("Record ID pattern regular expression compilation failed.", nil, nil, errX)
	}
	
	dayPattern = recordIDPattern
	
	timePattern, errA := regexp.Compile ("(0[1-9]|1[0-9]|2[0-4])-(0[1-9]|[1-5][0-9]|60)")
	if errA != nil {
		initReport = err.New ("Time pattern regular expression compilation failed.", nil, nil, errA)
	}
}

package domain

import "errors"

type LogLevel string

var ErrLogIsBlank = errors.New("log field is blank")

const (
	LevelInfo    LogLevel = "INFO"
	LevelWarning LogLevel = "WARNING"
	LevelError   LogLevel = "ERROR"
	LevelDebug   LogLevel = "DEBUG"
	LevelFatal   LogLevel = "FATAL"
	LevelUnknown LogLevel = "UNKNOWN"
)

type Log struct {
	Id              string   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Level           LogLevel `json:"level"`
	FromApplication string   `json:"source"`
	DateTime        string   `json:"date_time"`
	Message         string   `json:"message"`
}

func CreateLog(level LogLevel, fromApplication string, dateTime string, message string) (log *Log, err error) {
	// Input validation
	if level == "" || fromApplication == "" || dateTime == "" || message == "" {
		return nil, ErrLogIsBlank
	}

	// Create the log
	log = &Log{
		Level:           level,
		FromApplication: fromApplication,
		DateTime:        dateTime,
		Message:         message,
	}
	return log, nil
}

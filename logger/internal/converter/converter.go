package converter

import (
	"github.com/Kifiya-Financial-Technology/Logger/internal/domain"
	"github.com/Kifiya-Financial-Technology/Logger/loggerpb"
)

func ConvertLogLevelToPb(level domain.LogLevel) loggerpb.LogStatus {
	switch level {
	case domain.LevelDebug:
		return loggerpb.LogStatus_DEBUG
	case domain.LevelInfo:
		return loggerpb.LogStatus_INFO
	case domain.LevelWarning:
		return loggerpb.LogStatus_WARNING
	case domain.LevelError:
		return loggerpb.LogStatus_ERROR
	case domain.LevelFatal:
		return loggerpb.LogStatus_FATAL
	default:
		return loggerpb.LogStatus_UNKNOWN
	}
}
func ConvertLogLevelToDomain(level loggerpb.LogStatus) domain.LogLevel {
	switch level {
	case loggerpb.LogStatus_DEBUG:
		return domain.LevelDebug
	case loggerpb.LogStatus_ERROR:
		return domain.LevelError
	case loggerpb.LogStatus_WARNING:
		return domain.LevelWarning
	case loggerpb.LogStatus_INFO:
		return domain.LevelInfo
	case loggerpb.LogStatus_FATAL:
		return domain.LevelFatal
	default:
		return domain.LevelUnknown
	}
}

func ConvertLoggerToPb(from *domain.Log) (to *loggerpb.Log) {
	to = &loggerpb.Log{
		Id:              from.Id,
		Status:          ConvertLogLevelToPb(from.Level),
		FromApplication: from.FromApplication,
		DateTime:        from.DateTime,
		Message:         from.Message,
	}
	return
}

func ConvertLoggerPbToDomain(from *loggerpb.Log) (to *domain.Log) {
	to = &domain.Log{
		Id:              from.Id,
		Level:           ConvertLogLevelToDomain(from.Status),
		FromApplication: from.FromApplication,
		DateTime:        from.DateTime,
		Message:         from.Message,
	}
	return
}

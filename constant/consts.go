package constant

import (
	"time"
)

const (
	AppVersion string = "1.0.0"

	DbDateFormat      string = "2006-01-02"
	DbTimestampFormat string = "2006-01-02 15:04:05"

	ApiDateFormat      string = "2006-01-02"
	ApiTimestampFormat string = "2006-01-02T15:04:05"

	SessionCookieName string        = "session"
	SessionValidity   time.Duration = 1 * time.Hour

	ContextKeyTransactionHolder string = "transaction-holder"
	ContextKeySessionHolder     string = "session-holder"
	ContextKeySecurityContext   string = "security-context"

	SettingKeyShowOverviewDetails string = "show_overview_details"

	ViewPathDefault string = "/mediaserver"

	ApiPath string = "/mediaserver"

	ConfigFileName string = "config/media-service.yaml"

	CoreServiceMediaTokenDecryptTemplate = "%s://%s:%d/api/v4/internal/media/decrypt/%s"
	CoreServiceTokenHeader               = "X-Sds-Service-Token"
	CoreServiceForwardHostHeader         = "X-Dracoon-Forwarded-Host"
	CoreServiceForwardProtoHeader        = "X-Dracoon-Forwarded-Proto"
	CoreServiceForwardPortHeader         = "X-Dracoon-Forwarded-Port"
)

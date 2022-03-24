package constant

import "time"

const (
	AppVersion string = "1.0.0"

	DbDateFormat      string = "2006-01-02"
	DbTimestampFormat string = "2006-01-02 15:04:05"

	ApiDateFormat      string = "2006-01-02"
	ApiTimestampFormat string = "2006-01-02T15:04:05"

	ContextKeyJWT            = "jwt"
	ContextKeyJWTTokenClaims = "jwt-token-claims"

	ContextKeyDracoonForwardedProtoHeader  = "header-forwarded-proto"
	ContextKeyDracoonForwardedHostHeader   = "header-forwarded-host"
	ContextKeyDracoonForwardedPortHeader   = "header-forwarded-port"
	ContextKeyDracoonClientAddressHeader   = "header-client-address"
	ContextKeyDracoonClientUserAgentHeader = "header-client-user-agent"

	ContextOAuth2      = "token"
	ContextBasicAuth   = "basic"
	ContextAccessToken = "accesstoken"
	ContextAPIKey      = "apikey"

	SettingKeyShowOverviewDetails string = "show_overview_details"

	ApiPath string = "/mediaserver"

	ConfigFileName string = "config/media-service.yaml"

	CoreServiceBasePathTemplate = "%s://%s:%d/api"

	DracoonForwardedProtoHeader  = "X-Dracoon-Forwarded-Proto"
	DracoonForwardedHostHeader   = "X-Dracoon-Forwarded-Host"
	DracoonForwardedPortHeader   = "X-Dracoon-Forwarded-Port"
	DracoonClientAddressHeader   = "X-Dracoon-Client-Address"
	DracoonClientUserAgentHeader = "X-Dracoon-Client-User-Agent"

	RabbitMQExchangeName           = "pc-ziegert.mediaservice"
	RabbitMQWorkerRoutingKey       = "worker"
	RabbitMQMediaServiceRoutingKey = "mediaservice"
	RabbitMQQueueWorkerName        = "pc-ziegert.mediaservice.worker"
	RabbitMQQueueMediaServiceName  = "pc-ziegert.mediaservice.mediaservice"
	RabbitMQImageMessageType       = "mediaService.image"

	S3KeyTemplate = "%d/%v/%v/imgcache/%019d-%dx%d-%t"

	JWTIssuer = "dracoon auth service"

	TenantBackoffDuration        = time.Second * 2
	TenantMaxDuration            = time.Second * 15
	TenantInitialBackoffDuration = time.Second * 1

	ImageInitialExpiration = time.Second * 10
)

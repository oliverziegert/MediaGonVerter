package model

type MediaToken struct {
	nodeId        uint64
	nodesType     string
	nodesPath     string
	tenantDomain  string
	nodesSize     uint64
	token         string
	downloadToken string
}

func NewMediaToken(
	nodeId uint64,
	nodesType string,
	nodesPath string,
	tenantDomain string,
	nodesSize uint64,
	token string,
	downloadToken string) *MediaToken {
	return &MediaToken{
		nodeId:        nodeId,
		nodesType:     nodesType,
		nodesPath:     nodesPath,
		tenantDomain:  tenantDomain,
		nodesSize:     nodesSize,
		token:         token,
		downloadToken: downloadToken,
	}
}

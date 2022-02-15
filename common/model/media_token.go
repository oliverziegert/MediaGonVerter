package model

type MediaToken struct {
	NodeId        uint64
	NodeType      string
	NodePath      string
	TenantDomain  string
	NodeSize      uint64
	Token         string
	DownloadToken string
}

func NewMediaToken(
	nodeId uint64,
	nodeType string,
	nodePath string,
	tenantDomain string,
	nodeSize uint64,
	token string,
	downloadToken string) *MediaToken {
	return &MediaToken{
		NodeId:        nodeId,
		NodeType:      nodeType,
		NodePath:      nodePath,
		TenantDomain:  tenantDomain,
		NodeSize:      nodeSize,
		Token:         token,
		DownloadToken: downloadToken,
	}
}

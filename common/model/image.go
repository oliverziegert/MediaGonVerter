package model

import (
	"encoding/json"
	"fmt"
	error2 "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	"pc-ziegert.de/media_service/service/utils"
	"strconv"
	"strings"
)

type ImageState uint8

func (is ImageState) Uint() uint8 {
	return uint8(is)
}

const (
	ImageStateUnknown ImageState = iota
	ImageStateDownloading
	ImageStateConverting
	ImageStateCached
	ImageStateConversionError
)

type Conversion struct {
	Width       uint64 `json:"width,omitempty"`
	Height      uint64 `json:"height,omitempty"`
	Crop        bool   `json:"crop,omitempty"`
	S3UploadUrl string `json:"s3UploadUrl,omitempty"`
}

type Image struct {
	TenantId      string        `json:"tenantUuid,omitempty"`
	NodeId        uint64        `json:"nodeId,omitempty"`
	ServiceName   string        `json:"serviceName,omitempty"`
	NodeType      string        `json:"nodeType,omitempty"`
	S3DownloadUrl string        `json:"s3DownloadUrl,omitempty"`
	NodeSize      uint64        `json:"nodeSize,omitempty"`
	State         ImageState    `json:"state,omitempty"`
	Conversions   []*Conversion `json:"conversions,omitempty"`
}

func NewImage(tenantId string, nodeId uint64) *Image {
	return &Image{
		TenantId: tenantId,
		NodeId:   nodeId,
		State:    ImageStateUnknown,
	}
}

func (i *Image) GetKeyString() string {
	return fmt.Sprintf("%s:%d", i.TenantId, i.NodeId)
}

func (i *Image) GetValueMap() map[string]string {
	valueMap := make(map[string]string)
	valueMap["ServiceName"] = i.ServiceName
	valueMap["NodeType"] = i.NodeType
	valueMap["S3DownloadUrl"] = i.S3DownloadUrl
	valueMap["NodeSize"] = fmt.Sprintf("%d", i.NodeSize)
	valueMap["State"] = fmt.Sprintf("%d", i.State.Uint())
	for _, conversion := range i.Conversions {
		key := fmt.Sprintf("conversion-%dx%d-%t", conversion.Height, conversion.Width, conversion.Crop)
		value := fmt.Sprintf("%s", conversion.S3UploadUrl)
		valueMap[key] = value
	}
	return valueMap
}

func (i *Image) MapToImage(valueMap map[string]string) error {
	var err error

	i.ServiceName = valueMap["ServiceName"]
	i.NodeType = valueMap["NodeType"]
	i.S3DownloadUrl = valueMap["S3DownloadUrl"]
	if i.NodeSize, err = strconv.ParseUint(valueMap["NodeSize"], 10, 64); err != nil {
		return err
	}
	if i.State, err = stringToImageState(valueMap["State"]); err != nil {
		return err
	}
	for k, v := range valueMap {
		if strings.HasPrefix(k, "conversion-") {
			var conversion = Conversion{}
			conversion.S3UploadUrl = v
			ks := strings.Split(k, "-")
			conversion.Crop = utils.StringToBool(ks[2])
			size := strings.Split(ks[1], "x")
			conversion.Height = utils.StringToUint64(size[0])
			conversion.Width = utils.StringToUint64(size[1])
			i.Conversions = append(i.Conversions, &conversion)
		}
	}
	return nil
}

func (i *Image) JSON() ([]byte, *error2.Error) {
	j, err := json.Marshal(i)
	if err != nil {
		err := error2.WrapError(error2.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return nil, err
	}
	return j, nil
}

func (i *Image) JsonToImage(j []byte) *error2.Error {
	err := json.Unmarshal(j, i)
	if err != nil {
		err := error2.WrapError(error2.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return err
	}
	return nil
}

func stringToImageState(s string) (ImageState, error) {
	i, err := strconv.Atoi(s)
	return ImageState(i), err
}

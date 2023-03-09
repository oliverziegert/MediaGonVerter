package model

import (
	"encoding/json"
	"fmt"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
	u "pc-ziegert.de/media_service/service/utils"
	"strconv"
	"strings"
	"time"
)

type ConversionState uint8

func (is ConversionState) Uint() uint8 {
	return uint8(is)
}

const (
	ConversionStateUnknown ConversionState = iota
	ConversionStateCached
	ConversionStateConversionError
)

type Conversion struct {
	Width       uint64                   `json:"width,omitempty"`
	Height      uint64                   `json:"height,omitempty"`
	Crop        bool                     `json:"crop,omitempty"`
	S3UploadUrl *v4.PresignedHTTPRequest `json:"s3UploadUrl,omitempty"`
	State       ConversionState          `json:"state,omitempty"`
	Expires     *time.Time               `json:"expires,omitempty"`
}

type Image struct {
	TenantUuid    string        `json:"tenantUuid,omitempty"`
	NodeId        uint64        `json:"nodeId,omitempty"`
	ServiceName   string        `json:"serviceName,omitempty"`
	NodeType      string        `json:"nodeType,omitempty"`
	S3DownloadUrl string        `json:"s3DownloadUrl,omitempty"`
	NodeSize      uint64        `json:"nodeSize,omitempty"`
	Conversions   []*Conversion `json:"conversions,omitempty"`
}

func NewConversion(width uint64, height uint64, crop bool) *Conversion {
	e := time.Unix(0, 0)
	return &Conversion{
		Width:       width,
		Height:      height,
		Crop:        crop,
		S3UploadUrl: nil,
		State:       ConversionStateUnknown,
		Expires:     &e,
	}
}

func NewImage(tenantUuid string, nodeId uint64) *Image {
	return &Image{
		TenantUuid: tenantUuid,
		NodeId:     nodeId,
	}
}

func (i *Image) GetKeyString() string {
	return fmt.Sprintf("%s:%d", i.TenantUuid, i.NodeId)
}

func (i *Image) GetValueMap() map[string]string {
	valueMap := make(map[string]string)
	i.GetValueMapForValue(&valueMap, "serviceName")
	i.GetValueMapForValue(&valueMap, "nodeType")
	i.GetValueMapForValue(&valueMap, "nodeSize")
	for _, c := range i.Conversions {
		c.GetValueMap(&valueMap)
	}
	return valueMap
}

func (i *Image) GetValueMapForValue(valueMap *map[string]string, filed string) {
	vm := *valueMap
	switch filed {
	case "serviceName":
		vm["serviceName"] = i.ServiceName
	case "nodeType":
		vm["nodeType"] = i.NodeType
	case "nodeSize":
		vm["nodeSize"] = fmt.Sprintf("%d", i.NodeSize)
	}
}

func (i *Image) MapToImage(valueMap map[string]string) *e.Error {
	var err error

	i.ServiceName = valueMap["serviceName"]
	i.NodeType = valueMap["nodeType"]
	if i.NodeSize, err = strconv.ParseUint(valueMap["nodeSize"], 10, 64); err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		l.Debug(err.StackTrace())
		return err
	}

	i.Conversions = []*Conversion{}
	for k, v := range valueMap {
		if strings.HasPrefix(k, "conversion-") {
			var conversion = Conversion{}
			kt := strings.TrimLeft(k, "conversion-")

			ks := strings.Split(kt, "-")
			if len(ks) != 2 {
				err := e.NewError(e.ValIdInvalid, "Failed to register a consumer.")
				l.Debug(err.StackTrace())
				return err
			}
			conversion.Crop = u.StringToBool(ks[1])

			size := strings.Split(ks[0], "x")
			conversion.Height = u.StringToUint64(size[0])
			conversion.Width = u.StringToUint64(size[1])

			vs := strings.Split(v, "-")
			if len(vs) != 2 {
				err := e.NewError(e.ValIdInvalid, "Failed to register a consumer.")
				l.Debug(err.StackTrace())
				return err
			}

			cs, err := stringToConversionState(vs[0])
			if err != nil {
				err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
				l.Debug(err.StackTrace())
				return err
			}
			conversion.State = *cs

			ce, err := stringToConversionExpires(vs[1])
			if err != nil {
				err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
				l.Debug(err.StackTrace())
				return err
			}
			conversion.Expires = ce

			i.Conversions = append(i.Conversions, &conversion)
		}
	}
	return nil
}

func (i *Image) JSON() ([]byte, *e.Error) {
	j, err := json.Marshal(i)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		l.Debug(err.StackTrace())
		return nil, err
	}
	return j, nil
}

func (i *Image) JsonToImage(j []byte) *e.Error {
	err := json.Unmarshal(j, i)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		l.Debug(err.StackTrace())
		return err
	}
	return nil
}

func (i *Image) GetConversionBySize(width uint64, height uint64, crop bool) (*Conversion, *e.Error) {
	for _, conversion := range i.Conversions {
		if conversion.Crop == crop &&
			conversion.Width == width &&
			conversion.Height == height {
			return conversion, nil
		}
	}
	err := e.NewError(e.ValIdInvalid, "No conversion found")
	l.Debug(err.StackTrace())
	return nil, err
}

func stringToConversionState(s string) (*ConversionState, *e.Error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "", err)
		l.Debug(err.StackTrace())
		return nil, err
	}
	cs := ConversionState(i)
	return &cs, nil
}

func stringToConversionExpires(s string) (*time.Time, *e.Error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "", err)
		l.Debug(err.StackTrace())
		return nil, err
	}
	t := time.Unix(int64(i), 0)
	return &t, nil
}

func (c *Conversion) GetValueMap(valueMap *map[string]string) {
	vm := *valueMap
	key := fmt.Sprintf("conversion-%dx%d-%t", c.Height, c.Width, c.Crop)
	value := fmt.Sprintf("%d-%d", c.State, c.Expires.Unix())
	vm[key] = value
}

package model

import (
	"fmt"
	e "pc-ziegert.de/media_service/common/error"
	l "pc-ziegert.de/media_service/common/log"
	"strconv"
)

type TenantState uint8

func (ts TenantState) UInt8() uint8 {
	return uint8(ts)
}

const (
	TenantStateUnknown TenantState = iota
	TenantStateInvalid
)

type Tenant struct {
	TenantUuid  string      `json:"tenantUuid,omitempty"`
	TenantState TenantState `json:"tenantState,omitempty"`
}

func NewTenant(tenantUuid string) *Tenant {
	return &Tenant{
		TenantUuid:  tenantUuid,
		TenantState: TenantStateUnknown,
	}
}

func (t *Tenant) GetKeyString() string {
	return t.TenantUuid
}

func (t *Tenant) GetValueMap() map[string]string {
	valueMap := make(map[string]string)
	valueMap["tenantState"] = fmt.Sprintf("%d", t.TenantState)
	return valueMap
}

func (t *Tenant) MapToTenant(valueMap map[string]string) *e.Error {
	ts, err := stringToTenantState(valueMap["tenantState"])
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "", err)
		l.Debug(err.StackTrace())
		return err
	}
	t.TenantState = *ts
	return nil
}

func stringToTenantState(s string) (*TenantState, *e.Error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "stringToTenantState failed", err)
		l.Debug(err.StackTrace())
		return nil, err
	}
	ts := TenantState(i)
	return &ts, nil
}

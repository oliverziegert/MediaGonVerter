package utils

import (
	error2 "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	"strconv"
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		err := error2.NewError(error2.ValIdInvalid, "Invalid size variable. ("+s+" cant be converted to int.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return i
}

func StringToUint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		err := error2.NewError(error2.ValIdInvalid, "Invalid size variable. ("+s+" cant be converted to uint64.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return i
}

func StringToBool(s string) bool {
	o, err := strconv.ParseBool(s)
	if err != nil {
		err := error2.NewError(error2.ValIdInvalid, "Invalid string variable. ("+s+" cant be converted to bool.)")
		log.Debug(err.StackTrace())
		panic(err)
	}
	return o
}

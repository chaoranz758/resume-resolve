package pkg

import (
	"errors"
	"strconv"
)

func Int64ToStringList(idList []int64) []string {
	keysString := make([]string, 0, len(idList))

	for i := 0; i < len(idList); i++ {
		keysString = append(keysString, strconv.FormatInt(idList[i], 10))
	}

	return keysString
}

func Int64ToStringListWithPrefix(prefix string, idList []int64) []string {
	keysString := make([]string, 0, len(idList))

	for i := 0; i < len(idList); i++ {
		keysString = append(keysString, prefix+strconv.FormatInt(idList[i], 10))
	}

	return keysString
}

func InterfaceToInt64(idList []interface{}) ([]int64, error) {
	idList1 := make([]int64, 0, len(idList))

	for i := 0; i < len(idList); i++ {
		result, ok := idList[i].(int64)
		if !ok {
			return nil, errors.New(ErrTypeConvert)
		}
		idList1 = append(idList1, result)
	}

	return idList1, nil
}

func InterfaceToStringToInt64(idList []interface{}) ([]int64, error) {
	idList1 := make([]int64, 0, len(idList))

	for i := 0; i < len(idList); i++ {
		result, ok := idList[i].(string)
		if !ok {
			return nil, errors.New(ErrTypeConvert)
		}
		result1, err := strconv.ParseInt(result, 10, 64)
		if err != nil {
			return nil, err
		}
		idList1 = append(idList1, result1)
	}

	return idList1, nil
}

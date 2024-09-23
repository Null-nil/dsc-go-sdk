package swagger

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

// Compare type from swagger declaration and
// real output from endpoint

// this is dump api for request only
type API struct {
	client *resty.Client
}

// NewAPI creates Decimal API instance.
func NewAPI(apiURL string) *API {
	return &API{
		client: resty.New().SetHostURL(apiURL).SetTimeout(time.Minute),
	}
}

type OptionalParams struct {
	Limit  int
	Offset int
}

func (opt *OptionalParams) String() string {
	return fmt.Sprintf("?limit=%d&offset=%d", opt.Limit, opt.Offset)
}

func compareJSON(json1, json2 []byte) []string {
	var v1, v2 interface{}
	err := json.Unmarshal(json1, &v1)
	if err != nil {
		return []string{
			fmt.Sprintf("json1 unmarshal: %s", err.Error()),
		}
	}
	err = json.Unmarshal(json2, &v2)
	if err != nil {
		return []string{
			fmt.Sprintf("json2 unmarshal: %s", err.Error()),
		}
	}
	return switchCompare("", v1, v2)
}

// v1 and v2 must be same type: map[string]interface{} / []interface{}
func switchCompare(key string, v1, v2 interface{}) []string {
	var result []string
	t1, t2 := reflect.TypeOf(v1), reflect.TypeOf(v2)
	if t1 != t2 {
		return []string{
			fmt.Sprintf("types for key '%s' not equal", key),
		}
	}
	switch v1.(type) {
	case map[string]interface{}:
		{
			result = compareMaps(v1.(map[string]interface{}), v2.(map[string]interface{}))
		}
	case []interface{}:
		{
			result = compareArrays(v1.([]interface{}), v2.([]interface{}))
		}
	}

	for i := range result {
		result[i] = fmt.Sprintf("%s. %s", key, result[i])
	}
	return result
}

func compareMaps(map1, map2 map[string]interface{}) []string {
	var result []string
	// common keys for both maps
	var commonKeys = make(map[string]bool)
	// check keys
	for k1 := range map1 {
		k1Exist := false
		for k2 := range map2 {
			if k1 == k2 {
				k1Exist = true
				commonKeys[k1] = true
				break
			}
		}
		if !k1Exist {
			result = append(result, fmt.Sprintf("key %s from swagger not found in response", k1))
		}
	}

	for k2 := range map2 {
		k2Exist := false
		for k1 := range map1 {
			if k1 == k2 {
				k2Exist = true
			}
		}
		if !k2Exist {
			result = append(result, fmt.Sprintf("key %s from response not found in swagger", k2))
		}
	}

	for k := range commonKeys {
		var subresult = switchCompare(k, map1[k], map2[k])
		result = append(result, subresult...)
	}

	return result
}

func compareArrays(arr1, arr2 []interface{}) []string {
	if len(arr1) != len(arr2) {
		return []string{
			fmt.Sprintf("arrays length not equal %d != %d", len(arr1), len(arr2)),
		}
	}
	var result []string
	for i := range arr1 {
		var subresult = switchCompare(strconv.FormatInt(int64(i), 10), arr1[i], arr2[i])
		result = append(result, subresult...)
	}
	return result
}

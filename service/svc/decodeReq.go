package svc

import (
	"fmt"
	"github.com/boram-gong/json-decorator/service/body"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"io/ioutil"

	"net/http"
)

func DecodeNull(c *gin.Context) (interface{}, error) {
	return nil, nil
}

func DecodeTagJsonReq(c *gin.Context) (interface{}, error) {
	reqBody := &body.JsonReq{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}

		switch reqBody.Data.(type) {
		case map[string]interface{}:
			reqBody.JsonMap = reqBody.Data.(map[string]interface{})
		case []interface{}:
			reqBody.JsonSlice = reqBody.Data.([]interface{})
		default:
			return nil, NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
	}
	return reqBody, nil
}

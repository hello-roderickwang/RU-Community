package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"web_app/pkg/snowflake"
	"web_app/settings"

	//"github.com/magiconair/properties/assert"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePostHandler(t *testing.T) {
	fmt.Println("settings.Conf.StartTime", settings.Conf.StartTime)
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)
	body := `{
		"Community_id": 1,
		"title": "test",
		"content": "just a test"
	}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	//assert.Contains(t, w.Body.String(), "Need Login")
	//res := new(ResponseData)
	//if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
	//	t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	//}
	//assert.Equal(t, res.Code, CodeNeedLogin)

}

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func OrderNotify(r *gin.Context) {
	var (
		err  error
		body = make(map[string]interface{})
	)

	err = r.ShouldBind(&body)
	if err != nil {
		r.XML(http.StatusInternalServerError, &map[string]interface{}{
			"return_code": "FAIL",
			"return_msg":  "支付失败",
		})
		return
	}

	log.Println(body)

	r.XML(http.StatusOK, map[string]interface{}{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
	})
	return
}

func main() {
	app := gin.Default()

	app.Any("/", OrderNotify)

	log.Fatal(app.Run())
}

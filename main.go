package main

import (
	"fmt"
	"gee-rewrite/gee"
	"html/template"
	"net/http"
	"time"
)

//
//import (
//	"gee-rewrite/gee"
//	"net/http"
//)
//
//func main() {
//	r := gee.New()
//	//r.GET("/", func(c *gee.Context) {
//	//	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
//	//})
//	//
//	//r.GET("/hello", func(c *gee.Context) {
//	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
//	//})
//	//
//	//r.GET("/hello/:name", func(c *gee.Context) {
//	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.GetParams("name"), c.Path)
//	//})
//	//
//	//r.GET("/assets/*filepath", func(c *gee.Context) {
//	//	c.Json(http.StatusOK, gee.H{"filepath": c.GetParams("filepath")})
//	//})
//	//
//	//fileHandle, err := ioutil.TempFile("./logtmp", "distill-log-test")
//	//if err != nil {
//	//	logger.Errorf("err:%v", err)
//	//}
//	//
//	//fileLogger := logger.NewStreamLogger("test", fileHandle)
//	//fileLogger.Errorf("this is file log")
//
//	//r.GET("/index", func(c *gee.Context) {
//	//	c.HTML(http.StatusOK, "<h1>Index Page</h1>")
//	//})
//
//	r.GET("/ping", func(c *gee.Context) {
//		c.String(200, "pong")
//	})
//
//	r.Use(gee.Logger())
//	//r.GET("/", func(c *gee.Context) {
//	//	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
//	//})
//	v1 := r.Group("/v1")
//	{
//		v1.Use(gee.Logger())
//		//v1.GET("/", func(c *gee.Context) {
//		//	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
//		//})
//
//		v1.GET("/hello", func(c *gee.Context) {
//			// expect /hello?name=geektutu
//			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
//		})
//
//		v1.POST("/login", func(c *gee.Context) {
//			dst := struct {
//				Username string `json:"username"`
//				Password string `json:"password"`
//			}{}
//			err := c.ParseBody(&dst)
//			if err != nil {
//				c.String(http.StatusInternalServerError, err.Error())
//				return
//			}
//
//			c.Json(http.StatusOK, dst)
//		})
//	}
//
//	v2 := r.Group("/v2")
//	{
//		v2.Use(gee.MiddlewareForV2())
//		v2.GET("/hello/:name", func(c *gee.Context) {
//			// expect /hello/geektutu
//			c.String(http.StatusOK, "hello %s, you're at %s\n", c.GetParams("name"), c.Path)
//		})
//
//		v2.POST("/login", func(c *gee.Context) {
//			c.Json(http.StatusOK, gee.H{
//				"username": c.PostForm("username"),
//				"password": c.PostForm("password"),
//			})
//		})
//
//	}
//
//	r.RUN("127.0.0.1:9999")
//}

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	// test
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.RUN(":9999")
}

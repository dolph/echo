package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func getEnv() map[string]string {
	env := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		env[pair[0]] = pair[1]
	}
	return env
}

func getRequest(c *gin.Context) map[string]interface{} {
	return map[string]interface{}{
		"Body":              c.Request.Body,
		"ContentLength":     c.Request.ContentLength,
		"Form":              c.Request.Form,
		"Header":            c.Request.Header,
		"Host":              c.Request.Host,
		"Method":            c.Request.Method,
		"MultipartForm":     c.Request.MultipartForm,
		"PostForm":          c.Request.PostForm,
		"Proto":             c.Request.Proto,
		"RemoteAddr":        c.Request.RemoteAddr,
		"RequestURI":        c.Request.RequestURI,
		"TLS":               c.Request.TLS,
		"Trailer":           c.Request.Trailer,
		"TransferEncodings": c.Request.TransferEncoding,
		"URL":               c.Request.URL,
	}
}

func getOS() map[string]interface{} {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}

	executable, err := os.Executable()
	if err != nil {
		executable = ""
	}

	wd, err := os.Getwd()
	if err != nil {
		wd = ""
	}

	groups, err := os.Getgroups()
	if err != nil {
		groups = nil
	}

	return map[string]interface{}{
		"Hostname":   hostname,
		"EGID":       os.Getegid(),
		"Groups":     groups,
		"Executable": executable,
		"PID":        os.Getpid(),
		"PPID":       os.Getppid(),
		"UID":        os.Getuid(),
		"EUID":       os.Geteuid(),
		"WD":         wd,
	}
}

func setupRouter() *gin.Engine {
	// Respond to all requests using the same func.
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{
			"OS":      getOS(),
			"Request": getRequest(c),
			"Env":     getEnv(),
		})
	})
	return router
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := setupRouter()
	err := router.Run()
	if err != nil {
		log.Fatalf("Error running router: %v", err)
	}
}

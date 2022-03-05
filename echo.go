package main

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func getParams(c *gin.Context) map[string]string {
	params := make(map[string]string)
	for _, param := range c.Params {
		params[param.Key] = param.Value
	}
	return params
}

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

func index(c *gin.Context) {

	c.IndentedJSON(200, gin.H{
		"OS":      getOS(),
		"Request": getRequest(c),
		"Env":     getEnv(),
		"Params":  getParams(c),
	})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.NoRoute(index)
	r.Run()
}

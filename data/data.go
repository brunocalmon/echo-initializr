package data

var files Files
var grave string = "`"

func Initializr(namespace string) Files {
	createFile(File{Namespace: namespace, Name: "README.md", Content: readme})
	createFile(File{Namespace: namespace, Name: "main.go", Content: main})
	createFile(File{Namespace: namespace, Name: "go.mod", Content: gomod})

	createFile(File{Namespace: namespace, Folder: "/utils", Name: "string_utils.go", Content: string_utils})
	createFile(File{Namespace: namespace, Folder: "/domain", Name: "model_error_message.go", Content: model_error_message})
	createFile(File{Namespace: namespace, Folder: "/controller", Name: "health_controler.go", Content: health_controler})
	createFile(File{Namespace: namespace, Folder: "/config", Name: "environment.go", Content: environment})
	createFile(File{Namespace: namespace, Folder: "/appcontext", Name: "context.go", Content: context})

	createFile(File{Namespace: namespace, Folder: "/gateway", Name: ".gitkeep"})
	createFile(File{Namespace: namespace, Folder: "/usecase", Name: ".gitkeep"})

	return files
}

func createFile(f File) {
	files = append(files, f)
}

var main string = `package main

import (
	"%s/config"
	"%s/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	g := e.Group("/%s/v1")
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	g.GET("/health", controller.CheckHealth)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
`
var gomod string = `module %s

go %s`

var string_utils string = `package utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"
)

//StringInSlice checks if a slice contains a specific string
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Md5Hash encodes a string using the MD5 algorithm
func Md5Hash(hashData string) string {
	hash := md5.Sum([]byte(hashData))
	hashStr := hex.EncodeToString(hash[:])
	return hashStr
}

//ValidateInt validates if a string is an valid integer
func ValidateInt(value string, name string) bool {
	intVal, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("ERROR: " + name + " is not an int value: \"" + value + "\" | " + err.Error())
		return false
	}
	if (intVal == 0) && (name == "productId") {
		log.Printf("ProductId 0 not allowed")
		return false
	}
	return true
}
`
var model_error_message string = `package domain

import (
	"encoding/json"
)

//ErrorMessage is a wrapper type to return a JSON object with an error message
type ErrorMessage struct {
	Message string
}

//Bytes returns the ErrorMessage JSON bytes
func (errorMessage *ErrorMessage) Bytes() []byte {
	errorMessageJSON, err := json.Marshal(errorMessage)
	if err != nil {
		return nil
	}

	return errorMessageJSON
}

//GetErrorMessageBytes returns the bytes of the error message
func GetErrorMessageBytes(message string, err error) []byte {
	error := ErrorMessage{Message: message + err.Error()}
	return error.Bytes()
}
`
var health_controler string = `package controller

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

type Health struct {
    Status string "json:'status'"
}

func CheckHealth(c echo.Context) error {
    health := Health{}
    health.Status = "UP"
    return c.JSON(http.StatusOK, health)
}
`
var environment string = `package config

import "os"

var (
	// Port to be listened by application
	Port string
)

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func init() {
	Port = getEnv("PORT", "%d")
}
`
var context string = `package appcontext

//List of consts containing the names of the available componentes in the Application Context - appcontext.Current
const ()

//Component is the Base interface for all Components
type Component interface{}

//ApplicationContext is the type defining a map of Components
type ApplicationContext struct {
	components map[string]Component
}

//Current keeps all components available, initialized in the application startup
var Current ApplicationContext

//Add a component By Name
func (applicationContext *ApplicationContext) Add(componentName string, component Component) {
	applicationContext.components[componentName] = component
}

//Get a component By Name
func (applicationContext *ApplicationContext) Get(componentName string) Component {
	return applicationContext.components[componentName]
}

//Delete a component By Name
func (applicationContext *ApplicationContext) Delete(componentName string) {
	delete(applicationContext.components, componentName)
}

//Count returns the count of components registered
func (applicationContext *ApplicationContext) Count() int {
	return len(applicationContext.components)
}

//CreateApplicationContext creates a new ApplicationContext instance
func CreateApplicationContext() ApplicationContext {
	return ApplicationContext{components: make(map[string]Component)}
}

func init() {
	Current = CreateApplicationContext()
}
`
var readme string = `%s
=======

This is your project and has the objective of being a skeleton for new micro services projects

# What is this?

A project using the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).  
With this skeleton you can create new APIs, where it will have [echo](https://echo.labstack.com/) installed already.

## Basic structure

* ` + grave + `appcontext/context.go` + grave + `: is our ` + grave + `container` + grave + ` of services and object within the project.
* ` + grave + `config/` + grave + `: defines all necessary configuration that your application needs to run.
* ` + grave + `domain/` + grave + `: defines all entities of the project
* ` + grave + `gateway/` + grave + `: defines all connections with the external world, e.g. MongoDB, MySQL
* ` + grave + `usecase/` + grave + `: defines the use cases and business rules of the application

# Build

` + grave + grave + grave + `
go build
` + grave + grave + grave + `
`

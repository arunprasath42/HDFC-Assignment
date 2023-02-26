# HDFC Assignment

# How to run the project

1. Clone the project
2. Run `go mod tidy` to install all the dependencies
3. cd into the project directory and run `go run main.go`
4. The server will start on port 8080


## Prerequisites

- Go 1.18+

![logo](https://github.com/slack-go/slack/blob/master/logo.png?raw=true)
# Table of Contents
1. [Project Structure](#project-structure)
2. [Structure overview](#structure-overview)

# Go Multimodule Workspace Structure
This API is a part of a multimodule workspace structure to ease the development and deployment of the project.

# Project structure
This project follows MVC pattern, except the View. The source code related the web application is present inside `src` folder and follows the folder structure of a Java Spring based application.

# Structure-
    api-|  
        |──config/
        |  ├──config.go
        ├──route.go
        |  ├──routes.go
        ├──src/
        |  ├──controllers
        |  ├──models
        |  ├──service
        ├──utils/
        |  ├──constant
        |  ├──validator
        └──app.yaml
        └──main.go
    ├──go.work
    ├──go.work.sum
    ├──README.md
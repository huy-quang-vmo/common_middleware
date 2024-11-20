# Common Middleware

A collection of reusable middleware utilities for Go applications, designed to integrate seamlessly into various projects, following clean architecture principles.

## Features

- **common middleware**

## Installation

```bash
  go get github.com/huy-quang-vmo/common_middleware
```

## Usage
**Import the Package**
```go
    import "github.com/huy-quang-vmo/common_middleware"
```
**Example: Setting Up Middleware in a Gin Router**
```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/huy-quang-vmo/common_middleware"
)

func main() {
    r := gin.Default()
    checkMaintenanceStatusStore := NewCheckMaintenanceStatusStore()
    maintenanceMiddleware := common_middleware.NewMaintenanceMiddleware(checkMaintenanceStatusStore)

    r.Use(func(c *gin.Context) {
        handler := maintenanceMiddleware.MaintenanceStatus(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            c.Next()
        }))
        handler.ServeHTTP(c.Writer, c.Request)
    })

    r.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })

    r.Run(":8080")
}
```

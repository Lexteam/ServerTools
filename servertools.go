package main

import (
    "github.com/lexteam/servertools/controllers/services"
    "gopkg.in/macaron.v1"
)

func main() {
    // Initialise Macaron
    m := macaron.Classic()

    // Services Routes
    m.Group("services", func() {
        m.Post("update/webhook", services.GetUpdateWebhook)
    })

    // Run ServerTools
    m.Run()
}

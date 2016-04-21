package services

import (
    "net/http"
    "encoding/json"
    "github.com/google/go-github/github"
    "gopkg.in/macaron.v1"
)

func GetUpdateWebhook(ctx *macaron.Context) {
    // Check the webhook was a push event
    if (ctx.Req.Header.Get("X-GitHub-Event") == "push") {
        var res github.PushEvent

        err := json.NewDecoder(ctx.Req.Body).Decode(&res)
        if err != nil {
            ctx.Error(http.StatusInternalServerError, "Failed to decode json!")
            return
        }
    }
}

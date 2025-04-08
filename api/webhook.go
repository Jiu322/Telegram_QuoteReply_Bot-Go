package main

import (
  "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  // 调用主逻辑（需在 main.go 中暴露 handleWebhook）
  handleWebhook(w, r)
}

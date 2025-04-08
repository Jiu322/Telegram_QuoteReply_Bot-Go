package main

   import (
     "net/http"
   )

   func Handler(w http.ResponseWriter, r *http.Request) {
     // 转发请求到主逻辑
     handleWebhook(w, r)
   }

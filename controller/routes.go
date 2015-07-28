package xiidec

import (
    "net/http"
)

func init() {
    http.HandleFunc("/", handler_index)
}

package productcontroller

import (
	"net/http"

	"github.com/jintoples/go-jwt-mux/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]any{
		{
			"id":     1,
			"prduct": "sepatu",
			"stok":   100,
		},
		{
			"id":     2,
			"prduct": "baju",
			"stok":   50,
		},
	}

	helper.ResponseJson(w, http.StatusOK, data)
}

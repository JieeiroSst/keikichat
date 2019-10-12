package toolbox

import (
	"crypto/sha256"
	"net/http"
)

func FuncHash(vars string) string {
	hashInput:=sha256.New()
	input:=hashInput.Sum([]byte(vars))
	return string(input)
}

func HashRender(w http.ResponseWriter,r *http.Request){

}
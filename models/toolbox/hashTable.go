package toolbox

import (
	"crypto/sha256"
	"crypto/md5"
)

//hash lan 1
//bao mat lan 1
func FuncHash(vars string) string {
	hashInput:=sha256.New()
	input:=hashInput.Sum([]byte(vars))
	return string(input)
}

//hash lan 2
//bao mat lan 2
func HashTable(vars string) string {
	hashInput:=md5.New()
	input:=hashInput.Sum([]byte(FuncHash(vars)))
	return string(input)
}

//hash lan 3
//bao mat lan 3
func Hashsession(vars string) string {
	hashInput:=sha256.New()
	input:=hashInput.Sum([]byte(HashTable(vars)))
	return string(input)
}
//hash lan
//bao mat lan 3
func HahsFinal(vars string) string{
	hashInput:=sha256.New()
	input:=hashInput.Sum([]byte(Hashsession(vars)))
	return string(input)
}
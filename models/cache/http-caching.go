package cache

import
(
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/patrickmn/go-cache"
)

var newCache *cache.Cache

func init() {
	newCache = cache.New(5*time.Minute, 10*time.Minute)
	newCache.Set("foo", "bar", cache.DefaultExpiration)
}
func GetFromCache(w http.ResponseWriter, r *http.Request) {
	foo, found := newCache.Get("foo")
	if found {
	log.Print("Key Found in Cache with value as :: ",
	foo.(string))
	fmt.Fprintf(w, "Hello "+foo.(string))
	} else
	{
	log.Print("Key Not Found in Cache :: ", "foo")
	fmt.Fprintf(w, "Key Not Found in Cache")
	}
}

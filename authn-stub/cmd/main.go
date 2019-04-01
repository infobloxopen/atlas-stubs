package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

var userinfoToJWT = map[string]string{
	// AccountID: 1
	"admin1:admin": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SUQiOjF9.GsXyFDDARjXe1t9DPo2LIBKHEal3O7t3vLI3edA7dGU",
	// AccountID: 2
	"admin2:admin": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SUQiOjJ9.BoQj3KNm0PZqKjhYpUGzo-YYb-5IDAzEYILYDGZqZL4",
}

func main() {
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		if token := r.Header.Get("Throw-Token"); token != "" {
			w.Header().Set("Grpc-Metadata-Authorization", "Bearer "+token)
			w.WriteHeader(http.StatusOK)
		}

		if v := userinfoToJWT[r.Header.Get("User-And-Pass")]; v == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
		} else {
			log.Infof("Performed translation: Userinfo[%s] -> JWTToken[%s]", r.Header.Get("Username"), v)

			w.Header().Set("Grpc-Metadata-Authorization", "Token "+v)
			w.WriteHeader(http.StatusOK)
		}

		return
	})

	log.Fatal(http.ListenAndServe(":8040", nil))
}

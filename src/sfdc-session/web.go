package main

import (
    "fmt"
    "net/http"
    "net/url"
    "os"
)

func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/oauth/_callback", handleCallback)
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func index(res http.ResponseWriter, req *http.Request) {
    // TODO state param
    authUrl := 
        "https://" + 
        os.Getenv("SFDC_OAUTH_LOGIN_HOST") + 
        "/services/oauth2/authorize?response_type=code&display=popup" + 
        "&client_id=" + url.QueryEscape(os.Getenv("SFDC_OAUTH_CLIENT_KEY")) + 
        "&redirect_uri=" + url.QueryEscape(os.Getenv("SFDC_OAUTH_CALLBACK_URL"))

    http.Redirect(res, req, authUrl, http.StatusFound)
}

func handleCallback(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "CALLED BACK")
}
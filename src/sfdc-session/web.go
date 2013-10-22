package main

import (
    "fmt"
    "net/http"
    "net/url"
    "os"
    "io/ioutil"
    "encoding/json"
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
        loginHost() + 
        "/services/oauth2/authorize?response_type=code&display=popup" + 
        "&client_id=" + url.QueryEscape(clientKey()) + 
        "&redirect_uri=" + url.QueryEscape(callbackUrl())

    http.Redirect(res, req, authUrl, http.StatusFound)
}

func handleCallback(res http.ResponseWriter, req *http.Request) {
    tokenUrl := "https://" + loginHost() + "/services/oauth2/token"

    v := url.Values{}
    v.Set("code", req.FormValue("code"))
    v.Set("grant_type", "authorization_code")
    v.Set("client_id", clientKey())
    v.Set("client_secret", clientSecret())
    v.Set("redirect_uri", callbackUrl())

    tokenRes, tokenErr := http.PostForm(tokenUrl, v)
    if tokenErr != nil {
        fmt.Println(tokenErr) // TODO: handle error
    }

    defer tokenRes.Body.Close()
    tokenResBody, err := ioutil.ReadAll(tokenRes.Body)
    if err != nil {
        fmt.Println(err)  // TODO: handle error
    }

    var tokenResMap map[string]interface{}

    if err := json.Unmarshal(tokenResBody, &tokenResMap); err != nil {
        fmt.Println(err)  // TODO: handle error
    }

    fmt.Fprintln(res, "export SFDC_INSTANCE_URL='" + tokenResMap["instance_url"].(string) + "'")
    fmt.Fprintln(res, "export SFDC_ACCESS_TOKEN='" + tokenResMap["access_token"].(string) + "'")

    fmt.Fprintln(res, "")
    fmt.Fprintln(res, "curl -H 'X-PrettyPrint: 1' -H \"Authorization: Bearer $SFDC_ACCESS_TOKEN\" $SFDC_INSTANCE_URL/services/data")

    fmt.Fprintln(res, "")
    fmt.Fprintln(res, "sudo gem install restforce")
    fmt.Fprintln(res, "irb")
    fmt.Fprintln(res, "require 'restforce'")
    fmt.Fprintln(res, "client = Restforce.new(:oauth_token => ENV['SFDC_ACCESS_TOKEN'], :instance_url  => ENV['SFDC_INSTANCE_URL'])")
}

func loginHost() string {
    return os.Getenv("SFDC_OAUTH_LOGIN_HOST")
}

func clientKey() string {
    return os.Getenv("SFDC_OAUTH_CLIENT_KEY")
}

func clientSecret() string {
    return os.Getenv("SFDC_OAUTH_CLIENT_SECRET")
}

func callbackUrl() string {
    return os.Getenv("SFDC_OAUTH_CALLBACK_URL")
}

package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
)

func main() {
    oauth := OAuthClient{
        os.Getenv("SFDC_OAUTH_LOGIN_HOST"), 
        os.Getenv("SFDC_OAUTH_CLIENT_KEY"), 
        os.Getenv("SFDC_OAUTH_CLIENT_SECRET"), 
        os.Getenv("SFDC_OAUTH_CALLBACK_URL"),
    }

    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        http.Redirect(res, req, oauth.authUrl(), http.StatusFound)
    })

    http.HandleFunc("/oauth/_callback", func(res http.ResponseWriter, req *http.Request) {
        tokenRes, tokenErr := http.PostForm(oauth.tokenUrl(), oauth.tokenParams(req.FormValue("code")))
        if tokenErr != nil {
            fmt.Println(tokenErr) // TODO: handle error
        }

        defer tokenRes.Body.Close()
        tokenResBody, err := ioutil.ReadAll(tokenRes.Body)
        if err != nil {
            fmt.Println(err) // TODO: handle error
        }

        var tokenResMap map[string]interface{}

        if err := json.Unmarshal(tokenResBody, &tokenResMap); err != nil {
            fmt.Println(err) // TODO: handle error
        }

        fmt.Fprintln(res, output(tokenResMap["instance_url"].(string), tokenResMap["access_token"].(string)))
    })

    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
        panic(err)
    }
}

type OAuthClient struct {
    loginHost string
    clientKey string
    clientSecret string
    callbackUrl string    
}

func (c OAuthClient) authUrl() string {
    return "https://" + c.loginHost +
           "/services/oauth2/authorize?response_type=code&display=popup" +
           "&client_id=" + url.QueryEscape(c.clientKey) +
           "&redirect_uri=" + url.QueryEscape(c.callbackUrl)
}

func (c OAuthClient) tokenUrl() string {
    return "https://" + c.loginHost + "/services/oauth2/token"
}

func (oauth OAuthClient) tokenParams(code string) url.Values {
    p := url.Values{}
    p.Set("code", code)
    p.Set("grant_type", "authorization_code")
    p.Set("client_id", oauth.clientKey)
    p.Set("client_secret", oauth.clientSecret)
    p.Set("redirect_uri", oauth.callbackUrl)
    return p
}

func output(instanceUrl string, accessToken string) string {
    return  "export SFDC_INSTANCE_URL='" + instanceUrl + "'\n" +
            "export SFDC_ACCESS_TOKEN='" +  accessToken + "'\n" +
            "curl -H 'X-PrettyPrint: 1' -H \"Authorization: Bearer $SFDC_ACCESS_TOKEN\" $SFDC_INSTANCE_URL/services/data"
}

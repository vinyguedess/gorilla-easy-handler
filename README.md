[![CIBuild](https://github.com/vinyguedess/gorilla-easy-handler/actions/workflows/ci.yaml/badge.svg)](https://github.com/vinyguedess/gorilla-easy-handler/actions/workflows/ci.yaml)

# Gorilla Easy Handler
Facilitator for handling [gorilla/mux](https://github.com/gorilla/mux) requests.

## Introduction
Basically `http.ResponseWriter` and `http.Request` are encapsulated in two structs:
`geh.Response` and `geh.Request`. Both come with methods that facilitate the implementation
and don't require developer to worry about some things.

### Getting started
First, we create a basic Gorrila/Mux API.

```go
func HomeHandler(request geh.Request, response geh.Response) {
	geh.Status(http.StatusOk).
		Json(map[string]interface{}{
            "hello": "I",
            "am": "okay"
        })
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", geh.Handler(HomeHandler))
}
```

### Request

#### GetParams
Get a URL param value. If not found return an empty string.
```go
func HomeHandler(request geh.Request, response geh.Response) {
	response.Text(fmt.sprintf(
		"My ID is %s", 
		request.GetParams("id"),
    ))
}
```

#### GetQueryString
Get query string from URL. If not found return an empty string.

```go
func HomeHandler(request geh.Request, geh.Response) {
	response.Text(fmt.sprintf(
	    "My ID in query string is %s",
	    request.GetQueryString("id")
    ))
}
```

### GetHeader
Get a header value. If not found return an empty string.

```go
func HomeHandler(request geh.Request, geh.Response) {
	response.Text(fmt.sprintf(
	    "My Content-Type is %s",
	    request.GetHeader("Content-Type")
    ))
}
```

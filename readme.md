# URL shortner in GOlang

A simple URL shortner made using golang and gin router.
The POST and GET requested has been tested on postman.

The microservice runs on [localhost:8080](http://localhost:8080) by default.

### Start:

    git clone
    go run .

### for POST request:

In postman's POST request type:

    [localhost:8080/shorten]

In postman -> Body -> raw , pass the JSON value as:

    {
    "url":"http://www.google.com"
    }

A respone base62 shortcode will be generated

### for GET request:

paste the received shortcode in the browser as:

    [localhost:8080/shortcode]

which redirects to the google's homepage

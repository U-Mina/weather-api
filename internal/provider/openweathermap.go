package provider

import (
	"net/http"
)

// to try the interface feature
// weatherClient will use WeatherProvide Interface!!
type WeatherClient struct {
	Client *http.Client
	APIKey string
}

// constructor Ft, creates new CLIENT obj
// * return PTR to WeatherClient, ie, mem address
func NewWeatherClient(apiKey string) *WeatherClient {
	return &WeatherClient{
		Client: &http.Client{},
		APIKey: apiKey,
	}
}

/*
NOTE:
& The "Address Of" Operator
& operator is placed before a value (like struct WeatherClient{...})
It asks the question, "Where in memory does this value live?"
and it returns that memory address.

&WeatherClient{} → Takes a struct value,
returns its memory address.

* The "Dereference" Operator
The * operator is placed before a pointer.
It does the inverse operation: it takes a memory address
and asks: "What is the value that lives at this address?"
This is called "dereferencing" the pointer.

*myPointerVariable → Takes a memory address,
returns the value stored there.

* symbol is also used when declaring a type.

var p *WeatherClient

func New(...) *WeatherClient

In this context, the * does not mean "get the value."
It means "this variable is of the type
'pointer-to-a-WeatherClient'."
*/

package client


import ( 
    "fmt" 
    "net/http" 
    "time"
)

const SfcpEndpoint = "sfcp.com"


// Client Configuration-
type Client struct {
	SfcpEndpoint      string
	CustomHTTPClient   *http.Client
}


// NewClient Initialize a new Client
func NewClient(sfcp_endpoint string) (*Client, error) {
	fmt.Println("Creating client for")
	// Initialize the client
	client := Client{
		CustomHTTPClient: &http.Client{Timeout: 120 * time.Second},
		// Set Default URLs
		SfcpEndpoint: sfcp_endpoint,
	}

	// Return the client
	return &client, nil
}
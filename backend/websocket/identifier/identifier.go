// package identifier provides functionality to uniquely identify websocket
// clients
package identifier

// Client struct should contain all information about the client connected via
// websocket
type Client struct {
	ClientUUID  string // ClientUUID is a completely unique identifier
	HubName     string // HubName is the name of the hub they are connected to
	DisplayName string // DisplayName is the name the user chose

	// this Client struct may also include things like a jwt token or something
	// later so that we can code in some reconnect functionality
}

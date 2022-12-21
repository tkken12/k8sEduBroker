package common

const (
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	GET    = "GET"

	HTTP  = "http"
	HTTPS = "https"

	HTTP_OK             = 200
	HTTP_FORBIDDEN      = 403
	HTTP_INTERNAL_ERROR = 500

	LABEL_MASTER_ROLE = "node-role.kubernetes.io/control-plane"
	ROLE_WORKER       = "Worker"
	ROLE_MASTER       = "Master"
)

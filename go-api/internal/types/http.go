package types

// HTTPCode represents HTTP status codes used in the API
type HTTPCode int

const (
	HTTPOk                  HTTPCode = 200
	HTTPCreated             HTTPCode = 201
	HTTPAccepted            HTTPCode = 202
	HTTPNoContent           HTTPCode = 204
	HTTPBadRequest          HTTPCode = 400
	HTTPNotFound            HTTPCode = 404
	HTTPTooManyRequests     HTTPCode = 429
	HTTPUnprocessableEntity HTTPCode = 422
	HTTPInternalServerError HTTPCode = 500
	HTTPServiceUnavailable  HTTPCode = 503
)

// ErrorResponse represents API error responses
type ErrorResponse struct {
	Error      string   `json:"error"`
	Message    string   `json:"message"`
	StatusCode HTTPCode `json:"statusCode,omitempty"`
}

// APIResponse represents a generic API response wrapper
type APIResponse[T any] struct {
	Data   *T       `json:"data,omitempty"`
	Error  *string  `json:"error,omitempty"`
	Status HTTPCode `json:"status"`
}

// InstagramPostRequest represents the request for Instagram post data
type InstagramPostRequest struct {
	Shortcode string `json:"shortcode" validate:"required"`
}

// InstagramPostResponse represents the response for Instagram post data
type InstagramPostResponse struct {
	Data IGGraphQLResponseDto `json:"data"`
}

// DownloadProxyRequest represents the download proxy request parameters
type DownloadProxyRequest struct {
	URL      string `form:"url" validate:"required,url"`
	Filename string `form:"filename"`
}

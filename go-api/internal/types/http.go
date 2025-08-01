package types

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

type ErrorResponse struct {
	Error      string   `json:"error"`
	Message    string   `json:"message"`
	StatusCode HTTPCode `json:"statusCode,omitempty"`
}

type APIResponse[T any] struct {
	Data   *T       `json:"data,omitempty"`
	Error  *string  `json:"error,omitempty"`
	Status HTTPCode `json:"status"`
}

type InstagramPostRequest struct {
	Shortcode string `json:"shortcode" validate:"required"`
}

type InstagramPostResponse struct {
	Data IGGraphQLResponseDto `json:"data"`
}

type DownloadProxyRequest struct {
	URL      string `form:"url" validate:"required,url"`
	Filename string `form:"filename"`
}

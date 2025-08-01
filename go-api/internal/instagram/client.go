package instagram

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"instagram-downloader-api/internal/types"
)

const (
	instagramGraphQLURL = "https://www.instagram.com/graphql/query"
	userAgent           = "Mozilla/5.0 (Linux; Android 11; SAMSUNG SM-G973U) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/14.2 Chrome/87.0.4280.141 Mobile Safari/537.36"
)

type Client struct {
	httpClient          *http.Client
	sessionIDs          []string
	mu                  sync.Mutex
	currentSessionIndex int
}

func NewClient(sessionIDs []string) *Client {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	initialIndex := 0
	if len(sessionIDs) > 0 {
		initialIndex = rand.Intn(len(sessionIDs))
	}
	return &Client{
		httpClient:          &http.Client{},
		sessionIDs:          sessionIDs,
		currentSessionIndex: initialIndex,
	}
}

func (c *Client) getNextSessionID() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.sessionIDs) == 0 {
		return ""
	}
	sessionID := c.sessionIDs[c.currentSessionIndex]
	c.currentSessionIndex = (c.currentSessionIndex + 1) % len(c.sessionIDs)
	return sessionID
}

func generateRequestBody(shortcode string) string {
	variables := map[string]interface{}{
		"shortcode":               shortcode,
		"fetch_tagged_user_count": nil,
		"hoisted_comment_id":      nil,
		"hoisted_reply_id":        nil,
	}
	docID := "8845758582119845"
	friendlyName := "PolarisPostActionLoadPostQueryQuery"
	variablesJSON, _ := json.Marshal(variables)
	values := url.Values{}
	values.Set("av", "0")
	values.Set("__d", "www")
	values.Set("__user", "0")
	values.Set("__a", "1")
	values.Set("__req", "b")
	values.Set("__hs", "20183.HYP:instagram_web_pkg.2.1...0")
	values.Set("dpr", "3")
	values.Set("__ccg", "GOOD")
	values.Set("__rev", "1021613311")
	values.Set("__s", "hm5eih:ztapmw:x0losd")
	values.Set("__hsi", "7489787314313612244")
	values.Set("__dyn", "7xeUjG1mxu1syUbFp41twpUnwgU7SbzEdF8aUco2qwJw5ux609vCwjE1EE2Cw8G11wBz81s8hwGxu786a3a1YwBgao6C0Mo2swtUd8-U2zxe2GewGw9a361qw8Xxm16wa-0oa2-azo7u3C2u2J0bS1LwTwKG1pg2fwxyo6O1FwlA3a3zhA6bwIxe6V8aUuwm8jwhU3cyVrDyo")
	values.Set("__csr", "goMJ6MT9Z48KVkIBBvRfqKOkinBtG-FfLaRgG-lZ9Qji9XGexh7VozjHRKq5J6KVqjQdGl2pAFmvK5GWGXyk8h9GA-m6V5yF4UWagnJzazAbZ5osXuFkVeGCHG8GF4l5yp9oOezpo88PAlZ1Pxa5bxGQ7o9VrFbg-8wwxp1G2acxacGVQ00jyoE0ijonyXwfwEnwWwkA2m0dLw3tE1I80hCg8UeU4Ohox0clAhAtsM0iCA9wap4DwhS1fxW0fLhpRB51m13xC3e0h2t2H801HQw1bu02j-")
	values.Set("__comet_req", "7")
	values.Set("lsd", "AVrqPT0gJDo")
	values.Set("jazoest", "2946")
	values.Set("__spin_r", "1021613311")
	values.Set("__spin_b", "trunk")
	values.Set("__spin_t", "1743852001")
	values.Set("__crn", "comet.igweb.PolarisPostRoute")
	values.Set("fb_api_caller_class", "RelayModern")
	values.Set("fb_api_req_friendly_name", friendlyName)
	values.Set("variables", string(variablesJSON))
	values.Set("server_timestamps", "true")
	values.Set("doc_id", docID)
	return values.Encode()
}

func (c *Client) GetPostGraphQL(shortcode string) (*types.IGGraphQLResponseDto, int, error) {
	requestBody := generateRequestBody(shortcode)
	if requestBody == "" {
		return nil, 0, fmt.Errorf("failed to generate request body")
	}
	req, err := http.NewRequest("POST", instagramGraphQLURL, strings.NewReader(requestBody))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-FB-Friendly-Name", "PolarisPostActionLoadPostQueryQuery")
	req.Header.Set("X-BLOKS-VERSION-ID", "0d99de0d13662a50e0958bcb112dd651f70dea02e1859073ab25f8f2a477de96")
	req.Header.Set("X-CSRFToken", "uy8OpI1kndx4oUHjlHaUfu")
	req.Header.Set("X-IG-App-ID", "1217981644879628")
	req.Header.Set("X-FB-LSD", "AVrqPT0gJDo")
	req.Header.Set("X-ASBD-ID", "359341")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Referer", fmt.Sprintf("https://www.instagram.com/p/%s/", shortcode))
	if sessionID := c.getNextSessionID(); sessionID != "" {
		req.Header.Set("Cookie", fmt.Sprintf("sessionid=%s", sessionID))
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to read response body: %w", err)
	}
	var igResponse types.IGGraphQLResponseDto
	if err := json.Unmarshal(body, &igResponse); err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to parse response JSON: %w", err)
	}
	return &igResponse, resp.StatusCode, nil
}

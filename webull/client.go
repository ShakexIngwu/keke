package webull

import (
	"bytes"
	"crypto/md5"
	"io/ioutil"

	"encoding/hex"
	"encoding/json"

	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ShakexIngwu/tradeserver/webull/ws"
	model "github.com/ShakexIngwu/tradeserver/webullmodel"
)

// Endpoints for the Webull API
const (
	QuotesEndpoint         = "https://quoteapi.webull.com/api"
	UserEndpoint           = "https://userapi.webull.com/api"
	BrokerQuotesEndpoint   = "https://quoteapi.webullbroker.com/api"
	BrokerQuotesGWEndpoint = "https://quotes-gw.webullbroker.com/api"
	SecuritiesEndpoint     = "https://securitiesapi.webullbroker.com/api"
	UserBrokerEndpoint     = "https://userapi.webullbroker.com/api"
	PaperTradeEndpoint     = "https://act.webullbroker.com/webull-paper-center/api"
	PaperTradeEndpointV     = "https://act.webullfintech.com/webull-paper-center/api"
	TradeEndpoint          = "https://tradeapi.webulltrade.com/api/trade"
	StockInfoEndpoint      = "https://infoapi.webull.com/api"
)
// ErrAuthExpired signals the user must retrieve a new token
//var ErrAuthExpired = errors.New("Authentication token expired")

// AuthExpiredError returned when token needs to be refreshed
type AuthExpiredError struct {}

func (e *AuthExpiredError) Error() string {
	return fmt.Sprint("Authentication token expired")
}

// Client is a helpful abstraction around some common metadata required for
// API operations.
type Client struct {
	Username       string
	HashedPassword string
	AccountType    model.AccountType
	MFA            string
	UUID           string
	DeviceName     string

	AccessToken           string
	AccessTokenExpiration time.Time
	RefreshToken          string

	TradeToken           string
	TradeTokenExpiration time.Time

	DeviceID string

	httpClient *http.Client
}


// NewClient is a constructor for the Webull-Client client
func NewClient(creds *Credentials) (c *Client, err error) {
	c = &Client{
		httpClient: &http.Client{Timeout: time.Second * 10},
	}
	if creds != nil {
		c.DeviceID = creds.DeviceID
		c.Username = creds.Username
		c.AccountType = creds.AccountType
		if creds.DeviceID == "" {
			c.DeviceID = DefaultDeviceID
		}
		hasher := md5.New()
		hasher.Write([]byte(PasswordSalt + creds.Password))
		c.HashedPassword = hex.EncodeToString(hasher.Sum(nil))
		_, err = c.Token()
		if err != nil {
			return nil, err
		}
	}
	return
}

// GetAndDecode retrieves from the endpoint and unmarshals resulting json into
// the provided destination interface, which must be a pointer.
func (c *Client) GetAndDecode(URL url.URL, dest interface{}, headers *map[string]string, urlValues *map[string]string) error {
	if time.Now().After(c.AccessTokenExpiration) {
		return &AuthExpiredError{}
	}
	v := url.Values{}
	if urlValues != nil {
		for key, val := range *urlValues {
			v.Add(key, val)
		}
	}
	URL.RawQuery = v.Encode()

	if req, err := http.NewRequest(http.MethodGet, URL.String(), nil); err != nil {
		return err
	} else if req == nil {
		return fmt.Errorf("unable to create request")
	} else {
		if headers != nil {
			for key, val := range *headers {
				req.Header.Add(key, val)
			}
		}
		return c.DoAndDecode(req, dest)
	}
}

// PostAndDecode retrieves from the endpoint and unmarshals resulting json into
// the provided destination interface, which must be a pointer.
func (c *Client) PostAndDecode(URL url.URL, dest interface{}, headers *map[string]string, urlValues *map[string]string, payload []byte) error {
	if c.AccessToken != "" {
		if time.Now().After(c.AccessTokenExpiration) {
			return &AuthExpiredError{}
		}
	}
	v := url.Values{}
	if urlValues != nil {
		for key, val := range *urlValues {
			v.Set(key, val)
		}
	}
	URL.RawQuery = v.Encode()
	if req, err := http.NewRequest(http.MethodPost, URL.String(), bytes.NewReader(payload)); err != nil {
		return err
	} else if req == nil {
		return fmt.Errorf("unable to create request")
	} else {
		if headers != nil {
			for key, val := range *headers {
				req.Header.Add(key, val)
			}
		}
		return c.DoAndDecode(req, dest)
	}
}

func parseAnything(data []byte) (output interface{}, err error){
	if err = json.Unmarshal(data, &output); err != nil {
		return nil, fmt.Errorf("Unable to marshal body as interface")
	}
	return output, nil
}

// ConnectWebsockets connects to a streaming API by Webull
// NOTE: client still unstable
func (c *Client) ConnectWebsockets(tickerIDs []string) (err error) {
	err = ws.ConnectStreamingQuotes(c.Username, c.HashedPassword, c.DeviceID, c.AccessToken, tickerIDs)
	return err
}

// DoAndDecode provides useful abstractions around common errors and decoding
// issues. Ideally unmarshals into `dest`. On error, it'll use the Webull `ErrorBody` model.
// Last fallback is a plain interface.
func (c *Client) DoAndDecode(req *http.Request, dest interface{}) (err error) {
	var anyBody interface{}
	req.Header.Add("Content-Type", "application/json")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Got read error on body: %s", err.Error())
	}
	if res.StatusCode/100 != 2 {
		b := &bytes.Buffer{}
		var e model.ErrorBody
		err = json.Unmarshal(body, &e)
		if err != nil {
			// anything
			if anyBody, err = parseAnything(body); err != nil {
				return fmt.Errorf("Unable to marshal body as interface")
			}
			dest = anyBody
			return fmt.Errorf("got response %q and could not decode error body %q", res.Status, b.String())
		}
		// anything
		if anyBody, err = parseAnything(body); err != nil {
			return fmt.Errorf("Unable to marshal body as interface")
		}
		dest = anyBody
		return fmt.Errorf(e.Msg)
	}
	if err = json.Unmarshal(body, &dest); err != nil {
		// anything
		if anyBody, err = parseAnything(body); err != nil {
			return fmt.Errorf("Unable to marshal body as interface")
		}
	}
	// anything: uncomment for viewing raw response
	//if anyBody, err := parseAnything(body); err != nil {
	//	return fmt.Errorf("Unable to marshal body as interface")
	//}
	return err
}

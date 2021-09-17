package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

var (
	exportData []ExportData
	usersData,
	rolesData,
	ldapRolesData,
	samlRolesData,
	teamsData,
	samlTeamsData,
	ldapTeamsData,
	samlIDpsData,
	ldapServersData []byte
)

const (
	UsersEndpoint = "/CxRestAPI/auth/Users"
	TeamsEndpoint = "/CxRestAPI/auth/Teams"
	RolesEndpoint = "/CxRestAPI/auth/Roles"

	LdapServersEndpoint            = "/CxRestAPI/auth/LDAPServers"
	LdapRoleMappingsEndpoint       = "/CxRestAPI/auth/LDAPRoleMappings"
	LdapTeamMappingsEndpoint       = "/CxRestAPI/auth/LDAPTeamMappings"
	SamlIdentityProvidersEndpoint  = "/CxRestAPI/auth/SamlIdentityProviders"
	SamlRoleMappingsEndpoint       = "/CxRestAPI/auth/SamlRoleMappings"
	TeamMappingsEndpoint           = "/CxRestAPI/auth/SamlTeamMappings"
	ReportsLastTriagedScanEndpoint = "/CxWebInterface/odata/v1/Results?$select=Id,ScanId,Date,Scan&$expand=Scan($select=ProjectId)&$filter="
	ReportsCheckStatusEndpoint     = "/CxRestAPI/help/reports/sastScan/%d/status"
	ReportsResultEndpoint          = "/CxRestAPI/help/reports/sastScan/%d"
	CreateReportIDEndpoint         = "/CxRestAPI/help/reports/sastScan"
	LastTriagedFilters             = "Date gt %s and Comment ne null"
)

var isDebug bool

type HTTPAdapter interface {
	Do(req *http.Request) (*http.Response, error)
}

type SASTClient struct {
	BaseURL string
	Adapter HTTPAdapter
	Token   *AccessToken
}

type SASTError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func NewSASTClient(baseURL string, adapter HTTPAdapter) (*SASTClient, error) {
	client := SASTClient{
		BaseURL: baseURL,
		Adapter: adapter,
	}
	return &client, nil
}

func (c *SASTClient) Authenticate(username, password string) error {
	req, err := CreateAccessTokenRequest(c.BaseURL, username, password)
	if err != nil {
		return err
	}

	resp, err := c.Adapter.Do(req)
	log.Debug().
		Err(err).
		Str("method", req.Method).
		Str("url", req.URL.String()).
		Int("statusCode", resp.StatusCode).
		Msg("request")
	if err != nil {
		log.Debug().Err(err).Msgf("error authenticating")
		return fmt.Errorf("authentication error")
	}

	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		responseBody, ioErr := ioutil.ReadAll(resp.Body)
		if ioErr != nil {
			return ioErr
		}
		c.Token = &AccessToken{}
		return json.Unmarshal(responseBody, c.Token)
	} else if resp.StatusCode == http.StatusBadRequest {
		defer resp.Body.Close()
		body, ioErr := ioutil.ReadAll(resp.Body)
		if ioErr != nil {
			log.Debug().
				Err(ioErr).
				Str("method", req.Method).
				Str("url", req.URL.String()).
				Int("statusCode", resp.StatusCode).
				Msg("request")
			return fmt.Errorf("error while trying to authenticate: could not read response")
		}
		var response SASTError
		unmarshalErr := json.Unmarshal(body, &response)
		if unmarshalErr != nil {
			log.Debug().
				Err(unmarshalErr).
				Str("method", req.Method).
				Str("url", req.URL.String()).
				Int("statusCode", resp.StatusCode).
				Msg("request")
			return fmt.Errorf("error while trying to authenticate: could not decode response")
		}
		if response.ErrorDescription == "invalid_username_or_password" {
			return fmt.Errorf("error while trying to authenticate: invalid user name or password")
		}
	}
	return fmt.Errorf("error while trying to authenticate")
}

func (c *SASTClient) GetResponseBody(endpoint string) ([]byte, error) {
	req, err := CreateRequest(http.MethodGet, c.BaseURL+endpoint, nil, c.Token)
	if err != nil {
		return []byte{}, err
	}

	resp, err := c.doRequest(req, http.StatusOK)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *SASTClient) PostResponseBody(endpoint string, body io.Reader) ([]byte, error) {
	req, err := CreateRequest(http.MethodPost, c.BaseURL+endpoint, body, c.Token)
	if err != nil {
		return []byte{}, err
	}

	resp, err := c.doRequest(req, http.StatusAccepted)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *SASTClient) doRequest(req *http.Request, expectStatusCode int) (*http.Response, error) {
	resp, err := c.Adapter.Do(req)
	log.Debug().
		Err(err).
		Str("method", req.Method).
		Str("url", req.URL.String()).
		Int("statusCode", resp.StatusCode).
		Msg("request")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != expectStatusCode {
		defer resp.Body.Close()
		return nil, fmt.Errorf("request %s %s failed with status code %d", req.Method, req.URL.String(), resp.StatusCode)
	}
	return resp, nil
}

func (c *SASTClient) GetReportStatusResponse(report ReportResponse) (*StatusResponse, error) {
	statusUnm, errGetStatus := c.GetReportIDStatus(report.ReportID)
	if errGetStatus != nil {
		return &StatusResponse{}, errGetStatus
	}

	var status StatusResponse
	errStatusSheriff := json.Unmarshal(statusUnm, &status)
	if errStatusSheriff != nil {
		return &StatusResponse{}, errStatusSheriff
	}

	return &status, nil
}

func (c *SASTClient) GetUsers() ([]byte, error) {
	return c.GetResponseBody(UsersEndpoint)
}

func (c *SASTClient) GetRoles() ([]byte, error) {
	return c.GetResponseBody(RolesEndpoint)
}

func (c *SASTClient) GetTeams() ([]byte, error) {
	return c.GetResponseBody(TeamsEndpoint)
}

func (c *SASTClient) GetLdapServers() ([]byte, error) {
	return c.GetResponseBody(LdapServersEndpoint)
}

func (c *SASTClient) GetLdapRoleMappings() ([]byte, error) {
	return c.GetResponseBody(LdapRoleMappingsEndpoint)
}

func (c *SASTClient) GetLdapTeamMappings() ([]byte, error) {
	return c.GetResponseBody(LdapTeamMappingsEndpoint)
}

func (c *SASTClient) GetSamlIdentityProviders() ([]byte, error) {
	return c.GetResponseBody(SamlIdentityProvidersEndpoint)
}

func (c *SASTClient) GetSamlRoleMappings() ([]byte, error) {
	return c.GetResponseBody(SamlRoleMappingsEndpoint)
}

func (c *SASTClient) GetSamlTeamMappings() ([]byte, error) {
	return c.GetResponseBody(TeamMappingsEndpoint)
}

func (c *SASTClient) GetTriagedScansFromDate(fromDate string, offset, limit int) ([]byte, error) {
	url := ReportsLastTriagedScanEndpoint
	url += GetEncodingURL(LastTriagedFilters, fromDate)
	url += fmt.Sprintf("&$skip=%d&$top=%d", offset, limit)
	return c.GetResponseBody(url)
}

func (c *SASTClient) GetReportIDStatus(reportID int) ([]byte, error) {
	return c.GetResponseBody(fmt.Sprintf(ReportsCheckStatusEndpoint, reportID))
}

func (c *SASTClient) GetReportResult(reportID int) ([]byte, error) {
	return c.GetResponseBody(fmt.Sprintf(ReportsResultEndpoint, reportID))
}

func (c *SASTClient) PostReportID(body io.Reader) ([]byte, error) {
	return c.PostResponseBody(CreateReportIDEndpoint, body)
}

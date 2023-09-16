package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"cas/jwks"
	"encoding/json"
)
var BaseURL string
func GetBaseURL(c *gin.Context) string {
	scheme := c.GetHeader("X-Forwarded-Proto")
	if scheme == "" {
		scheme = "http"
	}
	return scheme + "://" + c.Request.Host
}
func DiscoveryEndpoint(c *gin.Context) {
	BaseURL =GetBaseURL(c)

	c.IndentedJSON(http.StatusOK, gin.H{
		"issuer":                                 BaseURL,
		"authorization_endpoint":                 BaseURL + "/authorize",
		"token_endpoint":                         BaseURL + "/token",
		"jwks_uri":                               BaseURL + "/jwks",
		"userinfo_endpoint":                      BaseURL + "/userinfo",
		"response_types_supported":               []string{"code"},
		"subject_types_supported":                []string{"public"},
		"id_token_signing_alg_values_supported":  []string{"none"},
	})
}

func JWKSEndpoint(c *gin.Context) {
	jwks, err := jwks.ReadLocalJWKS()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read JWKS from local file"})
		return
	}

	formattedJSON, err := json.MarshalIndent(jwks, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to format JSON"})
		return
	}

	c.Data(http.StatusOK, "application/json", formattedJSON)
}

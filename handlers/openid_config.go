package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"cas/config"
	"cas/jwks"
	"encoding/json"
)

func DiscoveryEndpoint(c *gin.Context) {
	
		c.IndentedJSON(http.StatusOK, gin.H{
			"issuer":                                 config.IssuerURL,
			"authorization_endpoint":                 config.IssuerURL + "/authorize",
			"token_endpoint":                         config.IssuerURL + "/token",
			"jwks_uri":                              config.IssuerURL + "/jwks",
			"userinfo_endpoint":                               config.IssuerURL + "/userinfo",
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
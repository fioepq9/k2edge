/*
 * worker api
 *
 * worker api
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Token struct {
	Metadata *Metadata `json:"metadata"`
	Config *TokenConfig `json:"config"`
	Status *TokenStatus `json:"status"`
}

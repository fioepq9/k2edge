/*
 * 
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type HealthConfig struct {
	Test []string `json:"test"`
	//  Interval is the time to wait between checks.
	Interval int64 `json:"interval"`
	//  Timeout is the time to wait before considering the check to have hung.
	Timeout int64 `json:"timeout"`
	//  The start period for the container to initialize before the retries starts to count down.
	StartPeriod int64 `json:"start_period,omitempty"`
	Retries int32 `json:"retries,omitempty"`
}

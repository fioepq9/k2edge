/*
 * 
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Job struct {
	Metadata *Metadata `json:"metadata"`
	Config *JobConfig `json:"config"`
	Status *JobStatus `json:"status"`
}

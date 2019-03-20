/*
 * Flagr
 *
 * Flagr is a feature flagging, A/B testing and dynamic configuration microservice. The base path for all the APIs is \"/api/v1\". 
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package goflagr

type CreateSegmentRequest struct {
	Description string `json:"description"`
	RolloutPercent int64 `json:"rolloutPercent"`
}

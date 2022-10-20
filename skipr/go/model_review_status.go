/*
 * Protected API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dto

type ReviewStatus string

// List of ReviewStatus
const (
	REVIEWSTATUS_PENDING ReviewStatus = "PENDING"
	REVIEWSTATUS_APPROVED ReviewStatus = "APPROVED"
	REVIEWSTATUS_REFUSED ReviewStatus = "REFUSED"
	REVIEWSTATUS_INFO_REQUIRED ReviewStatus = "INFO_REQUIRED"
)

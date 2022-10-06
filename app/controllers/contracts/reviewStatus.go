/* This file is auto-generated, manual edits in this file will be overwritten! */
package contracts


type ReviewStatus string

const (
REVIEW_STATUS_APPROVED ReviewStatus = "APPROVED"
REVIEW_STATUS_INFO_REQUIRED ReviewStatus = "INFO_REQUIRED"
REVIEW_STATUS_PENDING ReviewStatus = "PENDING"
REVIEW_STATUS_REFUSED ReviewStatus = "REFUSED"
)

var ReviewStatuses = []string{
string(REVIEW_STATUS_APPROVED),
string(REVIEW_STATUS_INFO_REQUIRED),
string(REVIEW_STATUS_PENDING),
string(REVIEW_STATUS_REFUSED),
}
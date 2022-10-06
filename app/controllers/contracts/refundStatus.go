/* This file is auto-generated, manual edits in this file will be overwritten! */
package contracts


type RefundStatus string

const (
REFUND_STATUS_ACCEPTED RefundStatus = "ACCEPTED"
REFUND_STATUS_PENDING RefundStatus = "PENDING"
REFUND_STATUS_REFUSED RefundStatus = "REFUSED"
)

var RefundStatuses = []string{
string(REFUND_STATUS_ACCEPTED),
string(REFUND_STATUS_PENDING),
string(REFUND_STATUS_REFUSED),
}
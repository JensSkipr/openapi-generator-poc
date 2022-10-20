/* This file is auto-generated, manual edits in this file will be overwritten! */
package contracts


type UserRole string

const (
USER_ROLE_ADMIN UserRole = "ADMIN"
USER_ROLE_EMPLOYEE UserRole = "EMPLOYEE"
USER_ROLE_OPERATOR UserRole = "OPERATOR"
USER_ROLE_REVIEWER UserRole = "REVIEWER"
)

var UserRoles = []string{
string(USER_ROLE_ADMIN),
string(USER_ROLE_EMPLOYEE),
string(USER_ROLE_OPERATOR),
string(USER_ROLE_REVIEWER),
}
/* This file is auto-generated, manual edits in this file will be overwritten! */
package contracts


type ExpenseCategory string

const (
EXPENSE_CATEGORY_PRODUCT ExpenseCategory = "PRODUCT"
EXPENSE_CATEGORY_PROVIDER ExpenseCategory = "PROVIDER"
EXPENSE_CATEGORY_SERVICE ExpenseCategory = "SERVICE"
)

var ExpenseCategories = []string{
string(EXPENSE_CATEGORY_PRODUCT),
string(EXPENSE_CATEGORY_PROVIDER),
string(EXPENSE_CATEGORY_SERVICE),
}
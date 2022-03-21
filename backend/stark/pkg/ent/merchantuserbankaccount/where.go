// Code generated by entc, DO NOT EDIT.

package merchantuserbankaccount

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), v))
	})
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), v))
	})
}

// BankName applies equality check predicate on the "bank_name" field. It's identical to BankNameEQ.
func BankName(v int32) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBankName), v))
	})
}

// AccountNumber applies equality check predicate on the "account_number" field. It's identical to AccountNumberEQ.
func AccountNumber(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountNumber), v))
	})
}

// AccountName applies equality check predicate on the "account_name" field. It's identical to AccountNameEQ.
func AccountName(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountName), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MerchantUserBankAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MerchantUserBankAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MerchantUserBankAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MerchantUserBankAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), v))
	})
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedBy), v))
	})
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.MerchantUserBankAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.MerchantUserBankAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedBy), v...))
	})
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedBy), v))
	})
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedBy), v))
	})
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedBy), v))
	})
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedBy), v))
	})
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCreatedBy), v))
	})
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCreatedBy), v))
	})
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCreatedBy), v))
	})
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCreatedBy), v))
	})
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCreatedBy), v))
	})
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.MerchantUserBankAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedBy), v...))
	})
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.MerchantUserBankAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedBy), v...))
	})
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUpdatedBy), v))
	})
}

// BankNameEQ applies the EQ predicate on the "bank_name" field.
func BankNameEQ(v int32) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBankName), v))
	})
}

// BankNameNEQ applies the NEQ predicate on the "bank_name" field.
func BankNameNEQ(v int32) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBankName), v))
	})
}

// BankNameIn applies the In predicate on the "bank_name" field.
func BankNameIn(vs ...int32) predicate.MerchantUserBankAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBankName), v...))
	})
}

// BankNameNotIn applies the NotIn predicate on the "bank_name" field.
func BankNameNotIn(vs ...int32) predicate.MerchantUserBankAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBankName), v...))
	})
}

// BankNameGT applies the GT predicate on the "bank_name" field.
func BankNameGT(v int32) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBankName), v))
	})
}

// BankNameGTE applies the GTE predicate on the "bank_name" field.
func BankNameGTE(v int32) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBankName), v))
	})
}

// BankNameLT applies the LT predicate on the "bank_name" field.
func BankNameLT(v int32) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBankName), v))
	})
}

// BankNameLTE applies the LTE predicate on the "bank_name" field.
func BankNameLTE(v int32) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBankName), v))
	})
}

// AccountNumberEQ applies the EQ predicate on the "account_number" field.
func AccountNumberEQ(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberNEQ applies the NEQ predicate on the "account_number" field.
func AccountNumberNEQ(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberIn applies the In predicate on the "account_number" field.
func AccountNumberIn(vs ...string) predicate.MerchantUserBankAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAccountNumber), v...))
	})
}

// AccountNumberNotIn applies the NotIn predicate on the "account_number" field.
func AccountNumberNotIn(vs ...string) predicate.MerchantUserBankAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAccountNumber), v...))
	})
}

// AccountNumberGT applies the GT predicate on the "account_number" field.
func AccountNumberGT(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberGTE applies the GTE predicate on the "account_number" field.
func AccountNumberGTE(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberLT applies the LT predicate on the "account_number" field.
func AccountNumberLT(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberLTE applies the LTE predicate on the "account_number" field.
func AccountNumberLTE(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberContains applies the Contains predicate on the "account_number" field.
func AccountNumberContains(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberHasPrefix applies the HasPrefix predicate on the "account_number" field.
func AccountNumberHasPrefix(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberHasSuffix applies the HasSuffix predicate on the "account_number" field.
func AccountNumberHasSuffix(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberEqualFold applies the EqualFold predicate on the "account_number" field.
func AccountNumberEqualFold(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberContainsFold applies the ContainsFold predicate on the "account_number" field.
func AccountNumberContainsFold(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccountNumber), v))
	})
}

// AccountNameEQ applies the EQ predicate on the "account_name" field.
func AccountNameEQ(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountName), v))
	})
}

// AccountNameNEQ applies the NEQ predicate on the "account_name" field.
func AccountNameNEQ(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccountName), v))
	})
}

// AccountNameIn applies the In predicate on the "account_name" field.
func AccountNameIn(vs ...string) predicate.MerchantUserBankAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAccountName), v...))
	})
}

// AccountNameNotIn applies the NotIn predicate on the "account_name" field.
func AccountNameNotIn(vs ...string) predicate.MerchantUserBankAccount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAccountName), v...))
	})
}

// AccountNameGT applies the GT predicate on the "account_name" field.
func AccountNameGT(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccountName), v))
	})
}

// AccountNameGTE applies the GTE predicate on the "account_name" field.
func AccountNameGTE(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccountName), v))
	})
}

// AccountNameLT applies the LT predicate on the "account_name" field.
func AccountNameLT(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccountName), v))
	})
}

// AccountNameLTE applies the LTE predicate on the "account_name" field.
func AccountNameLTE(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccountName), v))
	})
}

// AccountNameContains applies the Contains predicate on the "account_name" field.
func AccountNameContains(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccountName), v))
	})
}

// AccountNameHasPrefix applies the HasPrefix predicate on the "account_name" field.
func AccountNameHasPrefix(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccountName), v))
	})
}

// AccountNameHasSuffix applies the HasSuffix predicate on the "account_name" field.
func AccountNameHasSuffix(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccountName), v))
	})
}

// AccountNameEqualFold applies the EqualFold predicate on the "account_name" field.
func AccountNameEqualFold(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccountName), v))
	})
}

// AccountNameContainsFold applies the ContainsFold predicate on the "account_name" field.
func AccountNameContainsFold(v string) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccountName), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MerchantUserBankAccount) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MerchantUserBankAccount) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MerchantUserBankAccount) predicate.MerchantUserBankAccount {
	return predicate.MerchantUserBankAccount(func(s *sql.Selector) {
		p(s.Not())
	})
}

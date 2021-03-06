// Code generated by entc, DO NOT EDIT.

package paymenttelcodetail

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"gitlab.com/mcuc/monorepo/backend/stark/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), v))
	})
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), v))
	})
}

// TelcoName applies equality check predicate on the "telco_name" field. It's identical to TelcoNameEQ.
func TelcoName(v int32) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelcoName), v))
	})
}

// SerialNumber applies equality check predicate on the "serial_number" field. It's identical to SerialNumberEQ.
func SerialNumber(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSerialNumber), v))
	})
}

// CardID applies equality check predicate on the "card_id" field. It's identical to CardIDEQ.
func CardID(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCardID), v))
	})
}

// ChargedAmount applies equality check predicate on the "charged_amount" field. It's identical to ChargedAmountEQ.
func ChargedAmount(v uint64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChargedAmount), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v uint64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedBy), v))
	})
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedBy), v))
	})
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
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
func CreatedByNotIn(vs ...string) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
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
func CreatedByGT(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedBy), v))
	})
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedBy), v))
	})
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedBy), v))
	})
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedBy), v))
	})
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCreatedBy), v))
	})
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCreatedBy), v))
	})
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCreatedBy), v))
	})
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCreatedBy), v))
	})
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCreatedBy), v))
	})
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
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
func UpdatedByNotIn(vs ...string) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
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
func UpdatedByGT(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUpdatedBy), v))
	})
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUpdatedBy), v))
	})
}

// TelcoNameEQ applies the EQ predicate on the "telco_name" field.
func TelcoNameEQ(v int32) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelcoName), v))
	})
}

// TelcoNameNEQ applies the NEQ predicate on the "telco_name" field.
func TelcoNameNEQ(v int32) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTelcoName), v))
	})
}

// TelcoNameIn applies the In predicate on the "telco_name" field.
func TelcoNameIn(vs ...int32) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTelcoName), v...))
	})
}

// TelcoNameNotIn applies the NotIn predicate on the "telco_name" field.
func TelcoNameNotIn(vs ...int32) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTelcoName), v...))
	})
}

// TelcoNameGT applies the GT predicate on the "telco_name" field.
func TelcoNameGT(v int32) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTelcoName), v))
	})
}

// TelcoNameGTE applies the GTE predicate on the "telco_name" field.
func TelcoNameGTE(v int32) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTelcoName), v))
	})
}

// TelcoNameLT applies the LT predicate on the "telco_name" field.
func TelcoNameLT(v int32) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTelcoName), v))
	})
}

// TelcoNameLTE applies the LTE predicate on the "telco_name" field.
func TelcoNameLTE(v int32) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTelcoName), v))
	})
}

// SerialNumberEQ applies the EQ predicate on the "serial_number" field.
func SerialNumberEQ(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSerialNumber), v))
	})
}

// SerialNumberNEQ applies the NEQ predicate on the "serial_number" field.
func SerialNumberNEQ(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSerialNumber), v))
	})
}

// SerialNumberIn applies the In predicate on the "serial_number" field.
func SerialNumberIn(vs ...string) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSerialNumber), v...))
	})
}

// SerialNumberNotIn applies the NotIn predicate on the "serial_number" field.
func SerialNumberNotIn(vs ...string) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSerialNumber), v...))
	})
}

// SerialNumberGT applies the GT predicate on the "serial_number" field.
func SerialNumberGT(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSerialNumber), v))
	})
}

// SerialNumberGTE applies the GTE predicate on the "serial_number" field.
func SerialNumberGTE(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSerialNumber), v))
	})
}

// SerialNumberLT applies the LT predicate on the "serial_number" field.
func SerialNumberLT(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSerialNumber), v))
	})
}

// SerialNumberLTE applies the LTE predicate on the "serial_number" field.
func SerialNumberLTE(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSerialNumber), v))
	})
}

// SerialNumberContains applies the Contains predicate on the "serial_number" field.
func SerialNumberContains(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSerialNumber), v))
	})
}

// SerialNumberHasPrefix applies the HasPrefix predicate on the "serial_number" field.
func SerialNumberHasPrefix(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSerialNumber), v))
	})
}

// SerialNumberHasSuffix applies the HasSuffix predicate on the "serial_number" field.
func SerialNumberHasSuffix(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSerialNumber), v))
	})
}

// SerialNumberEqualFold applies the EqualFold predicate on the "serial_number" field.
func SerialNumberEqualFold(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSerialNumber), v))
	})
}

// SerialNumberContainsFold applies the ContainsFold predicate on the "serial_number" field.
func SerialNumberContainsFold(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSerialNumber), v))
	})
}

// CardIDEQ applies the EQ predicate on the "card_id" field.
func CardIDEQ(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCardID), v))
	})
}

// CardIDNEQ applies the NEQ predicate on the "card_id" field.
func CardIDNEQ(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCardID), v))
	})
}

// CardIDIn applies the In predicate on the "card_id" field.
func CardIDIn(vs ...string) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCardID), v...))
	})
}

// CardIDNotIn applies the NotIn predicate on the "card_id" field.
func CardIDNotIn(vs ...string) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCardID), v...))
	})
}

// CardIDGT applies the GT predicate on the "card_id" field.
func CardIDGT(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCardID), v))
	})
}

// CardIDGTE applies the GTE predicate on the "card_id" field.
func CardIDGTE(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCardID), v))
	})
}

// CardIDLT applies the LT predicate on the "card_id" field.
func CardIDLT(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCardID), v))
	})
}

// CardIDLTE applies the LTE predicate on the "card_id" field.
func CardIDLTE(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCardID), v))
	})
}

// CardIDContains applies the Contains predicate on the "card_id" field.
func CardIDContains(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCardID), v))
	})
}

// CardIDHasPrefix applies the HasPrefix predicate on the "card_id" field.
func CardIDHasPrefix(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCardID), v))
	})
}

// CardIDHasSuffix applies the HasSuffix predicate on the "card_id" field.
func CardIDHasSuffix(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCardID), v))
	})
}

// CardIDEqualFold applies the EqualFold predicate on the "card_id" field.
func CardIDEqualFold(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCardID), v))
	})
}

// CardIDContainsFold applies the ContainsFold predicate on the "card_id" field.
func CardIDContainsFold(v string) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCardID), v))
	})
}

// ChargedAmountEQ applies the EQ predicate on the "charged_amount" field.
func ChargedAmountEQ(v uint64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChargedAmount), v))
	})
}

// ChargedAmountNEQ applies the NEQ predicate on the "charged_amount" field.
func ChargedAmountNEQ(v uint64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChargedAmount), v))
	})
}

// ChargedAmountIn applies the In predicate on the "charged_amount" field.
func ChargedAmountIn(vs ...uint64) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChargedAmount), v...))
	})
}

// ChargedAmountNotIn applies the NotIn predicate on the "charged_amount" field.
func ChargedAmountNotIn(vs ...uint64) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChargedAmount), v...))
	})
}

// ChargedAmountGT applies the GT predicate on the "charged_amount" field.
func ChargedAmountGT(v uint64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChargedAmount), v))
	})
}

// ChargedAmountGTE applies the GTE predicate on the "charged_amount" field.
func ChargedAmountGTE(v uint64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChargedAmount), v))
	})
}

// ChargedAmountLT applies the LT predicate on the "charged_amount" field.
func ChargedAmountLT(v uint64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChargedAmount), v))
	})
}

// ChargedAmountLTE applies the LTE predicate on the "charged_amount" field.
func ChargedAmountLTE(v uint64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChargedAmount), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v uint64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v uint64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...uint64) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...uint64) predicate.PaymentTelcoDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v uint64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v uint64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v uint64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v uint64) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// HasPayment applies the HasEdge predicate on the "payment" edge.
func HasPayment() predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PaymentTable, PaymentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymentWith applies the HasEdge predicate on the "payment" edge with a given conditions (other predicates).
func HasPaymentWith(preds ...predicate.Payment) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PaymentTable, PaymentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PaymentTelcoDetail) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PaymentTelcoDetail) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
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
func Not(p predicate.PaymentTelcoDetail) predicate.PaymentTelcoDetail {
	return predicate.PaymentTelcoDetail(func(s *sql.Selector) {
		p(s.Not())
	})
}

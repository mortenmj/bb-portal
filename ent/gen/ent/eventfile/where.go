// Code generated by ent, DO NOT EDIT.

package eventfile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.EventFile {
	return predicate.EventFile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.EventFile {
	return predicate.EventFile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.EventFile {
	return predicate.EventFile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.EventFile {
	return predicate.EventFile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.EventFile {
	return predicate.EventFile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.EventFile {
	return predicate.EventFile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.EventFile {
	return predicate.EventFile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.EventFile {
	return predicate.EventFile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.EventFile {
	return predicate.EventFile(sql.FieldLTE(FieldID, id))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEQ(FieldURL, v))
}

// ModTime applies equality check predicate on the "mod_time" field. It's identical to ModTimeEQ.
func ModTime(v time.Time) predicate.EventFile {
	return predicate.EventFile(sql.FieldEQ(FieldModTime, v))
}

// Protocol applies equality check predicate on the "protocol" field. It's identical to ProtocolEQ.
func Protocol(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEQ(FieldProtocol, v))
}

// MimeType applies equality check predicate on the "mime_type" field. It's identical to MimeTypeEQ.
func MimeType(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEQ(FieldMimeType, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEQ(FieldStatus, v))
}

// Reason applies equality check predicate on the "reason" field. It's identical to ReasonEQ.
func Reason(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEQ(FieldReason, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.EventFile {
	return predicate.EventFile(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.EventFile {
	return predicate.EventFile(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldContainsFold(FieldURL, v))
}

// ModTimeEQ applies the EQ predicate on the "mod_time" field.
func ModTimeEQ(v time.Time) predicate.EventFile {
	return predicate.EventFile(sql.FieldEQ(FieldModTime, v))
}

// ModTimeNEQ applies the NEQ predicate on the "mod_time" field.
func ModTimeNEQ(v time.Time) predicate.EventFile {
	return predicate.EventFile(sql.FieldNEQ(FieldModTime, v))
}

// ModTimeIn applies the In predicate on the "mod_time" field.
func ModTimeIn(vs ...time.Time) predicate.EventFile {
	return predicate.EventFile(sql.FieldIn(FieldModTime, vs...))
}

// ModTimeNotIn applies the NotIn predicate on the "mod_time" field.
func ModTimeNotIn(vs ...time.Time) predicate.EventFile {
	return predicate.EventFile(sql.FieldNotIn(FieldModTime, vs...))
}

// ModTimeGT applies the GT predicate on the "mod_time" field.
func ModTimeGT(v time.Time) predicate.EventFile {
	return predicate.EventFile(sql.FieldGT(FieldModTime, v))
}

// ModTimeGTE applies the GTE predicate on the "mod_time" field.
func ModTimeGTE(v time.Time) predicate.EventFile {
	return predicate.EventFile(sql.FieldGTE(FieldModTime, v))
}

// ModTimeLT applies the LT predicate on the "mod_time" field.
func ModTimeLT(v time.Time) predicate.EventFile {
	return predicate.EventFile(sql.FieldLT(FieldModTime, v))
}

// ModTimeLTE applies the LTE predicate on the "mod_time" field.
func ModTimeLTE(v time.Time) predicate.EventFile {
	return predicate.EventFile(sql.FieldLTE(FieldModTime, v))
}

// ProtocolEQ applies the EQ predicate on the "protocol" field.
func ProtocolEQ(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEQ(FieldProtocol, v))
}

// ProtocolNEQ applies the NEQ predicate on the "protocol" field.
func ProtocolNEQ(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldNEQ(FieldProtocol, v))
}

// ProtocolIn applies the In predicate on the "protocol" field.
func ProtocolIn(vs ...string) predicate.EventFile {
	return predicate.EventFile(sql.FieldIn(FieldProtocol, vs...))
}

// ProtocolNotIn applies the NotIn predicate on the "protocol" field.
func ProtocolNotIn(vs ...string) predicate.EventFile {
	return predicate.EventFile(sql.FieldNotIn(FieldProtocol, vs...))
}

// ProtocolGT applies the GT predicate on the "protocol" field.
func ProtocolGT(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldGT(FieldProtocol, v))
}

// ProtocolGTE applies the GTE predicate on the "protocol" field.
func ProtocolGTE(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldGTE(FieldProtocol, v))
}

// ProtocolLT applies the LT predicate on the "protocol" field.
func ProtocolLT(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldLT(FieldProtocol, v))
}

// ProtocolLTE applies the LTE predicate on the "protocol" field.
func ProtocolLTE(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldLTE(FieldProtocol, v))
}

// ProtocolContains applies the Contains predicate on the "protocol" field.
func ProtocolContains(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldContains(FieldProtocol, v))
}

// ProtocolHasPrefix applies the HasPrefix predicate on the "protocol" field.
func ProtocolHasPrefix(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldHasPrefix(FieldProtocol, v))
}

// ProtocolHasSuffix applies the HasSuffix predicate on the "protocol" field.
func ProtocolHasSuffix(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldHasSuffix(FieldProtocol, v))
}

// ProtocolEqualFold applies the EqualFold predicate on the "protocol" field.
func ProtocolEqualFold(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEqualFold(FieldProtocol, v))
}

// ProtocolContainsFold applies the ContainsFold predicate on the "protocol" field.
func ProtocolContainsFold(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldContainsFold(FieldProtocol, v))
}

// MimeTypeEQ applies the EQ predicate on the "mime_type" field.
func MimeTypeEQ(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEQ(FieldMimeType, v))
}

// MimeTypeNEQ applies the NEQ predicate on the "mime_type" field.
func MimeTypeNEQ(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldNEQ(FieldMimeType, v))
}

// MimeTypeIn applies the In predicate on the "mime_type" field.
func MimeTypeIn(vs ...string) predicate.EventFile {
	return predicate.EventFile(sql.FieldIn(FieldMimeType, vs...))
}

// MimeTypeNotIn applies the NotIn predicate on the "mime_type" field.
func MimeTypeNotIn(vs ...string) predicate.EventFile {
	return predicate.EventFile(sql.FieldNotIn(FieldMimeType, vs...))
}

// MimeTypeGT applies the GT predicate on the "mime_type" field.
func MimeTypeGT(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldGT(FieldMimeType, v))
}

// MimeTypeGTE applies the GTE predicate on the "mime_type" field.
func MimeTypeGTE(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldGTE(FieldMimeType, v))
}

// MimeTypeLT applies the LT predicate on the "mime_type" field.
func MimeTypeLT(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldLT(FieldMimeType, v))
}

// MimeTypeLTE applies the LTE predicate on the "mime_type" field.
func MimeTypeLTE(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldLTE(FieldMimeType, v))
}

// MimeTypeContains applies the Contains predicate on the "mime_type" field.
func MimeTypeContains(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldContains(FieldMimeType, v))
}

// MimeTypeHasPrefix applies the HasPrefix predicate on the "mime_type" field.
func MimeTypeHasPrefix(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldHasPrefix(FieldMimeType, v))
}

// MimeTypeHasSuffix applies the HasSuffix predicate on the "mime_type" field.
func MimeTypeHasSuffix(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldHasSuffix(FieldMimeType, v))
}

// MimeTypeEqualFold applies the EqualFold predicate on the "mime_type" field.
func MimeTypeEqualFold(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEqualFold(FieldMimeType, v))
}

// MimeTypeContainsFold applies the ContainsFold predicate on the "mime_type" field.
func MimeTypeContainsFold(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldContainsFold(FieldMimeType, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.EventFile {
	return predicate.EventFile(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.EventFile {
	return predicate.EventFile(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldContainsFold(FieldStatus, v))
}

// ReasonEQ applies the EQ predicate on the "reason" field.
func ReasonEQ(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEQ(FieldReason, v))
}

// ReasonNEQ applies the NEQ predicate on the "reason" field.
func ReasonNEQ(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldNEQ(FieldReason, v))
}

// ReasonIn applies the In predicate on the "reason" field.
func ReasonIn(vs ...string) predicate.EventFile {
	return predicate.EventFile(sql.FieldIn(FieldReason, vs...))
}

// ReasonNotIn applies the NotIn predicate on the "reason" field.
func ReasonNotIn(vs ...string) predicate.EventFile {
	return predicate.EventFile(sql.FieldNotIn(FieldReason, vs...))
}

// ReasonGT applies the GT predicate on the "reason" field.
func ReasonGT(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldGT(FieldReason, v))
}

// ReasonGTE applies the GTE predicate on the "reason" field.
func ReasonGTE(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldGTE(FieldReason, v))
}

// ReasonLT applies the LT predicate on the "reason" field.
func ReasonLT(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldLT(FieldReason, v))
}

// ReasonLTE applies the LTE predicate on the "reason" field.
func ReasonLTE(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldLTE(FieldReason, v))
}

// ReasonContains applies the Contains predicate on the "reason" field.
func ReasonContains(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldContains(FieldReason, v))
}

// ReasonHasPrefix applies the HasPrefix predicate on the "reason" field.
func ReasonHasPrefix(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldHasPrefix(FieldReason, v))
}

// ReasonHasSuffix applies the HasSuffix predicate on the "reason" field.
func ReasonHasSuffix(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldHasSuffix(FieldReason, v))
}

// ReasonIsNil applies the IsNil predicate on the "reason" field.
func ReasonIsNil() predicate.EventFile {
	return predicate.EventFile(sql.FieldIsNull(FieldReason))
}

// ReasonNotNil applies the NotNil predicate on the "reason" field.
func ReasonNotNil() predicate.EventFile {
	return predicate.EventFile(sql.FieldNotNull(FieldReason))
}

// ReasonEqualFold applies the EqualFold predicate on the "reason" field.
func ReasonEqualFold(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldEqualFold(FieldReason, v))
}

// ReasonContainsFold applies the ContainsFold predicate on the "reason" field.
func ReasonContainsFold(v string) predicate.EventFile {
	return predicate.EventFile(sql.FieldContainsFold(FieldReason, v))
}

// HasBazelInvocation applies the HasEdge predicate on the "bazel_invocation" edge.
func HasBazelInvocation() predicate.EventFile {
	return predicate.EventFile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BazelInvocationTable, BazelInvocationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBazelInvocationWith applies the HasEdge predicate on the "bazel_invocation" edge with a given conditions (other predicates).
func HasBazelInvocationWith(preds ...predicate.BazelInvocation) predicate.EventFile {
	return predicate.EventFile(func(s *sql.Selector) {
		step := newBazelInvocationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EventFile) predicate.EventFile {
	return predicate.EventFile(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EventFile) predicate.EventFile {
	return predicate.EventFile(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EventFile) predicate.EventFile {
	return predicate.EventFile(sql.NotPredicates(p))
}

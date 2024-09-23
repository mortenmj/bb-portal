// Code generated by ent, DO NOT EDIT.

package testfile

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TestFile {
	return predicate.TestFile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TestFile {
	return predicate.TestFile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TestFile {
	return predicate.TestFile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TestFile {
	return predicate.TestFile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TestFile {
	return predicate.TestFile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TestFile {
	return predicate.TestFile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TestFile {
	return predicate.TestFile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TestFile {
	return predicate.TestFile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TestFile {
	return predicate.TestFile(sql.FieldLTE(FieldID, id))
}

// Digest applies equality check predicate on the "digest" field. It's identical to DigestEQ.
func Digest(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldEQ(FieldDigest, v))
}

// File applies equality check predicate on the "file" field. It's identical to FileEQ.
func File(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldEQ(FieldFile, v))
}

// Length applies equality check predicate on the "length" field. It's identical to LengthEQ.
func Length(v int64) predicate.TestFile {
	return predicate.TestFile(sql.FieldEQ(FieldLength, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldEQ(FieldName, v))
}

// DigestEQ applies the EQ predicate on the "digest" field.
func DigestEQ(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldEQ(FieldDigest, v))
}

// DigestNEQ applies the NEQ predicate on the "digest" field.
func DigestNEQ(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldNEQ(FieldDigest, v))
}

// DigestIn applies the In predicate on the "digest" field.
func DigestIn(vs ...string) predicate.TestFile {
	return predicate.TestFile(sql.FieldIn(FieldDigest, vs...))
}

// DigestNotIn applies the NotIn predicate on the "digest" field.
func DigestNotIn(vs ...string) predicate.TestFile {
	return predicate.TestFile(sql.FieldNotIn(FieldDigest, vs...))
}

// DigestGT applies the GT predicate on the "digest" field.
func DigestGT(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldGT(FieldDigest, v))
}

// DigestGTE applies the GTE predicate on the "digest" field.
func DigestGTE(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldGTE(FieldDigest, v))
}

// DigestLT applies the LT predicate on the "digest" field.
func DigestLT(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldLT(FieldDigest, v))
}

// DigestLTE applies the LTE predicate on the "digest" field.
func DigestLTE(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldLTE(FieldDigest, v))
}

// DigestContains applies the Contains predicate on the "digest" field.
func DigestContains(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldContains(FieldDigest, v))
}

// DigestHasPrefix applies the HasPrefix predicate on the "digest" field.
func DigestHasPrefix(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldHasPrefix(FieldDigest, v))
}

// DigestHasSuffix applies the HasSuffix predicate on the "digest" field.
func DigestHasSuffix(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldHasSuffix(FieldDigest, v))
}

// DigestIsNil applies the IsNil predicate on the "digest" field.
func DigestIsNil() predicate.TestFile {
	return predicate.TestFile(sql.FieldIsNull(FieldDigest))
}

// DigestNotNil applies the NotNil predicate on the "digest" field.
func DigestNotNil() predicate.TestFile {
	return predicate.TestFile(sql.FieldNotNull(FieldDigest))
}

// DigestEqualFold applies the EqualFold predicate on the "digest" field.
func DigestEqualFold(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldEqualFold(FieldDigest, v))
}

// DigestContainsFold applies the ContainsFold predicate on the "digest" field.
func DigestContainsFold(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldContainsFold(FieldDigest, v))
}

// FileEQ applies the EQ predicate on the "file" field.
func FileEQ(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldEQ(FieldFile, v))
}

// FileNEQ applies the NEQ predicate on the "file" field.
func FileNEQ(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldNEQ(FieldFile, v))
}

// FileIn applies the In predicate on the "file" field.
func FileIn(vs ...string) predicate.TestFile {
	return predicate.TestFile(sql.FieldIn(FieldFile, vs...))
}

// FileNotIn applies the NotIn predicate on the "file" field.
func FileNotIn(vs ...string) predicate.TestFile {
	return predicate.TestFile(sql.FieldNotIn(FieldFile, vs...))
}

// FileGT applies the GT predicate on the "file" field.
func FileGT(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldGT(FieldFile, v))
}

// FileGTE applies the GTE predicate on the "file" field.
func FileGTE(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldGTE(FieldFile, v))
}

// FileLT applies the LT predicate on the "file" field.
func FileLT(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldLT(FieldFile, v))
}

// FileLTE applies the LTE predicate on the "file" field.
func FileLTE(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldLTE(FieldFile, v))
}

// FileContains applies the Contains predicate on the "file" field.
func FileContains(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldContains(FieldFile, v))
}

// FileHasPrefix applies the HasPrefix predicate on the "file" field.
func FileHasPrefix(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldHasPrefix(FieldFile, v))
}

// FileHasSuffix applies the HasSuffix predicate on the "file" field.
func FileHasSuffix(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldHasSuffix(FieldFile, v))
}

// FileIsNil applies the IsNil predicate on the "file" field.
func FileIsNil() predicate.TestFile {
	return predicate.TestFile(sql.FieldIsNull(FieldFile))
}

// FileNotNil applies the NotNil predicate on the "file" field.
func FileNotNil() predicate.TestFile {
	return predicate.TestFile(sql.FieldNotNull(FieldFile))
}

// FileEqualFold applies the EqualFold predicate on the "file" field.
func FileEqualFold(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldEqualFold(FieldFile, v))
}

// FileContainsFold applies the ContainsFold predicate on the "file" field.
func FileContainsFold(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldContainsFold(FieldFile, v))
}

// LengthEQ applies the EQ predicate on the "length" field.
func LengthEQ(v int64) predicate.TestFile {
	return predicate.TestFile(sql.FieldEQ(FieldLength, v))
}

// LengthNEQ applies the NEQ predicate on the "length" field.
func LengthNEQ(v int64) predicate.TestFile {
	return predicate.TestFile(sql.FieldNEQ(FieldLength, v))
}

// LengthIn applies the In predicate on the "length" field.
func LengthIn(vs ...int64) predicate.TestFile {
	return predicate.TestFile(sql.FieldIn(FieldLength, vs...))
}

// LengthNotIn applies the NotIn predicate on the "length" field.
func LengthNotIn(vs ...int64) predicate.TestFile {
	return predicate.TestFile(sql.FieldNotIn(FieldLength, vs...))
}

// LengthGT applies the GT predicate on the "length" field.
func LengthGT(v int64) predicate.TestFile {
	return predicate.TestFile(sql.FieldGT(FieldLength, v))
}

// LengthGTE applies the GTE predicate on the "length" field.
func LengthGTE(v int64) predicate.TestFile {
	return predicate.TestFile(sql.FieldGTE(FieldLength, v))
}

// LengthLT applies the LT predicate on the "length" field.
func LengthLT(v int64) predicate.TestFile {
	return predicate.TestFile(sql.FieldLT(FieldLength, v))
}

// LengthLTE applies the LTE predicate on the "length" field.
func LengthLTE(v int64) predicate.TestFile {
	return predicate.TestFile(sql.FieldLTE(FieldLength, v))
}

// LengthIsNil applies the IsNil predicate on the "length" field.
func LengthIsNil() predicate.TestFile {
	return predicate.TestFile(sql.FieldIsNull(FieldLength))
}

// LengthNotNil applies the NotNil predicate on the "length" field.
func LengthNotNil() predicate.TestFile {
	return predicate.TestFile(sql.FieldNotNull(FieldLength))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.TestFile {
	return predicate.TestFile(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.TestFile {
	return predicate.TestFile(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.TestFile {
	return predicate.TestFile(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.TestFile {
	return predicate.TestFile(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.TestFile {
	return predicate.TestFile(sql.FieldContainsFold(FieldName, v))
}

// PrefixIsNil applies the IsNil predicate on the "prefix" field.
func PrefixIsNil() predicate.TestFile {
	return predicate.TestFile(sql.FieldIsNull(FieldPrefix))
}

// PrefixNotNil applies the NotNil predicate on the "prefix" field.
func PrefixNotNil() predicate.TestFile {
	return predicate.TestFile(sql.FieldNotNull(FieldPrefix))
}

// HasTestResult applies the HasEdge predicate on the "test_result" edge.
func HasTestResult() predicate.TestFile {
	return predicate.TestFile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TestResultTable, TestResultPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestResultWith applies the HasEdge predicate on the "test_result" edge with a given conditions (other predicates).
func HasTestResultWith(preds ...predicate.TestResultBES) predicate.TestFile {
	return predicate.TestFile(func(s *sql.Selector) {
		step := newTestResultStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TestFile) predicate.TestFile {
	return predicate.TestFile(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TestFile) predicate.TestFile {
	return predicate.TestFile(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TestFile) predicate.TestFile {
	return predicate.TestFile(sql.NotPredicates(p))
}

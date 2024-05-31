// Code generated by ent, DO NOT EDIT.

package bazelinvocationproblem

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldLTE(FieldID, id))
}

// ProblemType applies equality check predicate on the "problem_type" field. It's identical to ProblemTypeEQ.
func ProblemType(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldEQ(FieldProblemType, v))
}

// ProblemTypeEQ applies the EQ predicate on the "problem_type" field.
func ProblemTypeEQ(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldEQ(FieldProblemType, v))
}

// ProblemTypeNEQ applies the NEQ predicate on the "problem_type" field.
func ProblemTypeNEQ(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldNEQ(FieldProblemType, v))
}

// ProblemTypeIn applies the In predicate on the "problem_type" field.
func ProblemTypeIn(vs ...string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldIn(FieldProblemType, vs...))
}

// ProblemTypeNotIn applies the NotIn predicate on the "problem_type" field.
func ProblemTypeNotIn(vs ...string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldNotIn(FieldProblemType, vs...))
}

// ProblemTypeGT applies the GT predicate on the "problem_type" field.
func ProblemTypeGT(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldGT(FieldProblemType, v))
}

// ProblemTypeGTE applies the GTE predicate on the "problem_type" field.
func ProblemTypeGTE(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldGTE(FieldProblemType, v))
}

// ProblemTypeLT applies the LT predicate on the "problem_type" field.
func ProblemTypeLT(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldLT(FieldProblemType, v))
}

// ProblemTypeLTE applies the LTE predicate on the "problem_type" field.
func ProblemTypeLTE(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldLTE(FieldProblemType, v))
}

// ProblemTypeContains applies the Contains predicate on the "problem_type" field.
func ProblemTypeContains(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldContains(FieldProblemType, v))
}

// ProblemTypeHasPrefix applies the HasPrefix predicate on the "problem_type" field.
func ProblemTypeHasPrefix(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldHasPrefix(FieldProblemType, v))
}

// ProblemTypeHasSuffix applies the HasSuffix predicate on the "problem_type" field.
func ProblemTypeHasSuffix(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldHasSuffix(FieldProblemType, v))
}

// ProblemTypeEqualFold applies the EqualFold predicate on the "problem_type" field.
func ProblemTypeEqualFold(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldEqualFold(FieldProblemType, v))
}

// ProblemTypeContainsFold applies the ContainsFold predicate on the "problem_type" field.
func ProblemTypeContainsFold(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldContainsFold(FieldProblemType, v))
}

// LabelEQ applies the EQ predicate on the "label" field.
func LabelEQ(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldEQ(FieldLabel, v))
}

// LabelNEQ applies the NEQ predicate on the "label" field.
func LabelNEQ(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldNEQ(FieldLabel, v))
}

// LabelIn applies the In predicate on the "label" field.
func LabelIn(vs ...string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldIn(FieldLabel, vs...))
}

// LabelNotIn applies the NotIn predicate on the "label" field.
func LabelNotIn(vs ...string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldNotIn(FieldLabel, vs...))
}

// LabelGT applies the GT predicate on the "label" field.
func LabelGT(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldGT(FieldLabel, v))
}

// LabelGTE applies the GTE predicate on the "label" field.
func LabelGTE(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldGTE(FieldLabel, v))
}

// LabelLT applies the LT predicate on the "label" field.
func LabelLT(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldLT(FieldLabel, v))
}

// LabelLTE applies the LTE predicate on the "label" field.
func LabelLTE(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldLTE(FieldLabel, v))
}

// LabelContains applies the Contains predicate on the "label" field.
func LabelContains(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldContains(FieldLabel, v))
}

// LabelHasPrefix applies the HasPrefix predicate on the "label" field.
func LabelHasPrefix(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldHasPrefix(FieldLabel, v))
}

// LabelHasSuffix applies the HasSuffix predicate on the "label" field.
func LabelHasSuffix(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldHasSuffix(FieldLabel, v))
}

// LabelEqualFold applies the EqualFold predicate on the "label" field.
func LabelEqualFold(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldEqualFold(FieldLabel, v))
}

// LabelContainsFold applies the ContainsFold predicate on the "label" field.
func LabelContainsFold(v string) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.FieldContainsFold(FieldLabel, v))
}

// HasBazelInvocation applies the HasEdge predicate on the "bazel_invocation" edge.
func HasBazelInvocation() predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BazelInvocationTable, BazelInvocationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBazelInvocationWith applies the HasEdge predicate on the "bazel_invocation" edge with a given conditions (other predicates).
func HasBazelInvocationWith(preds ...predicate.BazelInvocation) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(func(s *sql.Selector) {
		step := newBazelInvocationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BazelInvocationProblem) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BazelInvocationProblem) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BazelInvocationProblem) predicate.BazelInvocationProblem {
	return predicate.BazelInvocationProblem(sql.NotPredicates(p))
}

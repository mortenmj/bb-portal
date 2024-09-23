// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testcollection"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testsummary"
)

// TestCollection is the model entity for the TestCollection schema.
type TestCollection struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Label holds the value of the "label" field.
	Label string `json:"label,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TestCollectionQuery when eager-loading is set.
	Edges                        TestCollectionEdges `json:"edges"`
	test_collection_test_summary *int
	selectValues                 sql.SelectValues
}

// TestCollectionEdges holds the relations/edges for other nodes in the graph.
type TestCollectionEdges struct {
	// BazelInvocation holds the value of the bazel_invocation edge.
	BazelInvocation []*BazelInvocation `json:"bazel_invocation,omitempty"`
	// TestSummary holds the value of the test_summary edge.
	TestSummary *TestSummary `json:"test_summary,omitempty"`
	// TestResults holds the value of the test_results edge.
	TestResults []*TestResultBES `json:"test_results,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedBazelInvocation map[string][]*BazelInvocation
	namedTestResults     map[string][]*TestResultBES
}

// BazelInvocationOrErr returns the BazelInvocation value or an error if the edge
// was not loaded in eager-loading.
func (e TestCollectionEdges) BazelInvocationOrErr() ([]*BazelInvocation, error) {
	if e.loadedTypes[0] {
		return e.BazelInvocation, nil
	}
	return nil, &NotLoadedError{edge: "bazel_invocation"}
}

// TestSummaryOrErr returns the TestSummary value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TestCollectionEdges) TestSummaryOrErr() (*TestSummary, error) {
	if e.TestSummary != nil {
		return e.TestSummary, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: testsummary.Label}
	}
	return nil, &NotLoadedError{edge: "test_summary"}
}

// TestResultsOrErr returns the TestResults value or an error if the edge
// was not loaded in eager-loading.
func (e TestCollectionEdges) TestResultsOrErr() ([]*TestResultBES, error) {
	if e.loadedTypes[2] {
		return e.TestResults, nil
	}
	return nil, &NotLoadedError{edge: "test_results"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TestCollection) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case testcollection.FieldID:
			values[i] = new(sql.NullInt64)
		case testcollection.FieldLabel:
			values[i] = new(sql.NullString)
		case testcollection.ForeignKeys[0]: // test_collection_test_summary
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TestCollection fields.
func (tc *TestCollection) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case testcollection.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tc.ID = int(value.Int64)
		case testcollection.FieldLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field label", values[i])
			} else if value.Valid {
				tc.Label = value.String
			}
		case testcollection.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field test_collection_test_summary", value)
			} else if value.Valid {
				tc.test_collection_test_summary = new(int)
				*tc.test_collection_test_summary = int(value.Int64)
			}
		default:
			tc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TestCollection.
// This includes values selected through modifiers, order, etc.
func (tc *TestCollection) Value(name string) (ent.Value, error) {
	return tc.selectValues.Get(name)
}

// QueryBazelInvocation queries the "bazel_invocation" edge of the TestCollection entity.
func (tc *TestCollection) QueryBazelInvocation() *BazelInvocationQuery {
	return NewTestCollectionClient(tc.config).QueryBazelInvocation(tc)
}

// QueryTestSummary queries the "test_summary" edge of the TestCollection entity.
func (tc *TestCollection) QueryTestSummary() *TestSummaryQuery {
	return NewTestCollectionClient(tc.config).QueryTestSummary(tc)
}

// QueryTestResults queries the "test_results" edge of the TestCollection entity.
func (tc *TestCollection) QueryTestResults() *TestResultBESQuery {
	return NewTestCollectionClient(tc.config).QueryTestResults(tc)
}

// Update returns a builder for updating this TestCollection.
// Note that you need to call TestCollection.Unwrap() before calling this method if this TestCollection
// was returned from a transaction, and the transaction was committed or rolled back.
func (tc *TestCollection) Update() *TestCollectionUpdateOne {
	return NewTestCollectionClient(tc.config).UpdateOne(tc)
}

// Unwrap unwraps the TestCollection entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tc *TestCollection) Unwrap() *TestCollection {
	_tx, ok := tc.config.driver.(*txDriver)
	if !ok {
		panic("ent: TestCollection is not a transactional entity")
	}
	tc.config.driver = _tx.drv
	return tc
}

// String implements the fmt.Stringer.
func (tc *TestCollection) String() string {
	var builder strings.Builder
	builder.WriteString("TestCollection(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tc.ID))
	builder.WriteString("label=")
	builder.WriteString(tc.Label)
	builder.WriteByte(')')
	return builder.String()
}

// NamedBazelInvocation returns the BazelInvocation named value or an error if the edge was not
// loaded in eager-loading with this name.
func (tc *TestCollection) NamedBazelInvocation(name string) ([]*BazelInvocation, error) {
	if tc.Edges.namedBazelInvocation == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := tc.Edges.namedBazelInvocation[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (tc *TestCollection) appendNamedBazelInvocation(name string, edges ...*BazelInvocation) {
	if tc.Edges.namedBazelInvocation == nil {
		tc.Edges.namedBazelInvocation = make(map[string][]*BazelInvocation)
	}
	if len(edges) == 0 {
		tc.Edges.namedBazelInvocation[name] = []*BazelInvocation{}
	} else {
		tc.Edges.namedBazelInvocation[name] = append(tc.Edges.namedBazelInvocation[name], edges...)
	}
}

// NamedTestResults returns the TestResults named value or an error if the edge was not
// loaded in eager-loading with this name.
func (tc *TestCollection) NamedTestResults(name string) ([]*TestResultBES, error) {
	if tc.Edges.namedTestResults == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := tc.Edges.namedTestResults[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (tc *TestCollection) appendNamedTestResults(name string, edges ...*TestResultBES) {
	if tc.Edges.namedTestResults == nil {
		tc.Edges.namedTestResults = make(map[string][]*TestResultBES)
	}
	if len(edges) == 0 {
		tc.Edges.namedTestResults[name] = []*TestResultBES{}
	} else {
		tc.Edges.namedTestResults[name] = append(tc.Edges.namedTestResults[name], edges...)
	}
}

// TestCollections is a parsable slice of TestCollection.
type TestCollections []*TestCollection

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/buildbarn/bb-portal/ent/gen/ent/exectioninfo"
	"github.com/buildbarn/bb-portal/ent/gen/ent/timingbreakdown"
)

// ExectionInfo is the model entity for the ExectionInfo schema.
type ExectionInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Strategy holds the value of the "strategy" field.
	Strategy string `json:"strategy,omitempty"`
	// CachedRemotely holds the value of the "cached_remotely" field.
	CachedRemotely bool `json:"cached_remotely,omitempty"`
	// ExitCode holds the value of the "exit_code" field.
	ExitCode int32 `json:"exit_code,omitempty"`
	// Hostname holds the value of the "hostname" field.
	Hostname string `json:"hostname,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExectionInfoQuery when eager-loading is set.
	Edges                          ExectionInfoEdges `json:"edges"`
	exection_info_timing_breakdown *int
	selectValues                   sql.SelectValues
}

// ExectionInfoEdges holds the relations/edges for other nodes in the graph.
type ExectionInfoEdges struct {
	// TestResult holds the value of the test_result edge.
	TestResult []*TestResultBES `json:"test_result,omitempty"`
	// TimingBreakdown holds the value of the timing_breakdown edge.
	TimingBreakdown *TimingBreakdown `json:"timing_breakdown,omitempty"`
	// ResourceUsage holds the value of the resource_usage edge.
	ResourceUsage []*ResourceUsage `json:"resource_usage,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedTestResult    map[string][]*TestResultBES
	namedResourceUsage map[string][]*ResourceUsage
}

// TestResultOrErr returns the TestResult value or an error if the edge
// was not loaded in eager-loading.
func (e ExectionInfoEdges) TestResultOrErr() ([]*TestResultBES, error) {
	if e.loadedTypes[0] {
		return e.TestResult, nil
	}
	return nil, &NotLoadedError{edge: "test_result"}
}

// TimingBreakdownOrErr returns the TimingBreakdown value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ExectionInfoEdges) TimingBreakdownOrErr() (*TimingBreakdown, error) {
	if e.TimingBreakdown != nil {
		return e.TimingBreakdown, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: timingbreakdown.Label}
	}
	return nil, &NotLoadedError{edge: "timing_breakdown"}
}

// ResourceUsageOrErr returns the ResourceUsage value or an error if the edge
// was not loaded in eager-loading.
func (e ExectionInfoEdges) ResourceUsageOrErr() ([]*ResourceUsage, error) {
	if e.loadedTypes[2] {
		return e.ResourceUsage, nil
	}
	return nil, &NotLoadedError{edge: "resource_usage"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ExectionInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case exectioninfo.FieldCachedRemotely:
			values[i] = new(sql.NullBool)
		case exectioninfo.FieldID, exectioninfo.FieldExitCode:
			values[i] = new(sql.NullInt64)
		case exectioninfo.FieldStrategy, exectioninfo.FieldHostname:
			values[i] = new(sql.NullString)
		case exectioninfo.ForeignKeys[0]: // exection_info_timing_breakdown
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExectionInfo fields.
func (ei *ExectionInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case exectioninfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ei.ID = int(value.Int64)
		case exectioninfo.FieldStrategy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field strategy", values[i])
			} else if value.Valid {
				ei.Strategy = value.String
			}
		case exectioninfo.FieldCachedRemotely:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field cached_remotely", values[i])
			} else if value.Valid {
				ei.CachedRemotely = value.Bool
			}
		case exectioninfo.FieldExitCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field exit_code", values[i])
			} else if value.Valid {
				ei.ExitCode = int32(value.Int64)
			}
		case exectioninfo.FieldHostname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hostname", values[i])
			} else if value.Valid {
				ei.Hostname = value.String
			}
		case exectioninfo.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field exection_info_timing_breakdown", value)
			} else if value.Valid {
				ei.exection_info_timing_breakdown = new(int)
				*ei.exection_info_timing_breakdown = int(value.Int64)
			}
		default:
			ei.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ExectionInfo.
// This includes values selected through modifiers, order, etc.
func (ei *ExectionInfo) Value(name string) (ent.Value, error) {
	return ei.selectValues.Get(name)
}

// QueryTestResult queries the "test_result" edge of the ExectionInfo entity.
func (ei *ExectionInfo) QueryTestResult() *TestResultBESQuery {
	return NewExectionInfoClient(ei.config).QueryTestResult(ei)
}

// QueryTimingBreakdown queries the "timing_breakdown" edge of the ExectionInfo entity.
func (ei *ExectionInfo) QueryTimingBreakdown() *TimingBreakdownQuery {
	return NewExectionInfoClient(ei.config).QueryTimingBreakdown(ei)
}

// QueryResourceUsage queries the "resource_usage" edge of the ExectionInfo entity.
func (ei *ExectionInfo) QueryResourceUsage() *ResourceUsageQuery {
	return NewExectionInfoClient(ei.config).QueryResourceUsage(ei)
}

// Update returns a builder for updating this ExectionInfo.
// Note that you need to call ExectionInfo.Unwrap() before calling this method if this ExectionInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ei *ExectionInfo) Update() *ExectionInfoUpdateOne {
	return NewExectionInfoClient(ei.config).UpdateOne(ei)
}

// Unwrap unwraps the ExectionInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ei *ExectionInfo) Unwrap() *ExectionInfo {
	_tx, ok := ei.config.driver.(*txDriver)
	if !ok {
		panic("ent: ExectionInfo is not a transactional entity")
	}
	ei.config.driver = _tx.drv
	return ei
}

// String implements the fmt.Stringer.
func (ei *ExectionInfo) String() string {
	var builder strings.Builder
	builder.WriteString("ExectionInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ei.ID))
	builder.WriteString("strategy=")
	builder.WriteString(ei.Strategy)
	builder.WriteString(", ")
	builder.WriteString("cached_remotely=")
	builder.WriteString(fmt.Sprintf("%v", ei.CachedRemotely))
	builder.WriteString(", ")
	builder.WriteString("exit_code=")
	builder.WriteString(fmt.Sprintf("%v", ei.ExitCode))
	builder.WriteString(", ")
	builder.WriteString("hostname=")
	builder.WriteString(ei.Hostname)
	builder.WriteByte(')')
	return builder.String()
}

// NamedTestResult returns the TestResult named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ei *ExectionInfo) NamedTestResult(name string) ([]*TestResultBES, error) {
	if ei.Edges.namedTestResult == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ei.Edges.namedTestResult[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ei *ExectionInfo) appendNamedTestResult(name string, edges ...*TestResultBES) {
	if ei.Edges.namedTestResult == nil {
		ei.Edges.namedTestResult = make(map[string][]*TestResultBES)
	}
	if len(edges) == 0 {
		ei.Edges.namedTestResult[name] = []*TestResultBES{}
	} else {
		ei.Edges.namedTestResult[name] = append(ei.Edges.namedTestResult[name], edges...)
	}
}

// NamedResourceUsage returns the ResourceUsage named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ei *ExectionInfo) NamedResourceUsage(name string) ([]*ResourceUsage, error) {
	if ei.Edges.namedResourceUsage == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ei.Edges.namedResourceUsage[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ei *ExectionInfo) appendNamedResourceUsage(name string, edges ...*ResourceUsage) {
	if ei.Edges.namedResourceUsage == nil {
		ei.Edges.namedResourceUsage = make(map[string][]*ResourceUsage)
	}
	if len(edges) == 0 {
		ei.Edges.namedResourceUsage[name] = []*ResourceUsage{}
	} else {
		ei.Edges.namedResourceUsage[name] = append(ei.Edges.namedResourceUsage[name], edges...)
	}
}

// ExectionInfos is a parsable slice of ExectionInfo.
type ExectionInfos []*ExectionInfo
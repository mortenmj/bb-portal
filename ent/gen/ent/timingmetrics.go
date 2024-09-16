// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/buildbarn/bb-portal/ent/gen/ent/timingmetrics"
)

// TimingMetrics is the model entity for the TimingMetrics schema.
type TimingMetrics struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CPUTimeInMs holds the value of the "cpu_time_in_ms" field.
	CPUTimeInMs int64 `json:"cpu_time_in_ms,omitempty"`
	// WallTimeInMs holds the value of the "wall_time_in_ms" field.
	WallTimeInMs int64 `json:"wall_time_in_ms,omitempty"`
	// AnalysisPhaseTimeInMs holds the value of the "analysis_phase_time_in_ms" field.
	AnalysisPhaseTimeInMs int64 `json:"analysis_phase_time_in_ms,omitempty"`
	// ExecutionPhaseTimeInMs holds the value of the "execution_phase_time_in_ms" field.
	ExecutionPhaseTimeInMs int64 `json:"execution_phase_time_in_ms,omitempty"`
	// ActionsExecutionStartInMs holds the value of the "actions_execution_start_in_ms" field.
	ActionsExecutionStartInMs int64 `json:"actions_execution_start_in_ms,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TimingMetricsQuery when eager-loading is set.
	Edges        TimingMetricsEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TimingMetricsEdges holds the relations/edges for other nodes in the graph.
type TimingMetricsEdges struct {
	// Metrics holds the value of the metrics edge.
	Metrics []*Metrics `json:"metrics,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedMetrics map[string][]*Metrics
}

// MetricsOrErr returns the Metrics value or an error if the edge
// was not loaded in eager-loading.
func (e TimingMetricsEdges) MetricsOrErr() ([]*Metrics, error) {
	if e.loadedTypes[0] {
		return e.Metrics, nil
	}
	return nil, &NotLoadedError{edge: "metrics"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TimingMetrics) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case timingmetrics.FieldID, timingmetrics.FieldCPUTimeInMs, timingmetrics.FieldWallTimeInMs, timingmetrics.FieldAnalysisPhaseTimeInMs, timingmetrics.FieldExecutionPhaseTimeInMs, timingmetrics.FieldActionsExecutionStartInMs:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TimingMetrics fields.
func (tm *TimingMetrics) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case timingmetrics.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tm.ID = int(value.Int64)
		case timingmetrics.FieldCPUTimeInMs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cpu_time_in_ms", values[i])
			} else if value.Valid {
				tm.CPUTimeInMs = value.Int64
			}
		case timingmetrics.FieldWallTimeInMs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wall_time_in_ms", values[i])
			} else if value.Valid {
				tm.WallTimeInMs = value.Int64
			}
		case timingmetrics.FieldAnalysisPhaseTimeInMs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field analysis_phase_time_in_ms", values[i])
			} else if value.Valid {
				tm.AnalysisPhaseTimeInMs = value.Int64
			}
		case timingmetrics.FieldExecutionPhaseTimeInMs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field execution_phase_time_in_ms", values[i])
			} else if value.Valid {
				tm.ExecutionPhaseTimeInMs = value.Int64
			}
		case timingmetrics.FieldActionsExecutionStartInMs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field actions_execution_start_in_ms", values[i])
			} else if value.Valid {
				tm.ActionsExecutionStartInMs = value.Int64
			}
		default:
			tm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TimingMetrics.
// This includes values selected through modifiers, order, etc.
func (tm *TimingMetrics) Value(name string) (ent.Value, error) {
	return tm.selectValues.Get(name)
}

// QueryMetrics queries the "metrics" edge of the TimingMetrics entity.
func (tm *TimingMetrics) QueryMetrics() *MetricsQuery {
	return NewTimingMetricsClient(tm.config).QueryMetrics(tm)
}

// Update returns a builder for updating this TimingMetrics.
// Note that you need to call TimingMetrics.Unwrap() before calling this method if this TimingMetrics
// was returned from a transaction, and the transaction was committed or rolled back.
func (tm *TimingMetrics) Update() *TimingMetricsUpdateOne {
	return NewTimingMetricsClient(tm.config).UpdateOne(tm)
}

// Unwrap unwraps the TimingMetrics entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tm *TimingMetrics) Unwrap() *TimingMetrics {
	_tx, ok := tm.config.driver.(*txDriver)
	if !ok {
		panic("ent: TimingMetrics is not a transactional entity")
	}
	tm.config.driver = _tx.drv
	return tm
}

// String implements the fmt.Stringer.
func (tm *TimingMetrics) String() string {
	var builder strings.Builder
	builder.WriteString("TimingMetrics(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tm.ID))
	builder.WriteString("cpu_time_in_ms=")
	builder.WriteString(fmt.Sprintf("%v", tm.CPUTimeInMs))
	builder.WriteString(", ")
	builder.WriteString("wall_time_in_ms=")
	builder.WriteString(fmt.Sprintf("%v", tm.WallTimeInMs))
	builder.WriteString(", ")
	builder.WriteString("analysis_phase_time_in_ms=")
	builder.WriteString(fmt.Sprintf("%v", tm.AnalysisPhaseTimeInMs))
	builder.WriteString(", ")
	builder.WriteString("execution_phase_time_in_ms=")
	builder.WriteString(fmt.Sprintf("%v", tm.ExecutionPhaseTimeInMs))
	builder.WriteString(", ")
	builder.WriteString("actions_execution_start_in_ms=")
	builder.WriteString(fmt.Sprintf("%v", tm.ActionsExecutionStartInMs))
	builder.WriteByte(')')
	return builder.String()
}

// NamedMetrics returns the Metrics named value or an error if the edge was not
// loaded in eager-loading with this name.
func (tm *TimingMetrics) NamedMetrics(name string) ([]*Metrics, error) {
	if tm.Edges.namedMetrics == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := tm.Edges.namedMetrics[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (tm *TimingMetrics) appendNamedMetrics(name string, edges ...*Metrics) {
	if tm.Edges.namedMetrics == nil {
		tm.Edges.namedMetrics = make(map[string][]*Metrics)
	}
	if len(edges) == 0 {
		tm.Edges.namedMetrics[name] = []*Metrics{}
	} else {
		tm.Edges.namedMetrics[name] = append(tm.Edges.namedMetrics[name], edges...)
	}
}

// TimingMetricsSlice is a parsable slice of TimingMetrics.
type TimingMetricsSlice []*TimingMetrics
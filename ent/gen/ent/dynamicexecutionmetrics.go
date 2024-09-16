// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/buildbarn/bb-portal/ent/gen/ent/dynamicexecutionmetrics"
)

// DynamicExecutionMetrics is the model entity for the DynamicExecutionMetrics schema.
type DynamicExecutionMetrics struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DynamicExecutionMetricsQuery when eager-loading is set.
	Edges        DynamicExecutionMetricsEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DynamicExecutionMetricsEdges holds the relations/edges for other nodes in the graph.
type DynamicExecutionMetricsEdges struct {
	// Metrics holds the value of the metrics edge.
	Metrics []*Metrics `json:"metrics,omitempty"`
	// RaceStatistics holds the value of the race_statistics edge.
	RaceStatistics []*RaceStatistics `json:"race_statistics,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedMetrics        map[string][]*Metrics
	namedRaceStatistics map[string][]*RaceStatistics
}

// MetricsOrErr returns the Metrics value or an error if the edge
// was not loaded in eager-loading.
func (e DynamicExecutionMetricsEdges) MetricsOrErr() ([]*Metrics, error) {
	if e.loadedTypes[0] {
		return e.Metrics, nil
	}
	return nil, &NotLoadedError{edge: "metrics"}
}

// RaceStatisticsOrErr returns the RaceStatistics value or an error if the edge
// was not loaded in eager-loading.
func (e DynamicExecutionMetricsEdges) RaceStatisticsOrErr() ([]*RaceStatistics, error) {
	if e.loadedTypes[1] {
		return e.RaceStatistics, nil
	}
	return nil, &NotLoadedError{edge: "race_statistics"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DynamicExecutionMetrics) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dynamicexecutionmetrics.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DynamicExecutionMetrics fields.
func (dem *DynamicExecutionMetrics) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dynamicexecutionmetrics.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dem.ID = int(value.Int64)
		default:
			dem.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DynamicExecutionMetrics.
// This includes values selected through modifiers, order, etc.
func (dem *DynamicExecutionMetrics) Value(name string) (ent.Value, error) {
	return dem.selectValues.Get(name)
}

// QueryMetrics queries the "metrics" edge of the DynamicExecutionMetrics entity.
func (dem *DynamicExecutionMetrics) QueryMetrics() *MetricsQuery {
	return NewDynamicExecutionMetricsClient(dem.config).QueryMetrics(dem)
}

// QueryRaceStatistics queries the "race_statistics" edge of the DynamicExecutionMetrics entity.
func (dem *DynamicExecutionMetrics) QueryRaceStatistics() *RaceStatisticsQuery {
	return NewDynamicExecutionMetricsClient(dem.config).QueryRaceStatistics(dem)
}

// Update returns a builder for updating this DynamicExecutionMetrics.
// Note that you need to call DynamicExecutionMetrics.Unwrap() before calling this method if this DynamicExecutionMetrics
// was returned from a transaction, and the transaction was committed or rolled back.
func (dem *DynamicExecutionMetrics) Update() *DynamicExecutionMetricsUpdateOne {
	return NewDynamicExecutionMetricsClient(dem.config).UpdateOne(dem)
}

// Unwrap unwraps the DynamicExecutionMetrics entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dem *DynamicExecutionMetrics) Unwrap() *DynamicExecutionMetrics {
	_tx, ok := dem.config.driver.(*txDriver)
	if !ok {
		panic("ent: DynamicExecutionMetrics is not a transactional entity")
	}
	dem.config.driver = _tx.drv
	return dem
}

// String implements the fmt.Stringer.
func (dem *DynamicExecutionMetrics) String() string {
	var builder strings.Builder
	builder.WriteString("DynamicExecutionMetrics(")
	builder.WriteString(fmt.Sprintf("id=%v", dem.ID))
	builder.WriteByte(')')
	return builder.String()
}

// NamedMetrics returns the Metrics named value or an error if the edge was not
// loaded in eager-loading with this name.
func (dem *DynamicExecutionMetrics) NamedMetrics(name string) ([]*Metrics, error) {
	if dem.Edges.namedMetrics == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := dem.Edges.namedMetrics[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (dem *DynamicExecutionMetrics) appendNamedMetrics(name string, edges ...*Metrics) {
	if dem.Edges.namedMetrics == nil {
		dem.Edges.namedMetrics = make(map[string][]*Metrics)
	}
	if len(edges) == 0 {
		dem.Edges.namedMetrics[name] = []*Metrics{}
	} else {
		dem.Edges.namedMetrics[name] = append(dem.Edges.namedMetrics[name], edges...)
	}
}

// NamedRaceStatistics returns the RaceStatistics named value or an error if the edge was not
// loaded in eager-loading with this name.
func (dem *DynamicExecutionMetrics) NamedRaceStatistics(name string) ([]*RaceStatistics, error) {
	if dem.Edges.namedRaceStatistics == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := dem.Edges.namedRaceStatistics[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (dem *DynamicExecutionMetrics) appendNamedRaceStatistics(name string, edges ...*RaceStatistics) {
	if dem.Edges.namedRaceStatistics == nil {
		dem.Edges.namedRaceStatistics = make(map[string][]*RaceStatistics)
	}
	if len(edges) == 0 {
		dem.Edges.namedRaceStatistics[name] = []*RaceStatistics{}
	} else {
		dem.Edges.namedRaceStatistics[name] = append(dem.Edges.namedRaceStatistics[name], edges...)
	}
}

// DynamicExecutionMetricsSlice is a parsable slice of DynamicExecutionMetrics.
type DynamicExecutionMetricsSlice []*DynamicExecutionMetrics

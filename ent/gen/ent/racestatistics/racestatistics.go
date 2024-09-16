// Code generated by ent, DO NOT EDIT.

package racestatistics

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the racestatistics type in the database.
	Label = "race_statistics"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMnemonic holds the string denoting the mnemonic field in the database.
	FieldMnemonic = "mnemonic"
	// FieldLocalRunner holds the string denoting the local_runner field in the database.
	FieldLocalRunner = "local_runner"
	// FieldRemoteRunner holds the string denoting the remote_runner field in the database.
	FieldRemoteRunner = "remote_runner"
	// FieldLocalWins holds the string denoting the local_wins field in the database.
	FieldLocalWins = "local_wins"
	// FieldRenoteWins holds the string denoting the renote_wins field in the database.
	FieldRenoteWins = "renote_wins"
	// EdgeDynamicExecutionMetrics holds the string denoting the dynamic_execution_metrics edge name in mutations.
	EdgeDynamicExecutionMetrics = "dynamic_execution_metrics"
	// Table holds the table name of the racestatistics in the database.
	Table = "race_statistics"
	// DynamicExecutionMetricsTable is the table that holds the dynamic_execution_metrics relation/edge. The primary key declared below.
	DynamicExecutionMetricsTable = "dynamic_execution_metrics_race_statistics"
	// DynamicExecutionMetricsInverseTable is the table name for the DynamicExecutionMetrics entity.
	// It exists in this package in order to avoid circular dependency with the "dynamicexecutionmetrics" package.
	DynamicExecutionMetricsInverseTable = "dynamic_execution_metrics"
)

// Columns holds all SQL columns for racestatistics fields.
var Columns = []string{
	FieldID,
	FieldMnemonic,
	FieldLocalRunner,
	FieldRemoteRunner,
	FieldLocalWins,
	FieldRenoteWins,
}

var (
	// DynamicExecutionMetricsPrimaryKey and DynamicExecutionMetricsColumn2 are the table columns denoting the
	// primary key for the dynamic_execution_metrics relation (M2M).
	DynamicExecutionMetricsPrimaryKey = []string{"dynamic_execution_metrics_id", "race_statistics_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the RaceStatistics queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMnemonic orders the results by the mnemonic field.
func ByMnemonic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMnemonic, opts...).ToFunc()
}

// ByLocalRunner orders the results by the local_runner field.
func ByLocalRunner(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocalRunner, opts...).ToFunc()
}

// ByRemoteRunner orders the results by the remote_runner field.
func ByRemoteRunner(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemoteRunner, opts...).ToFunc()
}

// ByLocalWins orders the results by the local_wins field.
func ByLocalWins(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocalWins, opts...).ToFunc()
}

// ByRenoteWins orders the results by the renote_wins field.
func ByRenoteWins(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRenoteWins, opts...).ToFunc()
}

// ByDynamicExecutionMetricsCount orders the results by dynamic_execution_metrics count.
func ByDynamicExecutionMetricsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDynamicExecutionMetricsStep(), opts...)
	}
}

// ByDynamicExecutionMetrics orders the results by dynamic_execution_metrics terms.
func ByDynamicExecutionMetrics(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDynamicExecutionMetricsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDynamicExecutionMetricsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DynamicExecutionMetricsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, DynamicExecutionMetricsTable, DynamicExecutionMetricsPrimaryKey...),
	)
}

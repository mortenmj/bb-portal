// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BazelInvocationsColumns holds the columns for the "bazel_invocations" table.
	BazelInvocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "invocation_id", Type: field.TypeUUID, Unique: true},
		{Name: "started_at", Type: field.TypeTime},
		{Name: "ended_at", Type: field.TypeTime, Nullable: true},
		{Name: "change_number", Type: field.TypeInt32, Nullable: true},
		{Name: "patchset_number", Type: field.TypeInt32, Nullable: true},
		{Name: "summary", Type: field.TypeJSON},
		{Name: "bep_completed", Type: field.TypeBool, Nullable: true},
		{Name: "step_label", Type: field.TypeString},
		{Name: "related_files", Type: field.TypeJSON},
		{Name: "user_email", Type: field.TypeString, Nullable: true},
		{Name: "user_ldap", Type: field.TypeString, Nullable: true},
		{Name: "build_logs", Type: field.TypeString, Nullable: true},
		{Name: "build_invocations", Type: field.TypeInt, Nullable: true},
		{Name: "event_file_bazel_invocation", Type: field.TypeInt, Unique: true},
	}
	// BazelInvocationsTable holds the schema information for the "bazel_invocations" table.
	BazelInvocationsTable = &schema.Table{
		Name:       "bazel_invocations",
		Columns:    BazelInvocationsColumns,
		PrimaryKey: []*schema.Column{BazelInvocationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bazel_invocations_builds_invocations",
				Columns:    []*schema.Column{BazelInvocationsColumns[13]},
				RefColumns: []*schema.Column{BuildsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "bazel_invocations_event_files_bazel_invocation",
				Columns:    []*schema.Column{BazelInvocationsColumns[14]},
				RefColumns: []*schema.Column{EventFilesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "bazelinvocation_change_number_patchset_number",
				Unique:  false,
				Columns: []*schema.Column{BazelInvocationsColumns[4], BazelInvocationsColumns[5]},
			},
		},
	}
	// BazelInvocationProblemsColumns holds the columns for the "bazel_invocation_problems" table.
	BazelInvocationProblemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "problem_type", Type: field.TypeString},
		{Name: "label", Type: field.TypeString},
		{Name: "bep_events", Type: field.TypeJSON},
		{Name: "bazel_invocation_problems", Type: field.TypeInt, Nullable: true},
	}
	// BazelInvocationProblemsTable holds the schema information for the "bazel_invocation_problems" table.
	BazelInvocationProblemsTable = &schema.Table{
		Name:       "bazel_invocation_problems",
		Columns:    BazelInvocationProblemsColumns,
		PrimaryKey: []*schema.Column{BazelInvocationProblemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bazel_invocation_problems_bazel_invocations_problems",
				Columns:    []*schema.Column{BazelInvocationProblemsColumns[4]},
				RefColumns: []*schema.Column{BazelInvocationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BlobsColumns holds the columns for the "blobs" table.
	BlobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uri", Type: field.TypeString, Unique: true},
		{Name: "size_bytes", Type: field.TypeInt64, Nullable: true},
		{Name: "archiving_status", Type: field.TypeEnum, Enums: []string{"QUEUED", "ARCHIVING", "SUCCESS", "FAILED"}, Default: "QUEUED"},
		{Name: "reason", Type: field.TypeString, Nullable: true},
		{Name: "archive_url", Type: field.TypeString, Nullable: true},
	}
	// BlobsTable holds the schema information for the "blobs" table.
	BlobsTable = &schema.Table{
		Name:       "blobs",
		Columns:    BlobsColumns,
		PrimaryKey: []*schema.Column{BlobsColumns[0]},
	}
	// BuildsColumns holds the columns for the "builds" table.
	BuildsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "build_url", Type: field.TypeString, Unique: true},
		{Name: "build_uuid", Type: field.TypeUUID, Unique: true},
		{Name: "env", Type: field.TypeJSON},
	}
	// BuildsTable holds the schema information for the "builds" table.
	BuildsTable = &schema.Table{
		Name:       "builds",
		Columns:    BuildsColumns,
		PrimaryKey: []*schema.Column{BuildsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "build_env",
				Unique:  false,
				Columns: []*schema.Column{BuildsColumns[3]},
			},
		},
	}
	// EventFilesColumns holds the columns for the "event_files" table.
	EventFilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString, Unique: true},
		{Name: "mod_time", Type: field.TypeTime},
		{Name: "protocol", Type: field.TypeString},
		{Name: "mime_type", Type: field.TypeString},
		{Name: "status", Type: field.TypeString, Default: "DETECTED"},
		{Name: "reason", Type: field.TypeString, Nullable: true},
	}
	// EventFilesTable holds the schema information for the "event_files" table.
	EventFilesTable = &schema.Table{
		Name:       "event_files",
		Columns:    EventFilesColumns,
		PrimaryKey: []*schema.Column{EventFilesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "eventfile_status",
				Unique:  false,
				Columns: []*schema.Column{EventFilesColumns[5]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BazelInvocationsTable,
		BazelInvocationProblemsTable,
		BlobsTable,
		BuildsTable,
		EventFilesTable,
	}
)

func init() {
	BazelInvocationsTable.ForeignKeys[0].RefTable = BuildsTable
	BazelInvocationsTable.ForeignKeys[1].RefTable = EventFilesTable
	BazelInvocationProblemsTable.ForeignKeys[0].RefTable = BazelInvocationsTable
}

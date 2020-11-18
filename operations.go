package harperdb

import "time"

const (
	OP_ADD_NODE               = "add_node"
	OP_ADD_ROLE               = "add_role"
	OP_ADD_USER               = "add_user"
	OP_ALTER_ROLE             = "alter_role"
	OP_ALTER_USER             = "alter_user"
	OP_CLUSTER_STATUS         = "cluster_status"
	OP_CREATE_ATTRIBUTE       = "create_attribute"
	OP_CREATE_SCHEMA          = "create_schema"
	OP_CREATE_TABLE           = "create_table"
	OP_CSV_DATA_LOAD          = "csv_data_load"
	OP_CSV_FILE_LOAD          = "csv_file_load"
	OP_CSV_URL_LOAD           = "csv_url_load"
	OP_DELETE_FILES_BEFORE    = "delete_files_before"
	OP_DELETE_TRANSACTION_LOG = "delete_transaction_logs_before"
	OP_DESCRIBE_ALL           = "describe_all"
	OP_DESCRIBE_SCHEMA        = "describe_schema"
	OP_DESCRIBE_TABLE         = "describe_table"
	OP_DELETE                 = "delete"
	OP_DROP_ATTRIBUTE         = "drop_atttribute"
	OP_DROP_ROLE              = "drop_role"
	OP_DROP_SCHEMA            = "drop_schema"
	OP_DROP_TABLE             = "drop_table"
	OP_DROP_USER              = "drop_user"
	OP_EXPORT_LOCAL           = "export_local"
	OP_EXPORT_TO_S3           = "export_to_s3"
	OP_GET_FINGERPRINT        = "get_fingerprint"
	OP_GET_JOB                = "get_job"
	OP_INSERT                 = "insert"
	OP_LIST_ROLES             = "list_roles"
	OP_LIST_USERS             = "list_users"
	OP_READ_LOG               = "read_log"
	OP_READ_TRANSACTION_LOG   = "read_transaction_log"
	OP_REGISTRATION_INFO      = "registration_info"
	OP_REMOVE_NODE            = "remove_node"
	OP_SEARCH_BY_HASH         = "search_by_hash"
	OP_SEARCH_BY_VALUE        = "search_by_value"
	OP_SEARCH_JOBS            = "search_jobs_by_start_date"
	OP_SET_LICENSE            = "set_license"
	OP_SQL                    = "sql"
	OP_SYSTEM_INFORMATION     = "system_information"
	OP_UPDATE                 = "update"
	OP_UPDATE_NODE            = "update_node"
	OP_USER_INFO              = "user_info"
)

type operation struct {
	Action          string          `json:"action,omitempty"`
	Active          *bool           `json:"active,omitempty"`
	Attribute       string          `json:"attribute,omitempty"`
	Company         string          `json:"company"`
	CSVURL          string          `json:"csv_url,omitempty"`
	Data            string          `json:"data,omitempty"`
	Date            time.Time       `json:"date,omitempty"`
	FilePath        string          `json:"file_path,omitempty"`
	Format          string          `json:"format"`
	From            string          `json:"from"`
	FromDate        string          `json:"from_date,omitempty"`
	GetAttributes   AttributeList   `json:"get_attributes,omitempty"`
	HashAttribute   string          `json:"hash_attribute,omitempty"`
	HashValues      interface{}     `json:"hash_values,omitempty"`
	Host            string          `json:"host,omitempty"`
	ID              string          `json:"id,omitempty"`
	Key             string          `json:"key"`
	Limit           int             `json:"limit"`
	Name            string          `json:"name,omitempty"`
	Operation       string          `json:"operation"`
	Order           string          `json:"order"`
	Password        string          `json:"password,omitempty"`
	Path            string          `json:"path"`
	Permission      Permission      `json:"permission,omitempty"`
	Port            int             `json:"port"`
	Records         interface{}     `json:"records,omitempty"`
	Role            string          `json:"role,omitempty"`
	S3              S3Credentials   `json:"s3"`
	Schema          string          `json:"schema,omitempty"`
	SearchAttribute Attribute       `json:"search_attribute,omitempty"`
	SearchOperation SearchOperation `json:"search_operation"`
	SearchType      string          `json:"search_type.omitempty"`
	SearchValue     interface{}     `json:"search_value,omitempty"`
	SearchValues    interface{}     `json:"search_values,omitempty"`
	Start           int             `json:"start"`
	Subscriptions   []Subscription  `json:"subscriptions,omitempty"`
	SQL             string          `json:"sql,omitempty"`
	Table           string          `json:"table,omitempty"`
	Timestamp       int64           `json:"timestamp,omitempty"`
	ToDate          string          `json:"to_date,omitempty"`
	Until           string          `json:"until"`
	Username        string          `json:"username,omitempty"`
}

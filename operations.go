package harperdb

const (
	OP_REGISTRATION_INFO = "registration_info"
	OP_GET_FINGERPRINT   = "get_fingerprint"
	OP_CREATE_SCHEMA     = "create_schema"
	OP_DROP_SCHEMA       = "drop_schema"
	OP_DESCRIBE_SCHEMA   = "describe_schema"
	OP_CREATE_TABLE      = "create_table"
	OP_DROP_TABLE        = "drop_table"
	OP_INSERT            = "insert"
	OP_UPDATE            = "update"
	OP_DELETE            = "delete"
	OP_SEARCH_BY_HASH    = "search_by_hash"
	OP_SEARCH_BY_VALUE   = "search_by_value"
	OP_SQL               = "sql"
	OP_CSV_DATA_LOAD     = "csv_data_load"
	OP_CSV_FILE_LOAD     = "csv_file_load"
	OP_CSV_URL_LOAD      = "csv_url_load"
	OP_DESCRIBE_TABLE    = "describe_table"
	OP_GET_JOB           = "get_job"
	OP_SEARCH_JOBS       = "search_jobs_by_start_date"
	OP_ADD_ROLE          = "add_role"
	OP_ALTER_ROLE        = "alter_role"
	OP_DROP_ROLE         = "drop_role"
	OP_LIST_ROLES        = "list_roles"
	OP_ADD_USER          = "add_user"
	OP_ALTER_USER        = "alter_user"
	OP_DROP_USER         = "drop_user"
	OP_USER_INFO         = "user_info"
	OP_LIST_USERS        = "list_users"
)

type operation struct {
	Operation       string        `json:"operation"`
	Schema          string        `json:"schema,omitempty"`
	Table           string        `json:"table,omitempty"`
	HashAttribute   string        `json:"hash_attribute,omitempty"`
	HashValues      interface{}   `json:"hash_values,omitempty"`
	Records         interface{}   `json:"records,omitempty"`
	GetAttributes   AttributeList `json:"get_attributes,omitempty"`
	SearchAttribute Attribute     `json:"search_attribute,omitempty"`
	SearchValue     interface{}   `json:"search_value,omitempty"`
	SQL             string        `json:"sql,omitempty"`
	Action          string        `json:"action,omitempty"`
	Data            string        `json:"data,omitempty"`
	FilePath        string        `json:"file_path,omitempty"`
	CSVURL          string        `json:"csv_url,omitempty"`
	ID              string        `json:"id,omitempty"`
	FromDate        string        `json:"from_date,omitempty"`
	ToDate          string        `json:"to_date,omitempty"`
	Role            string        `json:"role,omitempty"`
	Permission      Permission    `json:"permission,omitempty"`
	Username        string        `json:"username,omitempty"`
	Password        string        `json:"password,omitempty"`
	Active          *bool         `json:"active,omitempty"`
}

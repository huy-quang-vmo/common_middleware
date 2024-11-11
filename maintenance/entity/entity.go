package entity

type ServiceStatus string

var (
	StatusMaintenance ServiceStatus = "Maintenance"
	StatusActive      ServiceStatus = "Active"
	StatusInactive    ServiceStatus = "Inactive"
)

type ServiceManagement struct {
	ID     interface{}   `json:"id" mongo:"_id" pg:"id"`
	Status ServiceStatus `json:"status" mongo:"status" pg:"status"`

	tableName struct{} `pg:"service_managements,alias:service_managements"` //nolint:all
}

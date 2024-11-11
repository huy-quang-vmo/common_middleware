package entity

type ServiceStatus string

var (
	StatusMaintenance ServiceStatus = "maintenance"
	StatusActive      ServiceStatus = "active"
	StatusInactive    ServiceStatus = "inactive"
)

type ServiceManagement struct {
	ID     int           `json:"id"`
	Status ServiceStatus `json:"status"`

	tableName struct{} `pg:"service_managements,alias:service_managements"` //nolint:all
}

package maintenance

import "net/http"

type MaintenanceMiddleware struct {
	maintenanceStatus GetMaintenanceStatus
}

func NewMaintenanceMiddleware(maintenanceStatus GetMaintenanceStatus) *MaintenanceMiddleware {
	return &MaintenanceMiddleware{maintenanceStatus: maintenanceStatus}
}

type GetMaintenanceStatus interface {
	IsMaintenance() (bool, error)
}

func (c *MaintenanceMiddleware) MaintenanceStatus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isMaintenanceMode, err := c.maintenanceStatus.IsMaintenance()
		if err != nil {
			http.Error(w, "Failed to check maintenance_test status", http.StatusInternalServerError)
			return
		}
		if isMaintenanceMode {
			http.Error(w, "Service is under maintenance_test", http.StatusServiceUnavailable)
			return
		}
		next.ServeHTTP(w, r)
	})
}

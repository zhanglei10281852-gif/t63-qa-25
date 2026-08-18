package httpapi

import (
	"net/http"
	"testing"
)

func TestRecordedMileageIsVisibleOnVehicle(t *testing.T) {
	h := newAPIHarness(t)
	created := h.createVehicle(t, "601")
	id := requireString(t, created, "id")
	h.request(t, http.MethodPost, "/api/v1/fuel", map[string]any{"vehicle_id": id, "fuel_type": "diesel", "quantity": 20, "cost_cents": 15000, "odometer_km": 150, "station_code": "H-01", "recorded_at": h.now}, http.StatusCreated)
	page := h.request(t, http.MethodGet, "/api/v1/vehicles?limit=10&q=601", nil, http.StatusOK)
	item := page["items"].([]any)[0].(map[string]any)
	if item["odometer_km"].(float64) != 150 {
		t.Fatalf("vehicle mileage=%v", item["odometer_km"])
	}
}

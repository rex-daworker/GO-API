# GO-API Parking Management System

## Setup
1. Clone the repo
2. Navigate to `api01`
3. Run `go mod tidy`
4. Start server: `go run cmd/api/main.go`

## Endpoints
- `POST /parking` — Create event
- `PUT /parking/{id}` — Update event
- `DELETE /parking/1` — Delete event
- `GET /parking/all` — List all
- `GET /parking/all?page=2&slot_id=A1` — Filtered list

## Auth
- Username: `admin_oghenerobo`
- Password: `StrongPass!2025`

## Testing
```bash
go test ./internal/api/service/parking

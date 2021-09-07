package pkg

//DNSQueryParams represents query params needed to make request to DNS
type DNSQueryParams struct {
	XCoord   float64 `json:"x"`
	YCoord   float64 `json:"y"`
	ZCoord   float64 `json:"z"`
	Velocity float64 `json:"velocity"`
}

//DNSResponse represents response for location retrieval
type DNSResponse struct {
	Location float64 `json:"loc"`
}

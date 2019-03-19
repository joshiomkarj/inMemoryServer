package runtime

// Register request struct
type RegisterRequest struct {
	ID     string `json:"id,omitempty"`
	VMName string `json:"vmname,omitempty"`
	VMID   string `json:"vmid,omitempty"`
	CPU    string `json:"cpuutilization,omitempty"`
}

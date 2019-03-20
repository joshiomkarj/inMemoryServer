package runtime

// Register request struct
type RegisterRequest struct {
	VMName string `json:"vmname,omitempty"`
	VMID   string `json:"id,omitempty"`
	CPU    string `json:"cpuutilization,omitempty"`
}

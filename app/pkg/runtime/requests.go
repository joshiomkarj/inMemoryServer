package runtime

// Register speedyn user request struct
type RegisterRequest struct {
	ID     string `json:"id"`
	VMName string `json:"vmname"`
	VMID   string `json:"vmid"`
	CPU    string `json:"cpuutilization"`
}

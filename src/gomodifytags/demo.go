package main

type Server struct {
	Name        string `json:"name,omitempty"`
	Port        int    `json:"port,omitempty"`
	EnableLogs  bool   `json:"enable_logs,omitempty"`
	BaseDomain  string `json:"base_domain,omitempty"`
	Credentials struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	} `json:"credentials,omitempty"`
}

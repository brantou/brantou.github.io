package main


type Server struct {
  Name        string
  Port        int
  EnableLogs  bool
  BaseDomain  string
  Credentials struct {
    Username string
    Password string
  }
}

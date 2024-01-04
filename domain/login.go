package domain

type LoginRequest struct {
	NIM				string		`json:"nim" validator:"required"`
	Password		string		`json:"password" validator:"required"`
}

type LoginResponse struct {
	AccessToken		string		`json:"access_token"`
	RefreshToken	string		`json:"refresh_token"`
}

type LoginUsecase interface {
	ValidateMahasiswa(mahasiswa *Mahasiswa, request *LoginRequest) error 
	CreateFirebaseToken(mahasiswa *Mahasiswa, expired int) (string, error)
}
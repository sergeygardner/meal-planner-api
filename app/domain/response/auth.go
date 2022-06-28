package response

type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthConfirmation struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Status  int    `protobuf:"varint,2,opt,name=status,proto3" json:"status"`
}

package Auth

type Auth struct {
	Login    string `protobuf:"string,1,opt,name=login"`
	Password string `protobuf:"string,2,opt,name=password"`
}

func (a *Auth) WriteDB() error {
	return nil
}

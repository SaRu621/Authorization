package modelReg

type Reg struct {
	Login 	 string `protobuf:"string,1,opt,name=login"`
	Password string `protobuf:"string,2,opt,name=password"`
	Email    string `protobuf:"string,3,opt,name=email"`
	Name 	 string `protobuf:"string,4,opt,name=name"`
	Surname  string `protobuf:"string,5,opt,name=surname"`
}


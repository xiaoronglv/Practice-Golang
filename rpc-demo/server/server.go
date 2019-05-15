package server

type Args struct {
	A, B int
}

type Arith int

func (a *Arith) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func (a *Arith) Sub(args *Args, reply *int) error {
	*reply = args.A - args.B
	return nil
}

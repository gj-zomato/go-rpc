package mathFunc

import "errors"

type Args struct {
	A, B int
}

// Arith provides arithmetic methods
type Arith struct{}

// Add adds two numbers and stores the result in reply
func (a *Arith) Add(args *Args, reply *int) error {
	if args == nil {
		return errors.New("nil arguments")
	}
	*reply = args.A + args.B
	return nil
}

// Subtract subtracts two numbers and stores the result in reply
func (s *Arith) Subtract(args *Args, reply *int) error {
	if args == nil {
		return errors.New("nil arguments")
	}
	*reply = args.A - args.B
	return nil
}

// Multiply multiplies two numbers and stores the result in reply
func (m *Arith) Multiply(args *Args, reply *int) error {
	if args == nil {
		return errors.New("nil arguments")
	}
	*reply = args.A * args.B
	return nil
}

// Divide divides two numbers and stores the result in reply
func (d *Arith) Divide(args *Args, reply *int) error {
	if args == nil {
		return errors.New("nil arguments")
	}
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*reply = args.A / args.B
	return nil
}

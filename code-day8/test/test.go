package main

import "fmt"

type Callback func()

type LoginCallback func() bool

type Command struct {
	Name     string
	Callback Callback
}

type manager struct {
	loginCallback LoginCallback
	cmds          map[int]*Command
}

func newManager() *manager {
	return &manager{
		cmds: make(map[int]*Command),
	}
}

var mgr = newManager()

func main() {
	mgr.loginCallback = func() bool {
		return false
	}

	fmt.Printf("%T,%#v\n", mgr.loginCallback, mgr.loginCallback)
	mgr.loginCallback()
	fmt.Printf("%T,%#v\n", mgr.loginCallback(), mgr.loginCallback())

	if !mgr.loginCallback() {
		fmt.Println("loginCallback ")
	}

}

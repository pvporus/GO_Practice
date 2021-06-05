package main

import "net/http/pprof"

func (a *App) mapProfileRoutes() {
	//handlers for profiles
	a.Router.HandleFunc("/debug/pprof/", pprof.Index)
	a.Router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	a.Router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	a.Router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	a.Router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	a.Router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	a.Router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	a.Router.Handle("/debug/pprof/block", pprof.Handler("block"))
	a.Router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
}

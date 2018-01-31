package main

type application struct {
}

func newApplication(opts ...appOption) *application {
	return &application{}
}

type appOption func(*application)

var renderToMemory appOption = func(a *application) {}

func (a *application) buildSite() {}

func (a *application) startServer() {}

func (a *application) createNewSite(name string) {}

func (a *application) addNewContent(name string) {}

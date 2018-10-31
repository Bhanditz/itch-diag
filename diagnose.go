package main

import (
	"runtime"
)

// Diagnose runs a battery of tests.
func (a *App) Diagnose() {
	a.Debugf("Running diagnostics...")

	a.Eval(`
		window.external.invoke(JSON.stringify({
			UserAgent: window.navigator.userAgent
		}));
	`)
	msg := &UserAgentMessage{}
	a.Receive(&msg)
	a.Infof("User-Agent is: %s", msg.UserAgent)

	switch runtime.GOOS {
	case "windows":
		a.DiagnoseWindows()
	}

	a.Test("Diagnosing itch app dependencies", a.DiagnoseAppData)

	a.Debugf("All done!")
}

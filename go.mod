module github.com/pinpt/agent.next.github

go 1.14

require (
	github.com/pinpt/agent.next v0.0.0-20200610125456-def776246343
	golang.org/x/sys v0.0.0-20200610111108-226ff32320da // indirect
)

// TODO: this is only set while we're in rapid dev. once we get out of that we should remove
replace github.com/pinpt/agent.next => ../agent.next

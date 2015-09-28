/*
Package communicate is the communication organ.

This is not fixed yet, but we will try to follow ros.org messages format :
see http://wiki.ros.org/msg
*/
package communicate

// Alive ask a robot if he is alive.
//
// Return a single string line status message:
// "status alive"
// "status energy saving mode"
//
// "status dead" is obviously not returned. It is assumed after a time period without a reply.
func Alive() string {
	return "status alive"
}

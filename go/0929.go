package main

import (
	"strings"
)

// https://leetcode.com/problems/unique-email-addresses/
func numUniqueEmails(emails []string) int {
	u := make(map[string]struct{})

	for _, email := range emails {
		parts := strings.SplitN(email, "@", 2)
		local, domain := parts[0], parts[1]

		local = strings.ReplaceAll(local, ".", "")
		local = strings.SplitN(local, "+", 2)[0]

		u[local+"@"+domain] = struct{}{}
	}

	return len(u)
}

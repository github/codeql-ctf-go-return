# codeql-ctf-go-return
Extension project for the Go and Don't Return CodeQL CTF

This repository is intended for use by contestants in the March 2021 CodeQL CTF. If you are unfamiliar, first read the [https://securitylab.github.com/ctf/go-and-dont-return](contest documentation).

This is a set of example programs that exhibit coding mistakes similar to CVE-2020-11012 found in MinIO, but which are in various ways trickier to accurately detect than the original MinIO bug. See the contest documentation linked above for instructions on building a CodeQL database for these examples and identifying bugs in them.

If you notice any errors or omissions that don't appear related to the intentional bugs that are the target of the contest, see CONTRIBUTING.md for instructions on submitting pull requests.

These examples are licensed under CC0. See the LICENSE file for full details.

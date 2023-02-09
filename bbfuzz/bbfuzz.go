// Package bbfuzz tries to match strings that aren't "bb".
//
// The package is inspired by [this Tweet], which asks:
//
//	Write a regular expression which matches any string except the string "bb".
//
// [this Tweet]: https://twitter.com/qntm/status/1623418949445595138
package bbfuzz

import "regexp"

var (
	// Reference answer.
	regexp0 = regexp.MustCompile(`^(|[\S\s]|[\S\s]{3,}|[^b][\S\s]|[\S\s][^b])$`)

	// Various users' submissions.
	regexp1  = regexp.MustCompile(`^(.|.{3,}|[^b].|.[^b])$`)
	regexp2  = regexp.MustCompile(`([^b].*)|(b[^b].*)|^$`)
	regexp3  = regexp.MustCompile(`[^b]*b[^b]*|[^b]*bbb+[^b]*|([^b]+(b|bbb+)[^b]+)*`)
	regexp4  = regexp.MustCompile(`(b|bbb+)?([^b]+(b|bbb+)[^b]*)*|([^b]*(b|bbb+)[^b]+)*(b|bbb+)?`)
	regexp5  = regexp.MustCompile(`[^b].*|b([^b].*|b.+)|b?`)
	regexp6  = regexp.MustCompile(`^[^b].*|b[^b].*|bb.+$`)
	regexp7  = regexp.MustCompile(`^.?([^b].*)?$`)
	regexp8  = regexp.MustCompile(`^(|.|[^b].|.[^b]|...+)$`)
	regexp9  = regexp.MustCompile(`^(b[^b]|[^b]b|...+)$`)
	regexp10 = regexp.MustCompile(`^(b[^b]|[^b]b|.|...+)$`)
	regexp11 = regexp.MustCompile(`^[^b]?[^b]?.*`)
	regexp12 = regexp.MustCompile(`^(.?|.[^b]|[^b].|...+)$`)
	regexp13 = regexp.MustCompile(`^(b|[^b]{2}.*|bbb.*|bb[^b].*|[^b].*|b[^b].*)$`)
)

func Regexp0(s string) bool  { return regexp0.MatchString(s) }
func Regexp1(s string) bool  { return regexp1.MatchString(s) }
func Regexp2(s string) bool  { return regexp2.MatchString(s) }
func Regexp3(s string) bool  { return regexp3.MatchString(s) }
func Regexp4(s string) bool  { return regexp4.MatchString(s) }
func Regexp5(s string) bool  { return regexp5.MatchString(s) }
func Regexp6(s string) bool  { return regexp6.MatchString(s) }
func Regexp7(s string) bool  { return regexp7.MatchString(s) }
func Regexp8(s string) bool  { return regexp8.MatchString(s) }
func Regexp9(s string) bool  { return regexp9.MatchString(s) }
func Regexp10(s string) bool { return regexp10.MatchString(s) }
func Regexp11(s string) bool { return regexp11.MatchString(s) }
func Regexp12(s string) bool { return regexp12.MatchString(s) }
func Regexp13(s string) bool { return regexp13.MatchString(s) }

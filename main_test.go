package br_utils

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestCpfFormat(t *testing.T) {
	res := FormatCpfCnpj("42asd22$$3485898++----")
	if res != "422.234.858-98" {
		t.Fail()
	}
}

func TestCnpjFormat(t *testing.T) {
	res := FormatCpfCnpj("---6902a08540d$$00136)(")
	if res != "69.020.854/0001-36" {
		t.Fail()
	}
}

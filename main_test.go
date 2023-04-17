package br_utils

import (
	"github.com/Rafaellinos/br-utils/src"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestCpfCnpjFormat(t *testing.T) {
	res := src.FormatCpfCnpj("42asd22$$3485898++----")
	if res != "422.234.858-98" {
		t.Fail()
	}
}

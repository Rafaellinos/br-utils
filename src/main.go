package src

import "regexp"

func FormatCpfCnpj(cpfCnpj string) string {
	var re = regexp.MustCompile(`[^0-9]`)
	str := re.ReplaceAllString(cpfCnpj, "")

	switch size := len([]rune(str)); size {
	case 11:
		return FormatCpf(str)
	case 14:
		return FormatCnpj(str)
	}
	return ""
}

func FormatCpf(cpf string) string {
	return cpf[:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:]
}

func FormatCnpj(cnpj string) string {
	return cnpj[0:2] + "." + cnpj[2:5] + "." + cnpj[5:8] + "/" + cnpj[8:12] + "-" + cnpj[12:14]
}

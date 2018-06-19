package cnj

import (
	"fmt"
	"regexp"
	"errors"
	"strings"
)

var reCnj *regexp.Regexp
var ErrInvalidNumber error = errors.New("Numero inválido")
var ErrInvalidDigit error = errors.New("Dígito verificador inválido")

func init() {
	reCnj = regexp.MustCompile("[^0-9]")
}


// Valida dígitos de verificação de número de processo no padrão do CNJ.
// O parâmetro num deve conter o número, podendo ou não conter caracteres de
// separação como pontos, traços, espaços, etc.
func Valida(num string) error {
	var s string
	var n, d int64
	var err error

	s = reCnj.ReplaceAllString(num, "")
	s = strings.TrimLeft(s,"0")
	if len(s)==0 || len(s)>20 {
		return ErrInvalidNumber
	}

	if len(s) > 11 {
		_, err = fmt.Sscanf(s[:len(s)-13] + s[len(s)-11:] + "00","%d",&n)
		if err != nil {
			return err
		}
		if len(s) == 12 {
			_, err = fmt.Sscanf(s[:1],"%d",&d)
			if err != nil {
				return err
			}
		} else {
			_, err = fmt.Sscanf(s[len(s)-13:len(s)-11],"%d",&d)
			if err != nil {
				return err
			}
		}
	} else {
		_, err = fmt.Sscanf(s + "00","%d",&n)
		if err != nil {
			return err
		}
		d = 0
	}

	if (98 - (n%97)) == d {
		return nil
	}

	return ErrInvalidDigit
}


// Remove todos os caracteres diferentes de dígitos e insere zeros à esquerda para
// garantir que o número tenha 20 digitos
// Não faz validação do número, assumindo que o número fornecido já esteja validado
func Normaliza(num string) string {
	var n uint
	fmt.Sscanf(reCnj.ReplaceAllString(num, ""),"%d",&n)
	return fmt.Sprintf("%020d", n)
}



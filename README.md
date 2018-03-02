# cnj
Valida dígitos de verificação de número de processo no padrão do Conselho Nacional de Justiça (Resolução 65/2008 http://www.cnj.jus.br/busca-atos-adm?documento=2748).

```Go
  var err error

  if err = Valida(); err!=nil {
    if err == ErrInvalidNumber {
      // Trata numeros inválidos
    } else if err == ErrInvalidDigit {
       // Trata dígitos inválidos
    }
  }

```

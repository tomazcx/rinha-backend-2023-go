package person

import "errors"

var (
	ErrNicknameAlreadyRegistered = errors.New("Erro: Apelido já registrado")
	ErrPersonNotFound = errors.New("Erro: Pessoa não encontrada")
)

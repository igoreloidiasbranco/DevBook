package seguranca

import "golang.org/x/crypto/bcrypt"

func InserirHashNaSenha(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func CompararSenhaComHash(hashDaSenha []byte, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashDaSenha), []byte(senha))
}

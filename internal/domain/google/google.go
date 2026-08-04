package google

type GoogleAuthInput struct {
	IDToken string `json:"id_token"`
}

type GoogleTokenPayload struct {
	// Iss é o identificador do emissor para o emissor da resposta.
	// Normalmente https://accounts.google.com, mas accounts.google.com também é retornado para implementações legadas.
	Iss string `json:"iss"`

	// Sub é um identificador para o usuário, exclusivo entre todas as Contas do Google e nunca reutilizado.
	// Use sub no seu aplicativo como a chave de identificador único do usuário (comprimento máximo de 255 caracteres).
	Sub string `json:"sub"`

	// Azp é o identificador do cliente do apresentador autorizado, obtido no Console do Google Cloud.
	// Essa declaração só é necessária quando a parte que solicita o token de ID não é a mesma que o público-alvo.
	Azp string `json:"azp,omitempty"`

	// Aud é o público-alvo a que o token de ID se destina.
	// Esse é o identificador do cliente do seu aplicativo, obtido no console do Google Cloud.
	Aud string `json:"aud"`

	// Iat representa a hora em que o token de ID foi emitido, em segundos (tempo da época Unix).
	Iat int64 `json:"iat"`

	// Exp representa o prazo de validade em ou após o qual o token de ID não pode ser aceito, em segundos (época Unix).
	Exp int64 `json:"exp"`

	// Nonce é o valor do nonce fornecido pelo seu app na solicitação de autenticação.
	// Proteja-se contra ataques de repetição apresentando esse valor apenas uma vez.
	Nonce string `json:"nonce,omitempty"`

	// AuthTime indica o momento em que a autenticação do usuário ocorreu, em segundos decorridos desde a época Unix.
	AuthTime int64 `json:"auth_time,omitempty"`

	// AtHash é o hash do token de acesso, fornecendo validação de que ele está vinculado ao token de identidade.
	AtHash string `json:"at_hash,omitempty"`

	// Email é o endereço de e-mail do usuário.
	// Aviso: não use como identificador principal; use sempre o campo Sub.
	Email string `json:"email,omitempty"`

	// EmailVerified é verdadeiro se o endereço de e-mail do usuário tiver sido verificado; caso contrário, falso.
	EmailVerified bool `json:"email_verified,omitempty"`

	// Name é o nome completo do usuário em formato de exibição.
	Name string `json:"name,omitempty"`

	// Picture é a URL da foto do perfil do usuário.
	Picture string `json:"picture,omitempty"`

	// GivenName é o primeiro nome do usuário.
	GivenName string `json:"given_name,omitempty"`

	// FamilyName é o(s) sobrenome(s) do usuário.
	FamilyName string `json:"family_name,omitempty"`

	// Hd é o domínio associado à organização do Google Workspace ou do Cloud do usuário.
	// A ausência desta declaração indica que a conta não pertence a um domínio hospedado pelo Google.
	Hd string `json:"hd,omitempty"`
}

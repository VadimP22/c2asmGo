package casmutility


type Token struct {
	tokenType  string
	tokenValue string
}


func (t *Token) GetType() string {
	return t.tokenType
}


func (t *Token) GetValue() string {
	return t.tokenValue
}


func NewToken(tokenType, tokenValue string) Token {
	var tok Token
	tok.tokenType = tokenType
	tok.tokenValue = tokenValue
	return tok
}
package docs

import "cryptoStringRedisCache/internal/infrastructure/dto"

//go:generate swagger generate spec -i ../static/swagger.json -o ../static/swagger.json --scan-models

// swagger:route POST /api/crypto crypto Request
// Шифрует введеную строку по алгоритмам MD5 и SHA256
//
// responses:
//     200: Response

// swagger:parameters Request
type requestCrypto struct {
	// required: true
	// in: body
	Body dto.Request
}

// swagger:response Response
type responseCrypto struct {
	// required: true
	// in: body
	Body dto.Response
}

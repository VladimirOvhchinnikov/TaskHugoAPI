package main

// swagger:route POST /login login handleLogin
// handleLogin авторизует пользователя и возвращает JWT токен.
// При успешной авторизации возвращает токен, иначе сообщение об ошибке.
//
// Параметры:
// + name: body
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/LogPas"
//
// Ответы:
//   200: jwtResponse
//   400: badRequestResponse
//   500: internalServerErrorResponse
//   401: unauthorizedResponse

// LogPas определяет структуру входных данных для авторизации пользователя.
// swagger:model
type LogPas struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JwtResponse определяет структуру ответа с JWT токеном.
// swagger:response jwtResponse
type JwtResponse struct {
	// in:body
	Body struct {
		Token string `json:"token"`
	}
}

// badRequestResponse описывает ответ на неверный запрос.
// swagger:response badRequestResponse
type badRequestResponse struct {
	// in:body
	Body struct {
		Message string `json:"message"`
	}
}

// internalServerErrorResponse описывает ответ сервера на внутреннюю ошибку.
// swagger:response internalServerErrorResponse
type internalServerErrorResponse struct {
	// in:body
	Body struct {
		Message string `json:"message"`
	}
}

// unauthorizedResponse описывает ответ сервера на неавторизованный запрос.
// swagger:response unauthorizedResponse
type unauthorizedResponse struct {
	// in:body
	Body struct {
		Message string `json:"message"`
	}
}

// swagger:route GET /api/address/search address handleSearch
// handleSearch выполняет поиск адреса.
// responses:
//
//	200: searchResponse
//
// swagger:response searchResponse
type searchResponse struct {
	// in:body
	Body struct {
		Message string `json:"message"`
	}
}

// swagger:route GET /api/address/geocode address handleGeocode
// handleGeocode выполняет геокодирование адреса.
// responses:
//
//	200: geocodeResponse
//
// swagger:response geocodeResponse
type geocodeResponse struct {
	// in:body
	Body struct {
		Location string `json:"location"`
	}
}

// swagger:route GET /api address handleAll
// handleAll возвращает общую информацию.
// responses:
//
//	200: allResponse
//
// swagger:response allResponse
type allResponse struct {
	// in:body
	Body struct {
		Greeting string `json:"greeting"`
	}
}

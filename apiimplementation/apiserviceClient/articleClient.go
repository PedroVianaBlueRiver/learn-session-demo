package apiserviceclient

type apiUrl struct {
	Url string
}

func NewApi() *apiUrl {
	return &apiUrl{Url: "https://65e0b4ced3db23f76249e825.mockapi.io/"}
}

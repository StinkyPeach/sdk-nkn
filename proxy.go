package nkn

// This is just a random thing

var (
    HttpProxyUrl = ""
)

// no lock required

func GetProxyUrl() string {
    return HttpProxyUrl
}

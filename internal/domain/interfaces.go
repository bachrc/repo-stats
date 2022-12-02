package domain

type PongContent struct {
	Pong string
}

type ProfileFetcher interface {
	Pong() PongContent
}

type ProfileStatsDomain struct {
	fetcher ProfileFetcher
}

func NewProfileStats(fetcher ProfileFetcher) ProfileStatsDomain {
	return ProfileStatsDomain{fetcher: fetcher}
}

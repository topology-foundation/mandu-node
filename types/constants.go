package types

const (
	AccountAddressPrefix string = "mandu"
	AppName              string = "mandu"
	TokenDenom           string = "mandu"

	DefaultEpochSize      int64 = 10      // 60 seconds (~6s block time)
	RewardClaimWindow     int64 = 60 * 24 // 1 day (considering the default epoch size)
	ChallengeAnswerPeriod int64 = 1       // answer within 1 epoch
)

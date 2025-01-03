package types

const (
	// ModuleName defines the module name
	ModuleName = "subscription"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_subscription"

	SubscriptionRequestKeyPrefix          = "SubscriptionRequest/value"
	SubscriptionRequestRequesterKeyPrefix = "SubscriptionRequest/Requester/value"
	SubscriptionKeyPrefix                 = "Subscription/value"
	SubscriptionSubscriberKeyPrefix       = "Subscription/Subscriber/value"
	ProgressKeyPrefix                     = "Progress/value"
	ProgressObfuscatedKeyPrefix           = "Progress/Obfuscated/value"
	ProgressSizeKeyPrefix                 = "Progress/Size/value"
	HashSubmissionBlockKeyPrefix          = "HashSubmissionBlock/value"
)

var ParamsKey = []byte("p_subscription")

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetProviderStoreKey returns the key for the provider store for the given provider.
func GetSubscriberStoreKey(subscriber string) []byte {
	return KeyPrefix(SubscriptionSubscriberKeyPrefix + "/" + subscriber)
}

// GetRequesterStoreKey returns the key for the requester store for the given requester.
func GetRequesterStoreKey(requester string) []byte {
	return KeyPrefix(SubscriptionRequestRequesterKeyPrefix + "/" + requester)
}

func GetProgressSizeStoreKey(subscription string) []byte {
	return KeyPrefix(ProgressSizeKeyPrefix + "/" + subscription)
}

func GetHashSubmissionBlockStoreKey(provider string) []byte {
	return KeyPrefix(HashSubmissionBlockKeyPrefix + "/" + provider)
}

package validation

import (
	"testing"
)

func TestValidateNonEmptyString(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"Valid non-empty string", "Hello", false},
		{"Empty string", "", true},
		{"Whitespace only", "   ", true},
		{"Valid string with spaces", "  Hello  ", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateNonEmptyString(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("ValidateNonEmptyString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePositiveAmount(t *testing.T) {
	tests := []struct {
		name    string
		amount  int64
		wantErr bool
	}{
		{"Valid positive amount", 100, false},
		{"Zero amount", 0, true},
		{"Large positive amount", 1844674407379551615, false}, // Max uint64 value
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePositiveAmount(tt.amount); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePositiveAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		wantErr bool
	}{
		{"Valid bech32 address", "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g", false},
		{"Empty address", "", true},
		{"Invalid bech32 address", "cosmos1invalid", true},
		{"Wrong prefix", "osmo1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateAddress(tt.address); (err != nil) != tt.wantErr {
				t.Errorf("ValidateAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateBlockRange(t *testing.T) {
	tests := []struct {
		name       string
		startBlock int64
		endBlock   int64
		wantErr    bool
	}{
		{"Valid block range", 100, 200, false},
		{"Start block equals end block", 100, 100, true},
		{"Start block greater than end block", 200, 100, true},
		{"Zero start block", 0, 100, false},
		{"Large block numbers", 1844674407379551614, 1844674407379551615, false}, // Max uint64 value - 1 and Max uint64 value
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateBlockRange(tt.startBlock, tt.endBlock); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBlockRange() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateEpochRange(t *testing.T) {
	tests := []struct {
		name       string
		startEpoch int64
		endEpoch   int64
		wantErr    bool
	}{
		{"Valid epoch range", 100, 200, false},
		{"Start epoch equals end epoch", 100, 100, true},
		{"Start epoch greater than end epoch", 200, 100, true},
		{"Zero start epoch", 0, 100, false},
		{"Large epoch numbers", 1844674407379551614, 1844674407379551615, false}, // Max uint64 value - 1 and Max uint64 value
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateBlockRange(tt.startEpoch, tt.endEpoch); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBlockRange() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

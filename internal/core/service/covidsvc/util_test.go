package covidsvc_test

import (
	"math/rand"
	"strings"

	"covid/internal/dtos"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomRecords(n int) []dtos.CovidRecord {
	records := make([]dtos.CovidRecord, n)

	for i := range records {
		province := RandomString(10)
		age := int(RandomInt(1, 100))

		// 10% chance to set Province to empty string
		if rand.Float64() <= 0.1 {
			province = ""
		}

		// 10% chance to set Age to 0
		if rand.Float64() <= 0.1 {
			age = 0
		}

		records[i] = dtos.CovidRecord{
			Province: province,
			Age:      age,
		}
	}

	return records
}

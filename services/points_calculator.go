package services

import (
	"math"
	"receipt-processor-challenge-fetch/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CalculatePoints(receipt models.Receipt) int64 {
	var points int64

	// Rule 1: One point for every alphanumeric character in the retailer name
	alphanumeric := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += int64(len(alphanumeric.FindAllString(receipt.Retailer, -1)))

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == math.Floor(total) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if math.Mod(total*100, 25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt
	points += int64((len(receipt.Items) / 2) * 5)

	// Rule 5: If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int64(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}

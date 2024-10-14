package main

import (
	"fmt"
	"time"
)

// Function to generate a license
func generateLicense(userID string) (string, error) {
	// Creating a new license with a 30-day trial
	license := licenser.License{
		UserID:      userID,
		Expires:     time.Now().Add(30 * 24 * time.Hour), // 30 days trial
		LicenseType: "trial",
	}

	// Generate the license key using a secret key
	key, err := licenser.GenerateLicenseKey(license, "your-secret-key")
	if err != nil {
		return "", err
	}

	return key, nil
}

// Function to validate a license
func validateLicense(licenseKey string) (bool, error) {
	// Validate the license key using the same secret key
	license, err := licenser.ValidateLicenseKey(licenseKey, "your-secret-key")
	if err != nil {
		return false, err
	}

	// Check if the license has expired
	if license.Expires.Before(time.Now()) {
		return false, nil // License is expired
	}

	return true, nil
}

func main() {
	// Example user ID
	userID := "user123"

	// Generate a license for the user
	licenseKey, err := generateLicense(userID)
	if err != nil {
		fmt.Println("Error generating license:", err)
		return
	}
	fmt.Println("Generated License Key:", licenseKey)

	// Validate the generated license
	valid, err := validateLicense(licenseKey)
	if err != nil {
		fmt.Println("Error validating license:", err)
		return
	}

	if valid {
		fmt.Println("License is valid.")
	} else {
		fmt.Println("License is invalid or expired.")
	}
}

package utils

import "testing"

func TestIsValidateEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "Valid email basic format",
			email:    "test@example.com",
			expected: true,
		},
		{
			name:     "Valid email with subdomain",
			email:    "test@sub.example.com",
			expected: true,
		},
		{
			name:     "Valid email with numbers",
			email:    "test123@example.com",
			expected: true,
		},
		{
			name:     "Invalid email - missing @",
			email:    "testexample.com",
			expected: false,
		},
		{
			name:     "Invalid email - missing domain",
			email:    "test@",
			expected: false,
		},
		{
			name:     "Invalid email - spaces",
			email:    "test @example.com",
			expected: false,
		},
		{
			name:     "Invalid email - empty string",
			email:    "",
			expected: false,
		},
		{
			name:     "Invalid email - just string",
			email:    "teststring",
			expected: false,
		},
		{
			name:     "Invalid email - multiple @",
			email:    "test@test@example.com",
			expected: false,
		},
		{
			name:     "Invalid email - special chars",
			email:    "test!@example.com",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidateEmail(tt.email)
			if result != tt.expected {
				t.Errorf("isValidateEmail(%q) = %v, want %v", tt.email, result, tt.expected)
			}
		})
	}
}

func TestIsValidNickName(t *testing.T) {
    tests := []struct {
        name     string
        nickname string
        expected bool
    }{
        {
            name:     "Valid nickname basic",
            nickname: "john123",
            expected: true,
        },
        {
            name:     "Valid nickname with period",
            nickname: "john.doe",
            expected: true,
        },
        {
            name:     "Invalid - too short",
            nickname: "joe",
            expected: false,
        },
        {
            name:     "Invalid - starts with period",
            nickname: ".john",
            expected: false,
        },
        {
            name:     "Invalid - ends with period",
            nickname: "john.",
            expected: false,
        },
        {
            name:     "Invalid - consecutive periods",
            nickname: "john..doe",
            expected: false,
        },
        {
            name:     "Invalid - special characters",
            nickname: "john@doe",
            expected: false,
        },
        {
            name:     "Invalid - empty string",
            nickname: "",
            expected: false,
        },
        {
            name:     "Invalid - too long",
            nickname: "johndoejohndoejohndoe",
            expected: false,
        },
        {
            name    : "Invalid - just string",
            nickname: "johndo",
            expected: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := isValidNickName(tt.nickname)
            if result != tt.expected {
                t.Errorf("isValidNickName(%q) = %v, want %v", tt.nickname, result, tt.expected)
            }
        })
    }
}

package main

import (
	"fmt"
	"testing"
)

func TestVerifyJSONWithInvalidResource(t *testing.T) {
	jsonData := []byte(`{
		"PolicyName": "root",
		"PolicyDocument": {
			"Version": "2012-10-17",
			"Statement": [
				{
					"Sid": "IamListAccess",
					"Effect": "Allow",
					"Action": [
						"iam:ListRoles",
						"iam:ListUsers"
					],
					"Resource": "*"
				}
			]
		}
	}`)

	if verifyJSON(jsonData) {
		t.Error("Expected JSON file to be invalid, but got valid")
		return
	}
	fmt.Println("TestVerifyJSONWithInvalidResource passed")
}

func TestVerifyJSONWithValidResource(t *testing.T) {
	jsonData := []byte(`{
		"PolicyName": "root",
		"PolicyDocument": {
			"Version": "2012-10-17",
			"Statement": [
				{
					"Sid": "IamListAccess",
					"Effect": "Allow",
					"Action": [
						"iam:ListRoles",
						"iam:ListUsers"
					],
					"Resource": "arn:aws:iam::123456789012:role/*"
				}
			]
		}
	}`)

	if !verifyJSON(jsonData) {
		t.Error("Expected JSON file to be valid, but got invalid")
		return
	}
	fmt.Println("TestVerifyJSONWithValidResource passed")
}

func TestVerifyJSONWithMultipleStatements(t *testing.T) {
	jsonDataInvalid := []byte(`{
		"PolicyName": "root",
		"PolicyDocument": {
			"Version": "2012-10-17",
			"Statement": [
				{
					"Sid": "IamListAccess",
					"Effect": "Allow",
					"Action": [
						"iam:ListRoles",
						"iam:ListUsers"
					],
					"Resource": "arn:aws:iam::123456789012:user/username"
				},
				{
					"Sid": "IamReadAccess",
					"Effect": "Allow",
					"Action": "iam:GetUser",
					"Resource": "*"
				}
			]
		}
	}`)

	jsonDataValid := []byte(`{
		"PolicyName": "root",
		"PolicyDocument": {
			"Version": "2012-10-17",
			"Statement": [
				{
					"Sid": "IamListAccess",
					"Effect": "Allow",
					"Action": [
						"iam:ListRoles",
						"iam:ListUsers"
					],
					"Resource": "arn:aws:iam::123456789012:user/username"
				},
				{
					"Sid": "IamReadAccess",
					"Effect": "Allow",
					"Action": "iam:GetUser",
					"Resource": "arn:aws:iam::121236789012:user/*"
				}
			]
		}
	}`)

	if verifyJSON(jsonDataInvalid) {
		t.Error("Expected JSON file to be invalid, but got valid")
		return

	}
	if !verifyJSON(jsonDataValid) {
		t.Error("Expected JSON file to be valid, but got invalid")
		return

	}
	fmt.Println("TestVerifyJSONWithMultipleStatements passed")

}

func TestVerifyJSONWithMultipleResources(t *testing.T) {
	validJsonData := []byte(`{
		"PolicyName": "root",
		"PolicyDocument": {
			"Version": "2012-10-17",
			"Statement": [
				{
					"Sid": "IamListAccess",
					"Effect": "Allow",
					"Action": [
						"iam:ListRoles",
						"iam:ListUsers"
					],
					"Resource": [
							"arn:aws:iam::609103258633:group/Developers",
							"arn:aws:iam::609103258633:group/Operators"
								]
				}
			]
		}
	}`)

	invalidJsonData := []byte(`{
		"PolicyName": "root",
		"PolicyDocument": {
			"Version": "2012-10-17",
			"Statement": [
				{
					"Sid": "IamListAccess",
					"Effect": "Allow",
					"Action": [
						"iam:ListRoles",
						"iam:ListUsers"
					],
					"Resource": [
							"*"
								]
				}
			]
		}
	}`)

	if !verifyJSON(validJsonData) {
		t.Error("Expected JSON file to be valid, but got invalid")
		return

	}
	if verifyJSON(invalidJsonData) {
		t.Error("Expected JSON file to be invalid, but got valid")
		return

	}

	fmt.Println("TestVerifyJSONWithMultipleResources passed")

}

func TestVerifyJSONWithInvalidFormat(t *testing.T) {
	jsonDataWithoutResource := []byte(`{
		"PolicyName": "root",
		"PolicyDocument": {
			"Version": "2012-10-17",
			"Statement": [
				{
					"Sid": "IamListAccess",
					"Effect": "Allow",
					"Action": [
						"iam:ListRoles",
						"iam:ListUsers"
					]
				}
			]
		}
	}`)

	if verifyJSON(jsonDataWithoutResource) {
		t.Error("Expected JSON file to be invalid, but got valid")
		return
	}
	fmt.Println("TestVerifyJSONWithInvalidFormat passed")

}

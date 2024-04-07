package main

import (
	"encoding/json"
	"fmt"
)

//	{
//	   "PolicyName": "root",
//	   "PolicyDocument": {
//	       "Version": "2012-10-17",
//	       "Statement": [
//	           {
//	               "Sid": "IamListAccess",
//	               "Effect": "Allow",
//	               "Action": [
//	                   "iam:ListRoles",
//	                   "iam:ListUsers"
//	               ],
//	               "Resource": "*"
//	           }
//	       ]
//	   }
//	}
type AWSPolicyDocument struct {
	PolicyName     string `json:"PolicyName"`
	PolicyDocument struct {
		Version   string `json:"Version"`
		Statement []struct {
			Sid      string      `json:"Sid"`
			Effect   string      `json:"Effect"`
			Action   interface{} `json:"Action"`
			Resource interface{} `json:"Resource"`
		} `json:"Statement"`
	} `json:"PolicyDocument"`
}

func checkResource(resource string) bool {
	if resource == "*" {
		fmt.Println("Inalid resource: ", resource)
		return false

	}
	return true
}

func checkResourceArray(resource []interface{}) bool {
	for _, r := range resource {
		if r == "*" {
			fmt.Println("Invalid resource: ", r)
			return false
		}
	}
	return true
}

func verifyJSON(exampleJSON []byte) bool {

	var policy AWSPolicyDocument
	err := json.Unmarshal(exampleJSON, &policy)
	if err != nil {
		fmt.Println("Error parsing JSON file: ", err)
		return false
	}

	//policydocument->statement->resource

	policyDocument := policy.PolicyDocument
	if policyDocument.Version == "" || policy.PolicyName == "" {
		fmt.Println("Invalid JSON format: Missing Version or PolicyName	")
		return false

	}
	statement := policyDocument.Statement
	if statement == nil {
		fmt.Println("Invalid JSON format: Missing Statement")
		return false
	}
	if len(statement) == 0 {
		fmt.Println("Invalid JSON format: Empty Statement")
		return false

	}
	for _, s := range statement {
		resource := s.Resource
		if resource == nil || resource == "" {
			fmt.Println("Invalid JSON format: Missing Resource")
			return false
		}

		switch resource.(type) {
		case string:
			if !checkResource(resource.(string)) {
				return false
			}
		case []interface{}:
			if !checkResourceArray(resource.([]interface{})) {
				return false
			}

		}
	}
	return true

}

# JSON Verifier

## Description

A Go program to verify the correctness of AWS::IAM::Role Policy JSON files. 
Method return logical false if an input JSON Resource field contains a single asterisk and true in any other case

## Requirements

- Go (version 1.16 or newer)

## Installation and Execution

1. Clone the repository:

    ```bash
    git clone https://github.com/your_repository.git
    ```

2. Run the program with:

    ```bash
    make verify FILE=your_file.json
    ```
   You have to provide the path to the JSON file as an argument.


3. Run tests with:

    ```bash
    make test
    ```
4. You can use:
    ```bash
    make all FILE=your_file.json
    ```
    to run the program with tests.
    





## JSON Structure

Sample JSON structure that is supported by this program:

```json
{
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
}

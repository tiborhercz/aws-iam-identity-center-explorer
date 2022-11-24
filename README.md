<p align="center"><img src=".github/aws-identity-center-logo.png" height="60" alt="Project Logo"></p>
<h3 align="center">AWS IAM Identity Center explorer</h3>

# About

A simple tool to get structured information about accounts and groups from the IAM Identity Center (successor to AWS Single Sign-On).

## Why does this tool exist?

The IAM Identity Center (successor to AWS Single Sign-On) user interface in the browser console can be a hard and time-consuming to navigate.
Especially if you want to the view the accounts attached to a group. 

This tool doesn't do something you cannot do with the console or CLI. It just makes it easier to view it in one single overview.

## How does it work?

To get an overview of the accounts attached to groups you have first retrieve various data and then parse the data.
This is done as follows:

1. Get all accounts in the organization
2. Get the SSO permissions sets attached to those accounts
3. List the account SSO assignments and filter the principalId with principalType `GROUP`
4. Describe the SSO groups and get the DisplayName
5. Parse the data as seen below

### Data structure options

Data structure can be chosen by using the `groups` or `accounts` command

#### Groups
Accounts attached to groups:
```json
{
  "GROUP_DISPLAY_NAME": [
    {
      "AccountName": "ACCOUNT_NAME",
      "AccountId": "ACCOUNT_ID"
    },
    {
      "AccountName": "ACCOUNT_NAME",
      "AccountId": "ACCOUNT_ID"
    }
  ]
}
```

#### Accounts
Groups attached to accounts
```json
{
  "ACCOUNT_ID": {
    "AccountName": "ACCOUNT_NAME",
    "Groups": [
      "GROUP_DISPLAY_NAME",
      "GROUP_DISPLAY_NAME"
    ]
  }
}
```

# How to run

Required parameters:
- `--identityStoreId` - The globally unique identifier for the identity store.
- `--instanceArn` - The ARN of the IAM Identity Center instance under which the operation will be executed.
- `--region` - The AWS region.

Groups data structure
```shell
go run . groups --identityStoreId IDENTITY_STORE_ID --instanceArn INSTANCE_ARN --region AWS_REGION
```

Groups data structure
```shell
go run . accounts  --identityStoreId IDENTITY_STORE_ID --instanceArn INSTANCE_ARN --region AWS_REGION
```

## AWS Authentication

The Tool uses the Go AWS SDK v2, and it detects AWS credentials set in your environment and uses them to sign requests to AWS.

The tool looks for credentials in the following environment variables:

- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_SESSION_TOKEN` (optional)

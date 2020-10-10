*Deploy Reminder*

## Dev setup

```
brew tap aws/tap
brew install aws-sam-cli
sam --version
sam build
sam local start-api
```

Should be able to access at: http://localhost:3000/hello

## Deploy 

```
sam build
sam deploy --guided
```


TODO:

Bash script
  * Take argument of reminder time

  * Gathers metadata
    * username
    * user email
    * current timestamp
    * time of code commit
    * commit msg
    * commit hash
    * pipeline name

  * Posts data to storage (primary key = curr time + argument)


Webhook to post to lambda

* parse event based on platform

* store in table

Polling service (2nd lambda)
  * cron job of per min
  * if time now is past primary key, send email notification (in future, get slack name by email)
  * remove record from storage


1. set up pipeline
2. set up buildkite webhook to send to endpoint (API Gateway)

job.finished => reminder 
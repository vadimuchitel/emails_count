# Emails Count

Small fun project to count number of unique gmail emails.

## Run instructions:
go run cmd/api/main.go 

## Input format:
```
{
   "emails": ["+dededed@gmail.com", "testgm.ail@gmail.com", "test.gmail+something@gmail.com", " testgmail@gmail.com", "testgmail2@gmail.com", "testgmail@gmail.com@gmail.com" ]
}
```

## Output format:
```
{"num_emails":2}
```

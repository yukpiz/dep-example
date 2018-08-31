package main

import (
	"fmt"
	_ "github.com/aws/aws-sdk-go/aws"
	_ "github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/aws/aws-sdk-go/service/sqs"
	_ "github.com/guregu/dynamo"
)

func main() {
	fmt.Println("ʕ◔ϖ◔ʔ")
}

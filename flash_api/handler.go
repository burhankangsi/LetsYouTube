package flash_api

import (
	"fmt"
	"context"
	"io"
	"net/http"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/S3"
	"github.com/aws/aws-sdk-go/aws/session"	
)

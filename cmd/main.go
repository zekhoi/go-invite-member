package main

import (
	"fmt"
	"os"
	"time"

	"github.com/zekhoi/go-invite-member/pkg/service"
	"github.com/zekhoi/go-invite-member/pkg/utils"
)

func main() {
	usernames, err := utils.GetUsernames("usernames.txt")

	if err != nil {
		fmt.Printf("Error getting usernames: %v\n", err)
		os.Exit(1)
	}

	for _, username := range usernames {
		_ = service.SendInvite(username)
		time.Sleep(3 * time.Second)
	}
	// fmt.Print(result)

}

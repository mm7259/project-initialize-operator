package bitbucket

import (
	"os"

	"github.com/ktrysmt/go-bitbucket"
)

func GitInit() *bitbucket.Client {
	/* Simply authenticating with bitbucket */
	client := bitbucket.NewBasicAuth(os.Getenv("BITBUCKET_USER"), os.Getenv("BITBUCKET_PASSWORD"))
	return client
}

// func CreateRepoBitBucket(teamName string) error {
// 	/* TODO - create a new repo if one does not already exist on bitbucket*/
// 	/* You will need a provlidged service account on bitbucket that has this access and use the API to create a repo */
// 	/* Store credentials as secret*/

// 	//find the base repo that's going to get forked

// }

// func CreateRepoGitHub(teamName string) error {
// 	/* TODO - create a new repo if one does not already exist on github*/
// 	/* You will need a provlidged service account on github that has this access and use the API to create a repo */
// 	/* Store credentials as secret*/
// }

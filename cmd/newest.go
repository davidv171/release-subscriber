package cmd

import (
	"fmt"

	"github.com/davidv171/release-subscriber/utils"
	"github.com/google/go-github/github"
)

//Get newest release

/* GetReleases gets the newest release from a GitHubPage*/
/* return 888 if download fails after finding the repo */
/**/
/**/
func GetNewestRelease(owner, repo, destination string) int {

	ctx, client := utils.GetClient()

	release, response, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
	fmt.Println("Printing releases...")
	if err != nil {
		fmt.Println("Error getting it ", err.Error())
		return 999
	}
	fmt.Println("Asset release "+release.GetAssetsURL(), response.StatusCode)
	if release.TarballURL != nil {
		//TODO: Change from defaulting to tarball
		url := release.TarballURL
		fmt.Println("Download URL found:  ", *url)

		filename := createReleaseFileName(release, repo)

		succ := utils.Download(url, destination+filename)

		//Download failed
		if succ != 0 {
			fmt.Println("Could not download", err)
			return 888
		}
	}
	//For unit test, author needs to be the same
	return response.StatusCode

}

func createReleaseFileName(release *github.RepositoryRelease, repo string) string {
	name := *release.Author.Login
	tag := *(release.TagName)
	extension := ".tar.gz"

	return "gpm-" + repo + "-" + name + "-" + tag + extension
}

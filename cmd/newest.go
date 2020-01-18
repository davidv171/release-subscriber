package cmd

import (
	"errors"
	"fmt"

	"github.com/davidv171/release-subscriber/utils"
	"github.com/google/go-github/github"
)

//Get newest release

/* GetReleases gets the newest release from a GitHubPage*/
/* return 888 if download fails after finding the repo */
/**/
/**/
func GetNewestRelease(owner, repo, destination string) error {

	ctx, client := utils.GetClient()

	release, response, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
	fmt.Println("Printing releases...")
	if err != nil {
		fmt.Println("Error getting it ", err.Error())
		return err
	}
	fmt.Println("Asset release "+release.GetAssetsURL(), response.StatusCode)
	if release.TarballURL != nil {
		err := tarball(release, repo, destination)
		//Download failed
		if err != nil {
			return err
		}
	}
	//For unit test, author needs to be the same
	return err

}

func tarball(release *github.RepositoryRelease, repo string, destination string) error {
	//TODO: Change from defaulting to tarball
	url := release.TarballURL
	if url == nil {
		return errors.New("The URL is null for repo: " + repo)
	}
	fmt.Println("Download URL found:  ", *url)

	filename := createReleaseFileName(release, repo)

	err := utils.Download(*url, destination+filename)
	return err
}

func createReleaseFileName(release *github.RepositoryRelease, repo string) string {
	name := *release.Author.Login
	tag := *(release.TagName)
	extension := ".tar.gz"

	return "gpm-" + repo + "-" + name + "-" + tag + extension
}

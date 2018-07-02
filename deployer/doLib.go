package deployer

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/digitalocean/godo"
	"golang.org/x/crypto/ssh"
	"golang.org/x/oauth2"
)

func compareDOConfig(initialRegion DORegionConfig, testRegion DORegionConfig) bool {
	if initialRegion.Image == testRegion.Image &&
		initialRegion.Fingerprint == testRegion.Fingerprint &&
		initialRegion.PrivateKey == testRegion.PrivateKey &&
		initialRegion.Size == testRegion.Size &&
		initialRegion.DefaultUser == initialRegion.DefaultUser {
		return true
	}
	return false
}

func genDOKeyFingerprint(publicKey string) (keyFingerprint string) {
	key, err := ioutil.ReadFile(publicKey)

	if err != nil {
		fmt.Println("Specified DO public key does not exist")
		os.Exit(1)
	}
	pubKey, _, _, _, err := ssh.ParseAuthorizedKey(key)

	if err != nil {
		fmt.Println("Specified DO public key is not formatted correctly")
		os.Exit(1)
	}

	return strings.TrimSpace(ssh.FingerprintLegacyMD5(pubKey))
}

//GetDoRegions returns the list of available regions for digital ocean
func GetDoRegions() (regions []string) {
	client := newDOClient(doToken)
	regions, err := doRegions(client)
	if err != nil {
		fmt.Println("Error retrieving region list for validation")
	}
	return
}

func doRegions(client *godo.Client) ([]string, error) {
	var slugs []string
	regions, _, err := client.Regions.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return slugs, err
	}
	for _, r := range regions {
		slugs = append(slugs, r.Slug)
	}
	return slugs, nil
}

func (t *Token) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func newDOClient(token string) *godo.Client {
	t := &Token{AccessToken: token}
	oa := oauth2.NewClient(oauth2.NoContext, t)
	return godo.NewClient(oa)
}

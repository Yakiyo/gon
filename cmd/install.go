package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Yakiyo/gon/archives"
	"github.com/Yakiyo/gon/utils"
	"github.com/Yakiyo/gon/utils/where"
	"github.com/Yakiyo/gon/versions"
	"github.com/charmbracelet/log"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install version",
	Short: "Install a version of the Go compiler",
	Long: `Install a version of the Go compiler
	
If no argument is provided, it tries to find and use the version from a go.mod file in the current directory.
"latest" can also be given as argument to install the latest available stable version.
Otherwise it expects a valid semver compliant string as argument
`,
	SilenceUsage:  true,
	SilenceErrors: true,
	Example: "gon install latest 	# install latest version\n" +
		"gon install 1.20.1 	# install specific version\n" +
		"gon install        	# use a go.mod file in current directory",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var version string
		if len(args) == 1 {
			version = args[0]
			vers, err := versions.List()
			if err != nil {
				return err
			}
			if version == "latest" {
				version = vers[0]
				log.Info("Resolving latest version", "version", version)
			}
			version = strings.TrimSuffix(strings.ToLower(version), "v")
			if !lo.Contains(vers, version) {
				return fmt.Errorf("Invalid version, %v is not a valid version for Go", version)
			}
		} else {
			if !utils.PathExists("./go.mod") {
				return fmt.Errorf("No version provided and current directory does not contain a go.mod file. Provide a version explicitly.")
			}
			b, err := os.ReadFile("./go.mod")
			if err != nil {
				return fmt.Errorf("Unable to read local go.mod file, received error %s", err)
			}
			str := string(b)
			// find line that contains `go {{ version }}`
			vline, ok := lo.Find[string](strings.Split(str, "\n"), func(item string) bool {
				return strings.HasPrefix(item, "go")
			})
			if !ok {
				return fmt.Errorf("Invalid go.mod file, file does not specify go version")
			}
			version = versions.SafeVStr(strings.TrimSpace(strings.ReplaceAll(vline, "go", "")))
			if version == "" {
				return fmt.Errorf("Could not parse version from go.mod file. Line contains %s", vline)
			}
			log.Info("Resolving version from go.mod", "version", version)
		}
		intls := where.Installations()
		versionDir := filepath.Join(intls, version)
		if utils.PathExists(versionDir) {
			return fmt.Errorf("Version %s is already installed. Uninstall it first and then install it again if needed", version)
		}
		err := utils.EnsureDir(intls)
		if err != nil {
			return fmt.Errorf("Unable to create installations dir due to error %v", err)
		}
		url := versions.VersionArchiveUrl(version, viper.GetString("arch"))

		file, err := os.CreateTemp("", "archive")

		if err != nil {
			return err
		}
		log.Info("Downloading archive to temp directory", "url", url, "tmp", file.Name())
		defer os.Remove(file.Name())
		err = downloadToFile(url, file)
		if err != nil {
			return err
		}
		if runtime.GOOS == "windows" {
			log.Info("Extracting zip file")
			err = archives.Unzip(file, intls)
		} else {
			log.Info("Extracting tar file")
			err = archives.Untar(file.Name(), intls)
		}
		if err != nil {
			return fmt.Errorf("Failed to extract files from archive due to error %s", err)
		}
		err = os.Rename(filepath.Join(intls, "go"), versionDir)
		if err != nil {
			return fmt.Errorf("Failed to rename extracted folder due to error %v", err)
		}
		fmt.Printf("Successfully installed Go version %v\n", version)
		return nil
	},
}

// download archive from `url` to file `file`
func downloadToFile(url string, file *os.File) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			return fmt.Errorf("No archive found at url %v. Please ensure you are using correct architecture and version", url)
		}
		return fmt.Errorf("Failed to download archive, received status code %v %v", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	_, err = io.Copy(file, resp.Body)
	return err
}

func init() {
	rootCmd.AddCommand(installCmd)
}

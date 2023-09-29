# gon

Fast and simple version manager for the [Go](https://go.dev) compiler, written in Go, inspired from [nvm](https://github.com/nvm-sh/nvm).

## Installation

### Scoop (windows)
```
 $ scoop install https://github.com/Yakiyo/gon/raw/main/share/scoop/gon.json
 ```

### Brew
Theres a brew formula at [share/brew](./share/brew/gon.rb)

*More installation methods TBA*

## Setup
For using the Go compiler installed by gon, you need to add the downloaded path to your `PATH` environment variable. This can be manually added to the sys path/user path, or be automated in your shell's init script.

For example, in case of bash, add the following line to your `~/.bashrc`
```sh
export PATH="$(gon path bin)":$PATH
```
This works the same for zsh too.

For powershell, the following can be added to `profile.ps1`
```powershell
$Dir = gon path bin
$User = [System.EnvironmentVariableTarget]::User
$Path = [System.Environment]::GetEnvironmentVariable('Path', $User)
if (!(";${Path};".ToLower() -like "*;${Dir};*".ToLower())) {
  [System.Environment]::SetEnvironmentVariable('Path', "${Path};${InstallDir}", $User)
  $Env:Path += ";${Dir}"
}
```

## Usage
For installing a specific version, use the `gon install` command.
```sh
$ gon install latest # install latest Go version

$ gon install v1.20 # install specific version, the `v` suffix is optional

$ gon install # install the version mentioned in `./go.mod`
```

For uninstalling, the specific version must be mentioned
```sh
$ gon uninstall 1.20
```
Aliases can be added and removed via the `gon alias` and `gon unalias` commands.

For using an installed version, run 
```sh
$ gon use v1.20 # use specific version

$ gon use # use version mentioned in `./go.mod`
```

The name gon comes from the character [Gon Freecss](https://anilist.co/character/30/Gon-Freecss) from [Hunter X Hunter](https://anilist.co/anime/11061/Hunter-x-Hunter-2011/)
## Author

**gon** © [Yakiyo](https://github.com/Yakiyo). Authored and maintained by Yakiyo.

Released under [MIT](https://opensource.org/licenses/MIT) License

If you like this project, consider leaving a star ⭐ and sharing it with your friends and colleagues.
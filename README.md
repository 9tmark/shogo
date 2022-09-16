# shogo
[![MPLv2 License](https://img.shields.io/badge/license-MPLv2-blue.svg?style=flat-square)](https://www.mozilla.org/MPL/2.0/)

One-step url generator for [Short.io](https://short.io). Get a short url to share files or long links, straight from the command line into your clipboard (and stdout). Written in Go.

## Install a static build
> NOT AVAILABLE AT THE MOMENT

## Install from source
### Prerequisites
- [Go](https://golang.org/dl/) version 1.19 or higher
- A [Short.io](https://app.short.io/public/register) account

### Clone, prepare individual information and build
```sh
$ git clone https://github.com/9tmark/shogo.git
$ cd shogo
```
Have your [API key](https://short.io/features/api) ready:
```sh
$ echo "<your api key>" > key.txt
```
It's recommended to delete the key file after building.
Prepare your domain if you have one, otherwise use the default one (`short.io`). Then start installation:
```sh
$ echo "<your domain>" > domain.txt
$ make install
```
(The domain has to be setup for usage with Short.io beforehand)

## Usage
Use `shogo` or the shorter alias `shg` with a target url as passed argument. This is it. If you are in need of a more detailed output, activate verbose mode my appending `--verbose` to the command.

```sh
$ shogo --verbose "http://github.com"
```
```sh
$ shg "http://github.com"
```

### Provide your API key
If not already built-in or if you would like to exchange the built-in key dynamically, there are two options to provide a key.

1. You can put a **key.txt** file in the current directory containing the key.
2. You can set a key by storing it to the environmental variable `SHORT_IO_KEY`, temporarily or in your local profile.

The priority order of evaluation is:
1. **key.txt** in current directory
2. `SHORT_IO_KEY` in env
3. built-in key (only for self-built installations)

### Provide your domain
This works the same way you set up your API key.

1. You can put a **domain.txt** file in the current directory containing the domain of choice.
2. You can set a domain by storing it to the environmental variable `SHG_DOMAIN`, temporarily or in your local profile.

## Final notice
Please check [LICENSE](LICENSE) before contributing or using *shogo*. Please keep in mind: Use at own risk. This tool is not affiliated with Short.io. It is a personal project. Nice to know: If you provide Dropbox links, shogo will automatically append the `dl=0` parameter. Also, this tool is not affiliated with Dropbox.

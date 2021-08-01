# crtsh-cli - Subdomain Enumeration

`crtsh-cli` is a very simple Go program which allows you to query [crt.sh](https://crt.sh/) from the command line. Given a URL (e.g., *github.com*) it will provide a list (line by line) of associated (sub)domains retrieved via crt.sh's *Certificate Search*.

I primarily use this tool during CTF/Pentest reconnaissance.

<img alt="crtsh-cli Screenshot" src="https://user-images.githubusercontent.com/16179317/127785078-bce36f89-fc95-4322-a2d0-ecd146d3d577.png" width="500" />

## Running & Building

First, run `git clone https://github.com/IngoKl/crtsh-cli` into a directory of your choice.
Afterward, you can either run or compile the program:

`go run . github.com` or `go build`.

Of course, after building the tool, you can add it to your path for easy access.

Alternatively, you can also create a symlink in order to make the tool available to you more conveniently. 
This solution isn't very clean but works well enough.

On *Kali*, for example, you could do the following:

```bash
cd /home/kali && git clone https://github.com/IngoKl/crtsh-cli
cd crtsh-cli
go build
sudo ln -s /home/kali/crtsh-cli/crtsh-cli /usr/bin/crtsh-cli
```

## Usage Example

`crtsh-cli github.com`

```bash
Checking: github.com

r2.shared.global.fastly.net
f.cloud.github.com
github.com
*.github.com
skyline.github.com
support.enterprise.github.com
api.stars.github.com
dns-vetting1i.map.fastly.net
mac-installer.github.com
jira.github.com

[...]
```

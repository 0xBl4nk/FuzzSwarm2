<div align="center">

<img src="https://i.imgur.com/K1FYeZK.png" width=40%>

# Fuzz Swarm 2.0

![Go version][go_version_img]
[![Wiki][repo_wiki_img]][repo_wiki_url]

**FuzzSwarm** is a powerful **multi-threaded fuzzing tool** designed for **brute-forcing HTTP endpoints** and identifying vulnerabilities in web applications and **APIs**. It supports **GET and POST requests**, custom headers, and can load headers from a file. **Users can perform precise fuzzing with numeric ranges, wordlists, and filter responses by size**. FuzzSwarm also offers rate limit and timeout controls, as well as **proxy and SSL support** for secure testing environments, making it ideal for pentesters and security professionals seeking to uncover web security flaws.

</div>

## ⚡️ Quick start

First, [download][go_download_url] and install **Go**. Version `1.23` or
higher is required.

Clone FuzzSwarm Repository:

```bash
git clone https://github.com/0xBl4nk/FuzzSwarm2
```

Build the project or use the precompiled binary:
```bash
# Build project:
cd FuzzSwarm2
make build

# Usage:
./FuzzSwarm --help
```
## 📖 Project Wiki

The best way to better explore all the features of the **Fuzz Swarm 2**
is to read the project [Wiki][repo_wiki_url].

Yes, the most frequently asked questions (FAQ) are also
[here][repo_wiki_faq_url].

## ⚙️ Commands & Options

### `POST Example`

> Cheat sheet for some uses of Fuzz Swarm 2

```bash
./FuzzSwarm -X POST -u http://127.0.0.1/api/2fa \
 -R 1-10000,1 -d '{"number": FUZZ}' \
  -H "Content-Type: application/json" \
   -f 34 -v
```

| Option | Description                                              | Type   | Default | Required? |
| ------ | -------------------------------------------------------- | ------ | ------- | --------- |
| `-X `   | Select HTTP method | `string` | `GET` | No        |
| `-R` | Use number range instead of word list | `string` || yes, if you don't use word list |
| `-d` | Set POST data | `string` || No
| `-H` | Set custom headers |  `string` || No
| `-f` | Skip answer with answer length | `int` || No
| `-v` | Show response body | `bool` | `False` | No

![ImageUsage](https://i.imgur.com/guvTo1Y.png)

- 📖 Docs: https://0xbl4nk.github.io/FuzzSwarm2-wiki/

### `Scripts Example`

You can use scripts in FuzzSwarm to automate specific attack types, such as SSTI fuzzing, with predefined payloads for more targeted vulnerability testing.

```bash
./FuzzSwam --script ssti -u 'http://127.0.0.1/vulnerable?input=FUZZ' -v
```

| Option | Description                                                                                            | Type   | Default | Required? |
| ------ | ------------------------------------------------------------------------------------------------------ | ------ | ------- | --------- |
| `--script`   | Select the script to use | `string` | | No |

![cgapp_deploy](https://i.imgur.com/1rjekSu.png)

- 📖 Docs: https://0xbl4nk.github.io/FuzzSwarm2-wiki/

---

<!-- Go -->

[go_download_url]: https://golang.org/dl/
[go_version_img]: https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go
[repo_logo_url]: https://github.com/0xBl4nk/FuzzSwarm2
[repo_wiki_url]: https://0xbl4nk.github.io/FuzzSwarm2-wiki/
[repo_wiki_img]: https://img.shields.io/badge/docs-wiki_page-blue?style=for-the-badge&logo=none
[repo_wiki_faq_url]: https://0xbl4nk.github.io/FuzzSwarm2-wiki/faq/


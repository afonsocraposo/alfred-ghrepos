# Alfred GHRepos Workflow 🚀

## Description 📝

This Alfred workflow allows you to quickly search and access your GitHub repositories using `fzf` 🕵️‍♂️, a command-line fuzzy finder. It streamlines the process of finding and navigating to your repositories, enhancing productivity and workflow efficiency 🚀.

## Features 🌟

- **Fast Repository Search**: Utilizes `fzf` for efficient and rapid searching through your GitHub repositories 🏃‍♂️💨.
- **Easy Repository Access**: Directly navigate to the selected repository in your web browser 🌐.
- **Go**: Written in Golang - what more can I say 🐭.

## Setup 🛠️

To use this workflow, you need to set up a personal access token from GitHub with "repo" permissions. This token ensures secure and personalized access to your repositories.

### Steps to Setup:

1. **Generate a GitHub Personal Access Token** 🔑:
   - Go to [GitHub Tokens Settings](https://github.com/settings/tokens).
   - Generate a new token with "repo" permissions.

2. **Configure Alfred Workflow** 🧰:
   - Open Alfred Preferences and navigate to the GHRepos workflow.
   - Click on "Configure Workflow...".
   - Set the value of `GH_TOKEN` to your newly generated GitHub Personal Access Token.

## Usage 🖥️

After setting up, you can activate the workflow within Alfred:

1. Invoke Alfred 🔍.
2. Type `gh` followed by your search query to find a repository.
3. Use `fzf` search capabilities to quickly filter and select the desired repository 👓.
4. Press `Enter` to open the selected repository in your web browser 🌐.

To refresh the cache:
1. Invoke Alfred 🔍.
2. Type `gh` and press `Enter`.

## Support 🆘

For support, suggestions, or contributions, please visit the [workflow's GitHub repository](https://github.com/afonsocraposo/alfred-ghrepos).

## License 📄

[MIT License](LICENSE)

---

This workflow is not affiliated with, directly supported, or endorsed by GitHub or Alfred.

# Actions Case Police

🚨 Use the correct case, even on GitHub issues. Directly inspired by the original [case-police](https://github.com/antfu/case-police).

## Installation

> To use `actions-case-police`, you'll need to prepare a GitHub access token. Please refer to this [article](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) on how to generate an access token.

You can integrate `actions-case-police` to your existing GitHub actions workflow by using `Namchee/actions-case-police@<version>` in one of your jobs using `yaml`.

Below is the example of using `actions-case-police` job in your action workflow.

```
on:
  pull_request:
  issues:

jobs:
  case-police:
    runs-on: ubuntu-latest
    steps:
      - name: Use the correct case
        uses: Namchee/actions-case-police@v(version)
        with:
          access_token: <YOUR_GITHUB_ACCESS_TOKEN_HERE>
          fix         : true
```

Please refer to [GitHub workflow syntax](https://docs.github.com/en/free-pro-team@latest/actions/reference/workflow-syntax-for-github-actions#about-yaml-syntax-for-workflows) for more advanced usage.

> 💡 You can use a special syntax `${{ secrets.GITHUB_TOKEN }}` as your access token and the `github-actions` bot will run the job on behalf of you.

## Dictionary

Every words that can be case-policed is stored in dictionaries. Please refer to the [dictionary folder](./dict/) for all default words.

> 💡 You can also provide your own dictionary to be used with the default dictionary with the `dictionary` input.

## Inputs

You can customize this actions with these following options (fill it on `with` section):

| **Name**       | **Required?** | **Default Value** | **Description** |
| -------------- | ------------- | ----------------- | --------------- |
| `access_token` | `true`        | `-` | [GitHub access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token) to interact with the GitHub API. It is recommended to store this token with [GitHub Secrets](https://docs.github.com/en/free-pro-team@latest/actions/reference/encrypted-secrets).
| `fix`          | `false`        | `true`            | Determines if possible case fixes should be applied automatically. If set to `false`, `actions-case-police` will only log possible fixes in the action log.
| `preset`       | `false`        | ``["abbreviates", "brands", "general", "products", "softwares"]`` | Dictionary names to be used when validating word cases. By default, it will use all default dictionary. Comma-separated
| `exclude`      | `false`        | `""` | Words to be whitelisted on case police. Comma-separated
| `dictionary`   | `false`        | `{}` | Stringified JSON map that represents custom entiries for dictionary that will be used on case police. 

## Special Thanks

- [Anthony Fu](https://github.com/antfu), for the original [case-police](https://github.com/antfu/case-police) and for granting me the permission for this project.

## License

This project is licensed under the [MIT license](./LICENSE)
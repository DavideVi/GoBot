# GoBot

## Activity

POST to `/message` to send message to bot. Message format is JSON.

**Example:**
```
{
  "message": "Enable the VPN"
  "user": "Davide"
}
```

Bot receives message and looks into `handlers.json`.

The handlers file is used to determine which script to execute based on a regular
expression.

**Example:**
```
{
  "handlers": [
    {
      "regex": "/(?:enable the)(.*)|(?:Enable the)(.*)|(?:enable)(.*)/g",
      "handler": "enable.py",
      "executor": "python"
    },
    {
      "regex": "/(?:enable the)(.*)|(?:Enable the)(.*)|(?:enable)(.*)/g",
      "handler": "enable.sh",
      "executor": "bash"
    },
    {
      "regex": "/(?:enable the)(.*)|(?:Enable the)(.*)|(?:enable)(.*)/g",
      "handler": "enable.ps1",
      "executor": "powershell"
    }
  ]
}
```

All regexes are checked and if they match then the script is executed. Capture
groups will be passed as parameters. Stdout will be returned to the user.

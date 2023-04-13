
# Scripts

These scripts were created to run on Linux, and may not work on Windows.
Can be run in the terminal or in "Tasks" from vscode.


## terminal


    - open a terminal on script/linux folder
    - execute normally as ./myscript.sh

## vscode Tasks

- Add below configuration in .vscode/tasks.json file
- In vscode, press F1, and choose "Tasks:Run Task"


```
{
    // See https://go.microsoft.com/fwlink/?LinkId=733558
    // for the documentation about the tasks.json format
    // https://code.visualstudio.com/docs/editor/tasks

    "version": "2.0.0",
    "options": {
        "cwd": "${workspaceFolder}/scripts/linux"
      },
    "tasks": [
        
        {
            "label": "go-build-cmd",
            "type": "shell",
            "command": "./go-build-cmd.sh",            
        },
        {
            "label": "go-test",
            "type": "shell",
            "command": "./go-test.sh",            
        },
        {
            "label": "go-lint",
            "type": "shell",
            "command": "./go-lint.sh",            
        },
        {
            "label": "backup-this-project",
            "type": "shell",
            "command": "./backup-this-project.sh",           
        },
        {
            "label": "git-tools",
            "type": "shell",
            "command": "./git-tools.sh",
        }

    ]
}

```
# Global flags
complete -c notes -s h -l help -d "Show context-sensitive help."
complete -c notes -s A -l color-always -d "Enable color output always"
complete -c notes -l no-color -d "Disable color output"
complete -c notes -l version -d "Show application version."

# Subcommands
complete -c notes -n '__fish_use_subcommand' -xa 'help' -d "Show help."
complete -c notes -n '__fish_use_subcommand' -xa 'add' -d "Create/Edit a note with file name"
complete -c notes -n '__fish_use_subcommand' -xa 'list' -d "List notes By default"
complete -c notes -n '__fish_use_subcommand' -xa 'delete' -d "delete a note by file name"
complete -c notes -n '__fish_use_subcommand' -xa 'get' -d "Print the content of a note by file name"
complete -c notes -n '__fish_use_subcommand' -xa 'export' -d "Export notes to a zip file"
complete -c notes -n '__fish_use_subcommand' -xa 'import' -d "Import notes from a zip file"

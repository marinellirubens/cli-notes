# Global flags
complete -c notes -s h -l help -d "Show context-sensitive help."
complete -c notes -s v -l version -d "Show application version."

# Subcommands
complete -c notes -n '__fish_use_subcommand' -xa 'help' -d "Show help."
complete -c notes -n '__fish_use_subcommand' -xa 'add' -d "Create/Edit a note with file name"
complete -c notes -n '__fish_use_subcommand' -xa 'list' -d "List notes By default"
complete -c notes -n '__fish_use_subcommand' -xa 'delete' -d "delete a note by file name"
complete -c notes -n '__fish_use_subcommand' -xa 'get' -d "Print the content of a note by file name"
complete -c notes -n '__fish_use_subcommand' -xa 'export' -d "Export notes to a zip file"
complete -c notes -n '__fish_use_subcommand' -xa 'import' -d "Import notes from a zip file"

complete -c notes -n '__fish_seen_subcommand_from export' -s o -l output -d "Output file name, (default: .notes.zip)"
complete -c notes -n '__fish_seen_subcommand_from import' -s i -l input -d "Input file name, (default: .notes.zip)"


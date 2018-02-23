---
Title: File operations
Id: 3333
---

TODO:
* get file size
* get information about a file
* check if a file exists
* rename a file
* remove a file
* copy a file

- err := os.Remove(*name*) // Deletes a file. A non-nil error is returned if the file could not be deleted.

- err := os.Rename(*oldName*, *newName*) // Renames or moves a file (can be across directories). A non-nil error is returned if the file could not be moved.

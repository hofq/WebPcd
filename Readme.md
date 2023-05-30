# WebP Converter Daemon
The WebP Converter Daemon is an application that Checks a folder  for WebP Files and Converts them into more usable Image Formats.
The Application is needed because WebP is currently not Supported Natively in many Windows Applications.
The Application Converts all WebP Files in your Specified Folder to the Next best Image Format. (Currently: just png)

## Options
### RMOriginal
Default: `RMOriginal: false` Specifies if the Original File should be kept
### exclude_format (not yet implemented)
Default: `Not Set`
Example: `exclude_format: png,gif,jpg` Can be used to specify unwanted Image Formats
### use_format (not yet implemented)
Can be used to specify the wanted Image Formats
Default: `use_format: png,gif,jpg`
### OutputPath
Specifies the Path where the Converted Images Should be saved
Default: `OutputPath: "<Your Home Dir>/Downloads/"` Can be specified as a var
### InputPath
Specifies the Path where the System should look for new files
Default: `InputPath: "<Your Home Dir>/Downloads/"` Can be specified as a var

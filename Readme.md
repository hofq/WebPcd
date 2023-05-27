# WebP Converter Daemon
The WebP Converter Daemon is an application that Checks a folder  for WebP Files and Converts them into more usable Image Formats.
The Application is needed because WebP is currently not Supported Natively in many Windows Applications.
The Application Converts all WebP Files in your Specified Folder to the Next best Image Format.

## Options
### remove_original
Default: `remove_original: false` Specifies if the Original File should be kept
### exclude_format
Default: `Not Set`
Example: `exclude_format: png,gif,jpg` Can be used to specify unwanted Image Formats
### use_format
Default: `use_format: png,gif,jpg`
Can be used to specify the wanted Image Formats
Default: `locations: "C:\\Users\\%userprofile%\\Downloads\\"` Can be specified as a var or a list
list Example: 
```yml
locations:
  - "C:\\Users\\%userprofile%\\Downloads\\"
  - "C:\\Images\\"
  - "C:\\Users\\Documents\\Assets\\"
```
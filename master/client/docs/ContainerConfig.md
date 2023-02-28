# ContainerConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** |  Hostname | [optional] [default to null]
**Domainname** | **string** |  Domainname | [optional] [default to null]
**User** | **string** |  User that will run the command(s) inside the container, also support user:group | [optional] [default to null]
**AttachStdin** | **bool** |  Attach the standard input, makes possible user interaction | [optional] [default to null]
**AttachStdout** | **bool** |  Attach the standard output | [optional] [default to null]
**AttachStderr** | **bool** |  Attach the standard error | [optional] [default to null]
**ExposedPorts** | **[]string** |  List of exposed ports | [optional] [default to null]
**Tty** | **bool** |  Attach standard streams to a tty, including stdin if it is not closed. | [optional] [default to null]
**OpenStdin** | **bool** |  Open stdin | [optional] [default to null]
**StdinOnce** | **bool** |  If true, close stdin after the 1 attached client disconnects. | [optional] [default to null]
**Env** | **[]string** |  List of environment variable to set in the container | [optional] [default to null]
**Cmd** | **[]string** |  Command to run when starting the container | [optional] [default to null]
**Healthcheck** | [***HealthConfig**](HealthConfig.md) |  Healthcheck describes how to check the container is healthy | [optional] [default to null]
**ArgsEscaped** | **bool** |  True if command is already escaped (meaning treat as a command line) (Windows specific). | [optional] [default to null]
**Image** | **string** |  Name of the image as it was passed by the operator (e.g. could be symbolic) | [default to null]
**Volumes** | **[]string** |  List of volumes (mounts) used for the container | [optional] [default to null]
**WorkingDir** | **string** |  Current directory (PWD) in the command will be launched | [optional] [default to null]
**Entrypoint** | **[]string** |  Entrypoint to run when starting the container | [optional] [default to null]
**NetworkDisabled** | **bool** |  Is network disabled | [optional] [default to null]
**MacAddress** | **string** |  Mac Address of the container | [optional] [default to null]
**OnBuild** | **[]string** |  ONBUILD metadata that were defined on the image Dockerfile | [optional] [default to null]
**Labels** | **interface{}** |  List of labels set to this container | [optional] [default to null]
**StopSignal** | **string** |  Signal to stop a container | [optional] [default to null]
**StopTimeout** | **int32** |  Timeout (in seconds) to stop a container | [optional] [default to null]
**Shell** | **[]string** |  Shell for shell-form of RUN, CMD, ENTRYPOINT | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



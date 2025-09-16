// This file contains the version & build time information for the application.
package main

// Linker overwrites the strings below at compile time.
var (
	setAppName              = "tbd" // The name of the application.
	setBuildType            = "tbd" // The type of build (e.g., Local, Jenkins).
	setRemoteRepo           = "tbd" // The remote repository URL.
	setCurrentBranch        = "tbd" // The current branch of the repository.
	setBuilddate            = "tbd" // The date of the build.
	setBuildplatform        = "tbd" // The platform on which the build is performed.
	setBuildcompilerversion = "tbd" // The version of the compiler used for the build.
	setBuildTargetOS        = "tbd" // The target operating system for the build.
	setBuildTargetArch      = "tbd" // The target architecture for the build.
	setCGOEnabled           = "tbd" // Indicates whether CGO is enabled ("0" or "1").
	setCurrentCommitHash    = "tbd" // The Git hash of the current commit.
	setCurrentFullTag       = "tbd" // The full tag of the current commit.
	setVersion              = "tbd" // The version of the application.
	setCurrentTag           = "tbd" // The current tag of the repository.

	// buildDetails represents the build details as an object of the application as passed in when compiled.
	//buildDetails buildtime.Details = buildtime.Details{
	//	AppName:              setAppName,
	//	BuildType:            setBuildType,
	//	RemoteRepo:           setRemoteRepo,
	//	CurrentBranch:        setCurrentBranch,
	//	BuildDate:            setBuilddate,
	//	BuildPlatform:        setBuildplatform,
	//	BuildCompilerVersion: setBuildcompilerversion,
	//	BuildTargetOS:        setBuildTargetOS,
	//	BuildTargetArch:      setBuildTargetArch,
	//	CGOEnabled:           setCGOEnabled,
	//	CurrentCommitHash:    setCurrentCommitHash,
	//	CurrentFullTag:       setCurrentFullTag,
	//	Version:              setVersion,
	//	CurrentTag:           setCurrentTag,
	//}
)

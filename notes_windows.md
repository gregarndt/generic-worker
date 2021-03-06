User Data
=========

In order to set up a new AWS Provisioner Worker Type running on Windows, follow these steps:

1. Launch a Windows instance in AWS with the following UserData:

   ```powershell
   <powershell>
   
   $client = New-Object system.net.WebClient
   
   function Expand-ZIPFile($file, $destination, $url)
   {
       $client.DownloadFile($url, $file)
       $shell = new-object -com shell.application
       $zip = $shell.NameSpace($file)
       foreach($item in $zip.items())
       {
           $shell.Namespace($destination).copyhere($item)
       }
   }
   
   # install go
   md "C:\gopath"
   Expand-ZIPFile -File "C:\go1.4.2.windows-amd64.zip" -Destination "C:\" -Url "https://storage.googleapis.com/golang/go1.4.2.windows-amd64.zip"
   
   # install PSTools
   md "C:\PSTools"
   Expand-ZIPFile -File "C:\PSTools\PSTools.zip" -Destination "C:\PSTools" -Url "https://download.sysinternals.com/files/PSTools.zip"
   
   # install nssm
   Expand-ZIPFile -File "C:\nssm-2.24.zip" -Destination "C:\" -Url "http://www.nssm.cc/release/nssm-2.24.zip"
   
   # install git
   $client.DownloadFile("https://github.com/msysgit/msysgit/releases/download/Git-1.9.5-preview20150319/Git-1.9.5-preview20150319.exe", "C:\git-1.9.5-installer.exe")
   $p = Start-Process "C:\git-1.9.5-installer.exe" -ArgumentList "/SILENT" -Wait -PassThru
   $p.HasExited
   
   # set env vars
   [Environment]::SetEnvironmentVariable("GOROOT", "C:\go", "Machine")
   [System.Environment]::SetEnvironmentVariable("PATH", $Env:Path + ";C:\go\bin;C:\Program Files (x86)\Git\cmd", "Machine")
   [Environment]::SetEnvironmentVariable("GOPATH", "C:\gopath", "User")
   
   # download generic-worker
   md C:\generic-worker
   $client.DownloadFile("https://github.com/taskcluster/generic-worker/releases/download/v1.0.11/generic-worker-windows-amd64.exe", "C:\generic-worker\generic-worker.exe")
   
   # enable DEBUG logs for generic-worker install
   $env:DEBUG = "*"
   
   # install generic-worker
   $p = Start-Process C:\generic-worker\generic-worker.exe -ArgumentList "install --config C:\\generic-worker\\generic-worker.config" -wait -NoNewWindow -PassThru -RedirectStandardOutput C:\generic-worker\install.log -RedirectStandardError C:\generic-worker\install.err
   $p.HasExited
   </powershell>
   ```
2. Connect to the instance
3. Install any additional toolchains required for your tasks
4. Snapshot the instance, creating an AMI
5. Create a Worker Type in the AWS Provisioner, referencing the AMI you created

See https://www.youtube.com/watch?t=800&v=B1MAyJpUya8 for a complete walkthrough.

Windows Firefox Builds
======================

Here is the powershell script that can be supplied as UserData when launching
an AWS instance, to install the generic worker plus the toolchains required for
building Firefox.

[Firefox Powershell UserData](firefox.userdata)

Script to create AMI
====================

See the [create.sh](create.sh) script.

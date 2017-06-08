vRealize Automation provides UI where developers,users can request IT services and manage IT resources. vRA provider is created newly. 
It utilizes go library. It uses rest apis of vRA and send http request.   

Resources:

Execute Blueprint:
 It executes any blueprint present in VRA.
 Blueprint name is given in configuaration file.
 To execute that blueprint json data should pass from user in .json file  

The main.tf file contains the microservices of how to call the providers and resources. We need to specify required details for resource creation in this file.
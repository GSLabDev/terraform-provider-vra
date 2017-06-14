provider "vra" {
        host_url = "<vra_host_url/vra_host_ip>"
        tenant = "<tenant_name>"
        user_name = "<vra_user_name>"
        user_password = "<vra_user_password>"
}


resource "vra_execute_blueprint" "ExecuteBlueprint" {
        
       blueprint_name = "Create simple virtual machine"
       file_name = "data.json"
}

provider "vra" {
        host_url = "192.168.35.95"
        tenant = "team"
        user_name = "demo"
        user_password = "gsLab!23"
}


resource "vra_execute_blueprint" "ExecuteBlueprint" {
        
       blueprint_name = "Create simple virtual machine"
       file_name = "da.json"
}



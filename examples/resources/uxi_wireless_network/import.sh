# Import a wireless network using its ID
terraform import uxi_wireless_network.my_wireless_network <my_wireless_network_id>

# Import a wireless network using its ID with an import block
import {
    to = uxi_wireless_network.my_wireless_network
    id = "<my_wireless_network_id>"
}
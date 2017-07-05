Name                           : hoge.example1.com
TFTP Boot Files                : {}
Comment                        :
Enable gPXE?                   : False
Fetchable Files                : {}
Gateway                        : 192.168.0.254
Hostname                       : hoge.example.com
Image                          :
IPv6 Autoconfiguration         : False
IPv6 Default Device            :
Kernel Options                 : {'serial': '', 'console': ['tty0', 'ttyS0,115200n8'], 'edd': 'off'}
Kernel Options (Post Install)  : {}
Kickstart                      : /var/lib/cobbler/kickstarts/test.ks
Kickstart Metadata             :
LDAP Enabled                   : False
LDAP Management Type           :
Management Classes             : []
Management Parameters          : <<inherit>>
Monit Enabled                  : False
Name Servers                   : ['8.8.4.4', '8.8.8.8']
Name Servers Search Path       : ['example.com']
Netboot Enabled                : False
Owners                         : ['test_user']
Power Management Address       : 192.168.1.1
Power Management ID            :
Power Management Password      :
Power Management Type          :
Power Management Username      :
Profile                        :
Internal proxy                 : <<inherit>>
Red Hat Management Key         : <<inherit>>
Red Hat Management Server      : <<inherit>>
Repos Enabled                  : False
Server Override                : <<inherit>>
Status                         : production
Template Files                 : {}
Virt Auto Boot                 :
Virt CPUs                      :
Virt Disk Driver Type          :
Virt File Size(GB)             : 0
Virt Path                      : <<inherit>>
Virt PXE Boot                  : 0
Virt RAM (MB)                  :
Virt Type                      :
Interface =====                : bond0.1
Bonding Opts                   :
Bridge Opts                    :
CNAMES                         : []
InfiniBand Connected Mode      : False
DHCP Tag                       :
DNS Name                       :
Per-Interface Gateway          :
Master Interface               :
Interface Type                 :
IP Address                     : 1.1.1.1
IPv6 Address                   :
IPv6 Default Gateway           :
IPv6 MTU                       :
IPv6 Prefix                    :
IPv6 Secondaries               : []
IPv6 Static Routes             : []
MAC Address                    :
Management Interface           : False
MTU                            :
Subnet Mask                    : 255.255.255.0
Static                         : True
Static Routes                  : []
Virt Bridge                    :
Interface =====                : bond0
Bonding Opts                   : mode=1 miimon=100 primary=eth0 updelay=5000
Bridge Opts                    :
CNAMES                         : []
InfiniBand Connected Mode      : False
DHCP Tag                       :
DNS Name                       :
Per-Interface Gateway          :
Master Interface               :
Interface Type                 : bond
IP Address                     : 192.168.0.1
IPv6 Address                   :
IPv6 Default Gateway           :
IPv6 MTU                       :
IPv6 Prefix                    :
IPv6 Secondaries               : []
IPv6 Static Routes             : []
MAC Address                    :
Management Interface           : False
MTU                            :
Subnet Mask                    : 255.255.254.0
Static                         : True
Static Routes                  : ['192.168.0.0/16:172.16.0.1']
Virt Bridge                    :
Interface =====                : eth2
Bonding Opts                   :
Bridge Opts                    :
CNAMES                         : []
InfiniBand Connected Mode      : False
DHCP Tag                       :
DNS Name                       :
Per-Interface Gateway          :
Master Interface               : bond0
Interface Type                 : bond_slave
IP Address                     :
IPv6 Address                   :
IPv6 Default Gateway           :
IPv6 MTU                       :
IPv6 Prefix                    :
IPv6 Secondaries               : []
IPv6 Static Routes             : []
MAC Address                    : AA:AA:AA:AA:AA:22
Management Interface           : False
MTU                            :
Subnet Mask                    :
Static                         : False
Static Routes                  : []
Virt Bridge                    : 
Interface =====                : eth0
Bonding Opts                   :
Bridge Opts                    :
CNAMES                         : []
InfiniBand Connected Mode      : False
DHCP Tag                       :
DNS Name                       :
Per-Interface Gateway          :
Master Interface               : bond0
Interface Type                 : bond_slave
IP Address                     :
IPv6 Address                   :
IPv6 Default Gateway           :
IPv6 MTU                       :
IPv6 Prefix                    :
IPv6 Secondaries               : []
IPv6 Static Routes             : []
MAC Address                    : AA:AA:AA:AA:AA:11
Management Interface           : False
MTU                            :
Subnet Mask                    :
Static                         : False
Static Routes                  : []
Virt Bridge                    :

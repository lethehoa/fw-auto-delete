#!/bin/bash
cd /root/tan/proxy-tan && ansible-playbook vmware-generate-inventory.yml -l firewall1.vietnix.vn  &&  ansible-playbook vmw-deploy-autoproxy.yml -l firewall1.vietnix.vn  -v
cd /root/tan/proxy-tan && ansible-playbook vmware-generate-inventory.yml -l firewall2.vietnix.vn  &&  ansible-playbook vmw-deploy-autoproxy.yml -l firewall2.vietnix.vn  -v
cd /root/tan/proxy-tan && ansible-playbook vmware-generate-inventory.yml -l firewall3.vietnix.vn  &&  ansible-playbook vmw-deploy-autoproxy.yml -l firewall3.vietnix.vn  -v
cd /root/tan/proxy-tan && ansible-playbook vmware-generate-inventory.yml -l firewall4.vietnix.vn  &&  ansible-playbook vmw-deploy-autoproxy.yml -l firewall4.vietnix.vn  -v
#cd /root/tan/proxy-tan && ansible-playbook vmware-generate-inventory.yml -l firewall5.vietnix.vn  &&  ansible-playbook vmw-deploy-autoproxy.yml -l firewall5.vietnix.vn  -v
cd /root/tan/proxy-tan && ansible-playbook vmware-generate-inventory.yml -l firewall6.vietnix.vn  &&  ansible-playbook vmw-deploy-autoproxy.yml -l firewall6.vietnix.vn  -v
cd /root/tan/proxy-tan && ansible-playbook vmware-generate-inventory.yml -l firewall7.vietnix.vn  &&  ansible-playbook vmw-deploy-autoproxy.yml -l firewall7.vietnix.vn  -v
cd /root/tan/proxy-tan && ansible-playbook vmware-generate-inventory.yml -l firewall8.vietnix.vn  &&  ansible-playbook vmw-deploy-autoproxy.yml -l firewall8.vietnix.vn  -v
cd /root/tan/proxy-tan && ansible-playbook vmware-generate-inventory.yml -l firewall9.vietnix.vn  &&  ansible-playbook vmw-deploy-autoproxy.yml -l firewall9.vietnix.vn  -v
cd /root/tan/proxy-tan && ansible-playbook vmware-generate-inventory.yml -l firewall10.vietnix.vn  &&  ansible-playbook vmw-deploy-autoproxy.yml -l firewall10.vietnix.vn  -v


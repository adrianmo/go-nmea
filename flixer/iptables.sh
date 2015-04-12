iptables -t nat -F

iptables -t nat -A PREROUTING -i eth0 -j DNAT -s 27.32.237.11 --protocol tcp --dport 443 --to 10.3.4.1
iptables -t nat -A PREROUTING -i eth0 -j DNAT -s 27.32.237.11 --protocol tcp --dport 80 --to 10.3.4.1

iptables -t nat -A PREROUTING -i eth0 -j DNAT -s 72.14.231.112 --protocol tcp --dport 443 --to 10.3.4.3
iptables -t nat -A PREROUTING -i eth0 -j DNAT -s 72.14.231.112 --protocol tcp --dport 80 --to 10.3.4.3

#iptables -t nat -A PREROUTING -i eth0 -j DNAT -s 0.0.0.0 --protocol tcp --dport 80 --to 172.31.6.35:8080

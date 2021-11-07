#!/bin/bash

service_ip=
service_port=
url="http://__ip__:__port__/replace"
service_url=
which curl || {
	echo "Curl could not be found, exiting."
	exit 1
}


while [ 1 -eq 1 ]; do
    ch=2
	cat << EOT
==========================================
Menu choices:
==========================================
CHOICE NUMBER        CHOICE FUNCTIONALITY
------------------------------------------
1.                   Run a test
2.                   Exit
------------------------------------------
Enter your choice:
EOT

	read choice
		ch=${choice[0]}
		case $ch in
			1)
				if [ "x$service_ip" == "x" ]; then
					echo -e "Enter StringReplacer Service's Host IP Address:\n"
					read service_ip
					if [ "x$service_ip" == "x" ]; then
						echo "No input provided, try again."
						continue
					fi
				fi
				if [ "x$service_port" == "x" ]; then
					echo -e "Enter StringReplacer Service's Port Number:\n"
					read service_port
					if [ "x$service_port" == "x" ]; then
						echo "No input provided, try again."
						continue
					fi
				fi
				service_url="$(echo $url | sed "s,__ip__,$service_ip,g; s,__port__,$service_port,g")"
				InputText=""
				KeywordString=""
				ReplacementString=""
				echo -e "Enter text line to be tested for replacement operation:\n"
				read InputText
				if [ "x$InputText" == "x" ]; then
					echo "No input provided, try again."
					continue
				fi
				echo -e "Enter Match Keyword:\n"
				read KeywordString
				if [ "x$KeywordString" == "x" ]; then
					echo "No input provided, try again."
					continue
				fi
				echo -e "Enter Replacement String:\n"
				read ReplacementString
				if [ "x$ReplacementString" == "x" ]; then
					echo "No input provided, try again."
					continue
				fi
				echo "Forming json message..."
				m="{\"InputText\":\"${InputText}\",\"KeywordString\":\"${KeywordString}\",\"ReplacementString\":\"${ReplacementString}\"}"
				echo "Sending json:\n$m\nto url: ${service_url} ..."
				reply=`curl -X POST -q -d "$m"  "${service_url}"`
				echo -e "Received response from service:\n$(echo "$reply" | jq .)"
				continue
				;;
			2)
				echo "Terminating"
				exit 0
				;;
			*)
				echo "Invalid choice ($ch) , try again."
				;;
		esac
done

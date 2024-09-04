# NAME
**ping-shutdown** - service to shutdown unix server running with UPS on ping failure e.g. during power cut

# SYNOPSIS
**ping-shutdown**
[*ip address*]
[*timeout(mins)*]

# DESCRIPTION
**ping-shutdown**

The options are as follows:

*ip address*

> Specify the ip address of device to be pinged


*timeout*

> Timeout in minutes

# EXAMPLES

	ping-shutdown 192.168.1.1 10


	Example systemctl ping-shutdown.service file supplied


# AUTHORS

sjp27 &lt; https://github.com/sjp27 &gt;
implemented ping-shutdown.

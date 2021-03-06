package utils

/*
# Mapping is calculated from AWS EC2 API using the following formula:
# * First IP on each ENI is not used for pods
# * +2 for the pods that use host-networking (AWS CNI and kube-proxy)
#
#   # of ENI * (# of IPv4 per ENI - 1) + 2
#
# https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI
#
*/

type Ec2 struct {
	Ec2Type  string
	Ec2Limit int
}

func Ec2Limits() map[string]int {
	var ec2limits = make(map[string]int)
	ec2limits["a1.2xlarge"] = 58
	ec2limits["a1.4xlarge"] = 234
	ec2limits["a1.large"] = 29
	ec2limits["a1.medium"] = 8
	ec2limits["a1.metal"] = 234
	ec2limits["a1.xlarge"] = 58
	ec2limits["c1.medium"] = 12
	ec2limits["c1.xlarge"] = 58
	ec2limits["c3.2xlarge"] = 58
	ec2limits["c3.4xlarge"] = 234
	ec2limits["c3.8xlarge"] = 234
	ec2limits["c3.large"] = 29
	ec2limits["c3.xlarge"] = 58
	ec2limits["c4.2xlarge"] = 58
	ec2limits["c4.4xlarge"] = 234
	ec2limits["c4.8xlarge"] = 234
	ec2limits["c4.large"] = 29
	ec2limits["c4.xlarge"] = 58
	ec2limits["c5.12xlarge"] = 234
	ec2limits["c5.18xlarge"] = 737
	ec2limits["c5.24xlarge"] = 737
	ec2limits["c5.2xlarge"] = 58
	ec2limits["c5.4xlarge"] = 234
	ec2limits["c5.9xlarge"] = 234
	ec2limits["c5.large"] = 29
	ec2limits["c5.metal"] = 737
	ec2limits["c5.xlarge"] = 58
	ec2limits["c5a.12xlarge"] = 234
	ec2limits["c5a.16xlarge"] = 737
	ec2limits["c5a.24xlarge"] = 737
	ec2limits["c5a.2xlarge"] = 58
	ec2limits["c5a.4xlarge"] = 234
	ec2limits["c5a.8xlarge"] = 234
	ec2limits["c5a.large"] = 29
	ec2limits["c5a.metal"] = 737
	ec2limits["c5a.xlarge"] = 58
	ec2limits["c5ad.12xlarge"] = 234
	ec2limits["c5ad.16xlarge"] = 737
	ec2limits["c5ad.24xlarge"] = 737
	ec2limits["c5ad.2xlarge"] = 58
	ec2limits["c5ad.4xlarge"] = 234
	ec2limits["c5ad.8xlarge"] = 234
	ec2limits["c5ad.large"] = 29
	ec2limits["c5ad.metal"] = 737
	ec2limits["c5ad.xlarge"] = 58
	ec2limits["c5d.12xlarge"] = 234
	ec2limits["c5d.18xlarge"] = 737
	ec2limits["c5d.24xlarge"] = 737
	ec2limits["c5d.2xlarge"] = 58
	ec2limits["c5d.4xlarge"] = 234
	ec2limits["c5d.9xlarge"] = 234
	ec2limits["c5d.large"] = 29
	ec2limits["c5d.metal"] = 737
	ec2limits["c5d.xlarge"] = 58
	ec2limits["c5n.18xlarge"] = 737
	ec2limits["c5n.2xlarge"] = 58
	ec2limits["c5n.4xlarge"] = 234
	ec2limits["c5n.9xlarge"] = 234
	ec2limits["c5n.large"] = 29
	ec2limits["c5n.metal"] = 737
	ec2limits["c5n.xlarge"] = 58
	ec2limits["c6g.12xlarge"] = 234
	ec2limits["c6g.16xlarge"] = 737
	ec2limits["c6g.2xlarge"] = 58
	ec2limits["c6g.4xlarge"] = 234
	ec2limits["c6g.8xlarge"] = 234
	ec2limits["c6g.large"] = 29
	ec2limits["c6g.medium"] = 8
	ec2limits["c6g.metal"] = 737
	ec2limits["c6g.xlarge"] = 58
	ec2limits["c6gd.12xlarge"] = 234
	ec2limits["c6gd.16xlarge"] = 737
	ec2limits["c6gd.2xlarge"] = 58
	ec2limits["c6gd.4xlarge"] = 234
	ec2limits["c6gd.8xlarge"] = 234
	ec2limits["c6gd.large"] = 29
	ec2limits["c6gd.medium"] = 8
	ec2limits["c6gd.metal"] = 737
	ec2limits["c6gd.xlarge"] = 58
	ec2limits["c6gn.12xlarge"] = 234
	ec2limits["c6gn.16xlarge"] = 737
	ec2limits["c6gn.2xlarge"] = 58
	ec2limits["c6gn.4xlarge"] = 234
	ec2limits["c6gn.8xlarge"] = 234
	ec2limits["c6gn.large"] = 29
	ec2limits["c6gn.medium"] = 8
	ec2limits["c6gn.xlarge"] = 58
	ec2limits["cc2.8xlarge"] = 234
	ec2limits["cr1.8xlarge"] = 234
	ec2limits["d2.2xlarge"] = 58
	ec2limits["d2.4xlarge"] = 234
	ec2limits["d2.8xlarge"] = 234
	ec2limits["d2.xlarge"] = 58
	ec2limits["d3.2xlarge"] = 18
	ec2limits["d3.4xlarge"] = 38
	ec2limits["d3.8xlarge"] = 59
	ec2limits["d3.xlarge"] = 10
	ec2limits["d3en.12xlarge"] = 89
	ec2limits["d3en.2xlarge"] = 18
	ec2limits["d3en.4xlarge"] = 38
	ec2limits["d3en.6xlarge"] = 58
	ec2limits["d3en.8xlarge"] = 78
	ec2limits["d3en.xlarge"] = 10
	ec2limits["f1.16xlarge"] = 394
	ec2limits["f1.2xlarge"] = 58
	ec2limits["f1.4xlarge"] = 234
	ec2limits["g2.2xlarge"] = 58
	ec2limits["g2.8xlarge"] = 234
	ec2limits["g3.16xlarge"] = 737
	ec2limits["g3.4xlarge"] = 234
	ec2limits["g3.8xlarge"] = 234
	ec2limits["g3s.xlarge"] = 58
	ec2limits["g4ad.16xlarge"] = 234
	ec2limits["g4ad.4xlarge"] = 29
	ec2limits["g4ad.8xlarge"] = 58
	ec2limits["g4dn.12xlarge"] = 234
	ec2limits["g4dn.16xlarge"] = 58
	ec2limits["g4dn.2xlarge"] = 29
	ec2limits["g4dn.4xlarge"] = 29
	ec2limits["g4dn.8xlarge"] = 58
	ec2limits["g4dn.metal"] = 737
	ec2limits["g4dn.xlarge"] = 29
	ec2limits["h1.16xlarge"] = 737
	ec2limits["h1.2xlarge"] = 58
	ec2limits["h1.4xlarge"] = 234
	ec2limits["h1.8xlarge"] = 234
	ec2limits["hs1.8xlarge"] = 234
	ec2limits["i2.2xlarge"] = 58
	ec2limits["i2.4xlarge"] = 234
	ec2limits["i2.8xlarge"] = 234
	ec2limits["i2.xlarge"] = 58
	ec2limits["i3.16xlarge"] = 737
	ec2limits["i3.2xlarge"] = 58
	ec2limits["i3.4xlarge"] = 234
	ec2limits["i3.8xlarge"] = 234
	ec2limits["i3.large"] = 29
	ec2limits["i3.metal"] = 737
	ec2limits["i3.xlarge"] = 58
	ec2limits["i3en.12xlarge"] = 234
	ec2limits["i3en.24xlarge"] = 737
	ec2limits["i3en.2xlarge"] = 58
	ec2limits["i3en.3xlarge"] = 58
	ec2limits["i3en.6xlarge"] = 234
	ec2limits["i3en.large"] = 29
	ec2limits["i3en.metal"] = 737
	ec2limits["i3en.xlarge"] = 58
	ec2limits["inf1.24xlarge"] = 321
	ec2limits["inf1.2xlarge"] = 38
	ec2limits["inf1.6xlarge"] = 234
	ec2limits["inf1.xlarge"] = 38
	ec2limits["m1.large"] = 29
	ec2limits["m1.medium"] = 12
	ec2limits["m1.small"] = 8
	ec2limits["m1.xlarge"] = 58
	ec2limits["m2.2xlarge"] = 118
	ec2limits["m2.4xlarge"] = 234
	ec2limits["m2.xlarge"] = 58
	ec2limits["m3.2xlarge"] = 118
	ec2limits["m3.large"] = 29
	ec2limits["m3.medium"] = 12
	ec2limits["m3.xlarge"] = 58
	ec2limits["m4.10xlarge"] = 234
	ec2limits["m4.16xlarge"] = 234
	ec2limits["m4.2xlarge"] = 58
	ec2limits["m4.4xlarge"] = 234
	ec2limits["m4.large"] = 20
	ec2limits["m4.xlarge"] = 58
	ec2limits["m5.12xlarge"] = 234
	ec2limits["m5.16xlarge"] = 737
	ec2limits["m5.24xlarge"] = 737
	ec2limits["m5.2xlarge"] = 58
	ec2limits["m5.4xlarge"] = 234
	ec2limits["m5.8xlarge"] = 234
	ec2limits["m5.large"] = 29
	ec2limits["m5.metal"] = 737
	ec2limits["m5.xlarge"] = 58
	ec2limits["m5a.12xlarge"] = 234
	ec2limits["m5a.16xlarge"] = 737
	ec2limits["m5a.24xlarge"] = 737
	ec2limits["m5a.2xlarge"] = 58
	ec2limits["m5a.4xlarge"] = 234
	ec2limits["m5a.8xlarge"] = 234
	ec2limits["m5a.large"] = 29
	ec2limits["m5a.xlarge"] = 58
	ec2limits["m5ad.12xlarge"] = 234
	ec2limits["m5ad.16xlarge"] = 737
	ec2limits["m5ad.24xlarge"] = 737
	ec2limits["m5ad.2xlarge"] = 58
	ec2limits["m5ad.4xlarge"] = 234
	ec2limits["m5ad.8xlarge"] = 234
	ec2limits["m5ad.large"] = 29
	ec2limits["m5ad.xlarge"] = 58
	ec2limits["m5d.12xlarge"] = 234
	ec2limits["m5d.16xlarge"] = 737
	ec2limits["m5d.24xlarge"] = 737
	ec2limits["m5d.2xlarge"] = 58
	ec2limits["m5d.4xlarge"] = 234
	ec2limits["m5d.8xlarge"] = 234
	ec2limits["m5d.large"] = 29
	ec2limits["m5d.metal"] = 737
	ec2limits["m5d.xlarge"] = 58
	ec2limits["m5dn.12xlarge"] = 234
	ec2limits["m5dn.16xlarge"] = 737
	ec2limits["m5dn.24xlarge"] = 737
	ec2limits["m5dn.2xlarge"] = 58
	ec2limits["m5dn.4xlarge"] = 234
	ec2limits["m5dn.8xlarge"] = 234
	ec2limits["m5dn.large"] = 29
	ec2limits["m5dn.xlarge"] = 58
	ec2limits["m5n.12xlarge"] = 234
	ec2limits["m5n.16xlarge"] = 737
	ec2limits["m5n.24xlarge"] = 737
	ec2limits["m5n.2xlarge"] = 58
	ec2limits["m5n.4xlarge"] = 234
	ec2limits["m5n.8xlarge"] = 234
	ec2limits["m5n.large"] = 29
	ec2limits["m5n.xlarge"] = 58
	ec2limits["m5zn.12xlarge"] = 737
	ec2limits["m5zn.2xlarge"] = 58
	ec2limits["m5zn.3xlarge"] = 234
	ec2limits["m5zn.6xlarge"] = 234
	ec2limits["m5zn.large"] = 29
	ec2limits["m5zn.metal"] = 737
	ec2limits["m5zn.xlarge"] = 58
	ec2limits["m6g.12xlarge"] = 234
	ec2limits["m6g.16xlarge"] = 737
	ec2limits["m6g.2xlarge"] = 58
	ec2limits["m6g.4xlarge"] = 234
	ec2limits["m6g.8xlarge"] = 234
	ec2limits["m6g.large"] = 29
	ec2limits["m6g.medium"] = 8
	ec2limits["m6g.metal"] = 737
	ec2limits["m6g.xlarge"] = 58
	ec2limits["m6gd.12xlarge"] = 234
	ec2limits["m6gd.16xlarge"] = 737
	ec2limits["m6gd.2xlarge"] = 58
	ec2limits["m6gd.4xlarge"] = 234
	ec2limits["m6gd.8xlarge"] = 234
	ec2limits["m6gd.large"] = 29
	ec2limits["m6gd.medium"] = 8
	ec2limits["m6gd.metal"] = 737
	ec2limits["m6gd.xlarge"] = 58
	ec2limits["mac1.metal"] = 234
	ec2limits["p2.16xlarge"] = 234
	ec2limits["p2.8xlarge"] = 234
	ec2limits["p2.xlarge"] = 58
	ec2limits["p3.16xlarge"] = 234
	ec2limits["p3.2xlarge"] = 58
	ec2limits["p3.8xlarge"] = 234
	ec2limits["p3dn.24xlarge"] = 737
	ec2limits["p4d.24xlarge"] = 737
	ec2limits["r3.2xlarge"] = 58
	ec2limits["r3.4xlarge"] = 234
	ec2limits["r3.8xlarge"] = 234
	ec2limits["r3.large"] = 29
	ec2limits["r3.xlarge"] = 58
	ec2limits["r4.16xlarge"] = 737
	ec2limits["r4.2xlarge"] = 58
	ec2limits["r4.4xlarge"] = 234
	ec2limits["r4.8xlarge"] = 234
	ec2limits["r4.large"] = 29
	ec2limits["r4.xlarge"] = 58
	ec2limits["r5.12xlarge"] = 234
	ec2limits["r5.16xlarge"] = 737
	ec2limits["r5.24xlarge"] = 737
	ec2limits["r5.2xlarge"] = 58
	ec2limits["r5.4xlarge"] = 234
	ec2limits["r5.8xlarge"] = 234
	ec2limits["r5.large"] = 29
	ec2limits["r5.metal"] = 737
	ec2limits["r5.xlarge"] = 58
	ec2limits["r5a.12xlarge"] = 234
	ec2limits["r5a.16xlarge"] = 737
	ec2limits["r5a.24xlarge"] = 737
	ec2limits["r5a.2xlarge"] = 58
	ec2limits["r5a.4xlarge"] = 234
	ec2limits["r5a.8xlarge"] = 234
	ec2limits["r5a.large"] = 29
	ec2limits["r5a.xlarge"] = 58
	ec2limits["r5ad.12xlarge"] = 234
	ec2limits["r5ad.16xlarge"] = 737
	ec2limits["r5ad.24xlarge"] = 737
	ec2limits["r5ad.2xlarge"] = 58
	ec2limits["r5ad.4xlarge"] = 234
	ec2limits["r5ad.8xlarge"] = 234
	ec2limits["r5ad.large"] = 29
	ec2limits["r5ad.xlarge"] = 58
	ec2limits["r5b.12xlarge"] = 234
	ec2limits["r5b.16xlarge"] = 737
	ec2limits["r5b.24xlarge"] = 737
	ec2limits["r5b.2xlarge"] = 58
	ec2limits["r5b.4xlarge"] = 234
	ec2limits["r5b.8xlarge"] = 234
	ec2limits["r5b.large"] = 29
	ec2limits["r5b.metal"] = 737
	ec2limits["r5b.xlarge"] = 58
	ec2limits["r5d.12xlarge"] = 234
	ec2limits["r5d.16xlarge"] = 737
	ec2limits["r5d.24xlarge"] = 737
	ec2limits["r5d.2xlarge"] = 58
	ec2limits["r5d.4xlarge"] = 234
	ec2limits["r5d.8xlarge"] = 234
	ec2limits["r5d.large"] = 29
	ec2limits["r5d.metal"] = 737
	ec2limits["r5d.xlarge"] = 58
	ec2limits["r5dn.12xlarge"] = 234
	ec2limits["r5dn.16xlarge"] = 737
	ec2limits["r5dn.24xlarge"] = 737
	ec2limits["r5dn.2xlarge"] = 58
	ec2limits["r5dn.4xlarge"] = 234
	ec2limits["r5dn.8xlarge"] = 234
	ec2limits["r5dn.large"] = 29
	ec2limits["r5dn.xlarge"] = 58
	ec2limits["r5n.12xlarge"] = 234
	ec2limits["r5n.16xlarge"] = 737
	ec2limits["r5n.24xlarge"] = 737
	ec2limits["r5n.2xlarge"] = 58
	ec2limits["r5n.4xlarge"] = 234
	ec2limits["r5n.8xlarge"] = 234
	ec2limits["r5n.large"] = 29
	ec2limits["r5n.xlarge"] = 58
	ec2limits["r6g.12xlarge"] = 234
	ec2limits["r6g.16xlarge"] = 737
	ec2limits["r6g.2xlarge"] = 58
	ec2limits["r6g.4xlarge"] = 234
	ec2limits["r6g.8xlarge"] = 234
	ec2limits["r6g.large"] = 29
	ec2limits["r6g.medium"] = 8
	ec2limits["r6g.metal"] = 737
	ec2limits["r6g.xlarge"] = 58
	ec2limits["r6gd.12xlarge"] = 234
	ec2limits["r6gd.16xlarge"] = 737
	ec2limits["r6gd.2xlarge"] = 58
	ec2limits["r6gd.4xlarge"] = 234
	ec2limits["r6gd.8xlarge"] = 234
	ec2limits["r6gd.large"] = 29
	ec2limits["r6gd.medium"] = 8
	ec2limits["r6gd.metal"] = 737
	ec2limits["r6gd.xlarge"] = 58
	ec2limits["t1.micro"] = 4
	ec2limits["t2.2xlarge"] = 44
	ec2limits["t2.large"] = 35
	ec2limits["t2.medium"] = 17
	ec2limits["t2.micro"] = 4
	ec2limits["t2.nano"] = 4
	ec2limits["t2.small"] = 11
	ec2limits["t2.xlarge"] = 44
	ec2limits["t3.2xlarge"] = 58
	ec2limits["t3.large"] = 35
	ec2limits["t3.medium"] = 17
	ec2limits["t3.micro"] = 4
	ec2limits["t3.nano"] = 4
	ec2limits["t3.small"] = 11
	ec2limits["t3.xlarge"] = 58
	ec2limits["t3a.2xlarge"] = 58
	ec2limits["t3a.large"] = 35
	ec2limits["t3a.medium"] = 17
	ec2limits["t3a.micro"] = 4
	ec2limits["t3a.nano"] = 4
	ec2limits["t3a.small"] = 8
	ec2limits["t3a.xlarge"] = 58
	ec2limits["t4g.2xlarge"] = 58
	ec2limits["t4g.large"] = 35
	ec2limits["t4g.medium"] = 17
	ec2limits["t4g.micro"] = 4
	ec2limits["t4g.nano"] = 4
	ec2limits["t4g.small"] = 11
	ec2limits["t4g.xlarge"] = 58
	ec2limits["u-12tb1.metal"] = 147
	ec2limits["u-18tb1.metal"] = 737
	ec2limits["u-24tb1.metal"] = 737
	ec2limits["u-6tb1.metal"] = 147
	ec2limits["u-9tb1.metal"] = 147
	ec2limits["x1.16xlarge"] = 234
	ec2limits["x1.32xlarge"] = 234
	ec2limits["x1e.16xlarge"] = 234
	ec2limits["x1e.2xlarge"] = 58
	ec2limits["x1e.32xlarge"] = 234
	ec2limits["x1e.4xlarge"] = 58
	ec2limits["x1e.8xlarge"] = 58
	ec2limits["x1e.xlarge"] = 29
	ec2limits["z1d.12xlarge"] = 737
	ec2limits["z1d.2xlarge"] = 58
	ec2limits["z1d.3xlarge"] = 234
	ec2limits["z1d.6xlarge"] = 234
	ec2limits["z1d.large"] = 29
	ec2limits["z1d.metal"] = 737
	ec2limits["z1d.xlarge"] = 58

	return ec2limits
}

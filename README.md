# md-code-formatter

A tool to format the HCL code blocks in markdown files

# Install

Install `mdcoder` binary into `$GOPATH/bin` after running the following command

```sh
$ go install github.com:imjoey/md-code-formatter/cmd/mdcoder
```

# Usage

```sh
$ mdcoder -h
md-code-formatter: A tool to format the HCL code blocks in markdown files
Usage: md-code-formatter [-d directory]

Options:
  -d string
    	The directory that contains the markdown files (default "./")


$ mdcoder -d website/docs

2019/07/17 15:16:46 Formatted file <website/docs/d/antiddos_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/d/csbs_backup_policy_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/d/csbs_backup_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/d/cts_tracker_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/d/dcs_maintainwindow_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/d/dcs_product_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/d/dms_maintainwindow_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/d/dms_product_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/d/images_image_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/d/rds_flavors_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/d/vbs_backup_policy_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/as_configuration_v1.html.markdown> : 3 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/as_group_v1.html.markdown> : 3 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/as_policy_v1.html.markdown> : 3 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/cdm_cluster_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/cloudtable_cluster_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/compute_instance_v2.html.markdown> : 2 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/compute_secgroup_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/cs_peering_connect_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/csbs_backup_policy_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/csbs_backup_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/css_cluster_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/cts_tracker_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/dcs_instance_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/dis_stream_v2.html.markdown> : 2 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/dms_group_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/dms_instance_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/dms_queue_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/dns_recordset_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/dns_zone_v2.html.markdown> : 2 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/dws_cluster.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/elb_backendecs.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/elb_healthcheck.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/elb_listener.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/elb_loadbalancer.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/fw_firewall_group_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/ges_graph_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/iam_agency_v3.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/identity_group_membership_v3.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/identity_group_v3.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/identity_project_v3.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/identity_role_assignment_v3.html.markdown> : 2 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/identity_user_v3.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/images_image_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/lb_l7policy_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/lb_l7rule_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/mls_instance.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/mrs_cluster_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/mrs_job_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/nat_gateway_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/nat_snat_rule_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/networking_floatingip_associate_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/networking_secgroup_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/rds_instance_v1.html.markdown> : 3 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/rts_software_config_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/s3_bucket.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/s3_bucket_policy.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/sfs_file_system_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/smn_subscription_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/smn_topic_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/vbs_backup_policy_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/vbs_backup_v2.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Formatted file <website/docs/r/vpc_eip_v1.html.markdown> : 1 code blocks formatted
2019/07/17 15:16:46 Statistics: 63 of 108 markdown files formatted
```

package test

const (
	query string = `
				select 'addressdata' as job, count(*) as count  from addressdata_previous
				union
				select 'addressdataconnection' as job, count(*) as count from addressdataconnection_previous
				union
				select 'customerinvoicestructure' as job, count(*) as count from customerinvoicestructure_previous
				union
				select 'nationalbuyer' as job, count(*) as count from nationalbuyer_previous
				union
				select 'partnergln' as job, count(*) as count from partnergln_previous
				union
				select 'partnerhierarchy' as job, count(*) as count from partnerhierarchy_previous
				union
				select 'partnernationalbuyerconnection' as job, count(*) as count from partnernationalbuyerconnection_previous
				union
				select 'partnerserviceassignment' as job, count(*) as count from partnerserviceassignment_previous
				union
				select 'partnersolvency' as job, count(*) as count from partnersolvency_previous
				union
				select 'partnersupplementalcontractinfo' as job, count(*) as count from partnersupplementalcontractinfo_previous
				union
				select 'partnervatid' as job, count(*) as count from partnervatid_previous
				union
				select 'productgroup' as job, count(*) as count from productgroup_previous
				union
				select 'sawithaddressid' as job, count(*) as count from sawithaddressid_previous
				union
				select 'sawithgln' as job, count(*) as count from sawithgln_previous sp
				union
				select 'vatrate' as job, count(*) as count from vatrate_previous vp;`
)

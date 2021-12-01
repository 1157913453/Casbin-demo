## Casbin的一个demo

### 权限说明：
- gh和hh两个租户，每个租户都有dev和admin角色，每个租户都有data1和data2数据
- gh有角色gh_dev和gh_admin，hh有角色hh_dev和hh_admin
- 每个租户的dev都对其域的data1数据有读写权限
- 每个租户的admin除了拥有其dev的所有权限外，对域的data2还有读权限
- system对所有域的所有数据拥有所有权限

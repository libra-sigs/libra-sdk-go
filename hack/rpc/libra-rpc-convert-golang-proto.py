import os
from os import rename
from os import listdir
from os.path import isfile, join


proto_info = {}
rpc_root = os.sys.argv[1]
for proto in ["types", "mempool_status", "admission_control"]:
  proto_info[proto] = [f for f in listdir(join(rpc_root, proto)) if isfile(join(rpc_root, proto, f))]

print proto_info

rpc_prefix="github.com/libra-sigs/libra-sdk-go/libra/rpc"
for proto in ["types", "mempool_status", "admission_control"]:
  for f in proto_info[proto]:
    in_file = "%s/%s/%s" % (rpc_root, proto, f)
    out_file = "%s.tmp" % in_file
    lines =  [ l for l in open(in_file).readlines()]
    out = open(out_file, 'w+')
    go_package_title = "option go_package = \"%s/%s\";\n" % (rpc_prefix, proto)
    file_need_update = False
    for l in lines:
      package_title = "package %s;" % proto
      for t_proto in ["types", "mempool_status", "admission_control"]:
        need_update = False
        for t_f in proto_info[t_proto]:
          #if proto == t_proto:
          #  continue
          needle = "import \"%s\"" % t_f
          n_needle = "import \"%s/%s\"" % (t_proto, t_f)
          if l.find(needle) == 0:
            l = l.replace(needle, n_needle)
            need_update = True
            break
        if need_update:
          file_need_update = True
          break
      if l.find(package_title) == 0:
        out.write(go_package_title)
      out.write(l)
    out.close()
    rename(out_file, in_file)

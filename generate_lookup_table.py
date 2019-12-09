import pyvaruint as var

print("table = {")
for i in range(2**32):
    data = var.encode_varint(i)
    print('    b"%s": %s' % (data.hex(), i))
print("}")


import re

file_path = r'D:\AlexCiawi\mobile\vendor_portal\lib\features\vendor_registration\controllers\vendor_registration_controller.dart'

with open(file_path, 'r', encoding='utf-8') as f:
    content = f.read()

# We want to find inal payload = { ... }; inside SubmitRegistration()
# And remove the "Step*": { } wrappers, also the "Address": { }, "MainContact": { } wrappers

payload_start = content.find('final payload = {')
if payload_start == -1:
    print('Payload not found')
    exit(1)

payload_end = content.find('};', payload_start)

payload_block = content[payload_start:payload_end+2]

# Regex to find all key-value pairs like "Key": Value,
pairs = re.findall(r'"([A-Za-z0-9_]+)":\s*(.+?),', payload_block)

new_payload_lines = ["      final payload = {"]
for key, value in pairs:
    # Skip Step prefixes and object wrappers
    if key.startswith('Step') or key in ['Address', 'MainContact', 'InquiryContact', 'PoContact', 'PaymentContact', 'Attachments']:
        continue
    new_payload_lines.append(f'        "{key}": {value},')

new_payload_lines.append("      };")

new_payload_block = '\n'.join(new_payload_lines)

new_content = content[:payload_start] + new_payload_block + content[payload_end+2:]

with open(file_path, 'w', encoding='utf-8') as f:
    f.write(new_content)
print("Done")
